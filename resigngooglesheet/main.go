package main

import (
	"belajargolang/api/resign/helper"
	"belajargolang/api/resign/models"
	"belajargolang/api/resign/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	UploadSubmission()
}

func UploadSubmission() {

	resp, err := http.Get("https://script.googleusercontent.com/macros/echo?user_content_key=en9LPCpWxHIi6ypowj516ylITkOMcRm0eaorXRA-GMyKo-ng7B_4N3uJdiSRP8a2Q7gPMCZq1dbxs47rN-kyK-ATKocFrm8vm5_BxDlH2jW0nuo2oDemN9CCS2h10ox_1xSncGQajx_ryfhECjZEnF7XA4IZnfrhIsfLjiFuVgnfUn6GRRnJuLRSaK1nuhAc5gBTl6uWHFKtr26VJn7A0dwf6k5YrcAy1T9U3EAstbBKoSOdad6tfw&lib=MP2gm8nYgyptxZX7bL1YrGUbkDnjr7vlQ")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	// sb := string(body)
	// log.Printf(sb)
	// fmt.Println(reflect.TypeOf(sb))

	var data map[string][][]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Jumlah :", len(data["content"]))
	/*
	   // Sudah Bisa Convert
	   	for i := 0; i < len(data["content"]); i++ {
	   		var get_num_str, rep_nik, number_of_employees string
	   		get_num_str = fmt.Sprintf("%f", data["content"][i][1])
	   		rep_nik = strings.Replace(get_num_str, ".", "", -1)
	   		number_of_employees = rep_nik[0:10]
	   		fmt.Println(reflect.TypeOf(data["content"][i][0]), number_of_employees, reflect.TypeOf(data["content"][i][4]))
	   	}
	*/

	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbresign, _ = models.ConnResign()
	defer dbresign.Close()

	var notification []string
	notification = append(notification, "")
	var code int = 200

	for _, record := range records {

		var Count_employees int
		var Status_employee, Number_of_employees string

		err = db.
			QueryRow("select COUNT(id), COALESCE(number_of_employees, 'NULL'), COALESCE(status_employee, 'NULL') FROM employees where number_of_employees = ? ", record[1]).
			Scan(&Count_employees, &Number_of_employees, &Status_employee)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		switch Count_employees {
		case 1: // Karyawan di temukan

			var resultemployee = models.Employee{}
			// Menampilkan data karyawan
			err = db.QueryRow("select COALESCE(name, ''), COALESCE(date_of_birth, '0000-00-00'), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') as date_out, status_employee from employees where number_of_employees = ? ", Number_of_employees).
				Scan(&resultemployee.Name, &resultemployee.Date_of_birth, &resultemployee.Hire_date, &resultemployee.Date_out, &resultemployee.Status_employee)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var Count_resigns int
			// Cari data resign pada pada database hwi berdasarkan nik
			err = dbresign.QueryRow("select count(id) as count_resign from resigns where number_of_employees = ? ", Number_of_employees).
				Scan(&Count_resigns)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var Count_resign_submissions int
			var Count_status_resign_submissions string

			// Cari data pengajuan resign pada  database HWI
			err = dbresign.QueryRow("select count(id) as count_resign_submissions, COALESCE(status_resignsubmisssion, 'NULL') as  status_resignsubmisssion from resignation_submissions where number_of_employees = ? ", Number_of_employees).
				Scan(&Count_resign_submissions, &Count_status_resign_submissions)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// ===================== RESULT ==========================
			var resignation_submission_id int

			if Count_resigns > 0 && Count_resign_submissions > 0 && Count_status_resign_submissions != "cancel" {
				//Jika sudah resign dan sudah mengajukan dan pegajuannya tidak cancel maka tidak dapat mengajukan lagi
				each := fmt.Sprintf("NIK ini %s tidak dapat mengajukan resign karena sudah mengajukan resign dan status karyawan sudah resign. \n", Number_of_employees)
				notification = append(notification, each)
				code = 400

			} else if Count_resigns == 0 && Count_resign_submissions > 0 && Count_status_resign_submissions != "cancel" {
				//Jika belum resign dan sudah mengajukan dan pegajuannya tidak cancel maka tidak dapat mengajukan lagi
				each := fmt.Sprintf("NIK ini %s tidak dapat mengajukan resign karena sudah resign dengan status pengajuan wait. \n", Number_of_employees)
				notification = append(notification, each)
				code = 400

			} else if Count_resign_submissions == 0 {
				var typesubmissions, classificationsubmissions string
				typesubmissions = "false"
				classificationsubmissions = "Mengajukan permohonan resign setelah karyawan resign"

				switch Count_resigns {
				case 0:
					if Status_employee == "active" {
						typesubmissions = helper.DateSubmissionCompareRequest(record[0], record[4])
						if typesubmissions == "true" {
							classificationsubmissions = "Mengajukan permohonan resign sebelum karyawan resign"
						}
					}
				}

				var datasubmissions = map[string]interface{}{
					"number_of_employees":          Number_of_employees,
					"name":                         resultemployee.Name,
					"position":                     record[5],
					"department":                   record[9],
					"building":                     record[9],
					"address":                      record[3],
					"hire_date":                    resultemployee.Hire_date,
					"date_out":                     resultemployee.Date_out,
					"date_resignation_submissions": record[4],
					"type":                         typesubmissions,
					"reason":                       record[3],
					"detail_reason":                record[6],
					"periode_of_service":           helper.Periode_of_serve(resultemployee.Hire_date, record[4]),
					"age":                          helper.Age(resultemployee.Date_of_birth),
					"suggestion":                   record[7],
					"status_resignsubmisssion":     "wait",
					"using_media":                  "google",
					"classification":               classificationsubmissions,
					"print":                        0,
					"created_at":                   record[0],
					"updated_at":                   record[0],
				}
				repository.InsertResign("resignation_submissions", datasubmissions)
				err = dbresign.QueryRow("SELECT id FROM resignation_submissions WHERE number_of_employees = ? AND created_at = ? ", Number_of_employees, record[0]).
					Scan(&resignation_submission_id)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				var datakuesioners = map[string]interface{}{
					"resignation_submission_id": resignation_submission_id,
					"number_of_employees":       Number_of_employees,
					"k1":                        record[10],
					"k2":                        record[11],
					"k3":                        record[12],
					"k4":                        record[13],
					"k5":                        record[14],
					"k6":                        record[15],
					"k7":                        record[16],
					"created_at":                record[0],
					"updated_at":                record[0],
				}
				repository.InsertResign("kuesioners", datakuesioners)
			}

		default:
		}
	}

	//untuk membuat json pertama kita harus set Header
	w.Header().Set("Content-Type", "application/json")
	result := map[string]interface{}{
		"code":    code,
		"data":    notification,
		"message": "Succesfully",
	}
	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))
}
