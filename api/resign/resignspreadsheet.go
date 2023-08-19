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
	"sync"

	"strconv"
	"strings"
	"time"
)

// mutex to protect the shared resource
var mu sync.Mutex

// wait group to wait for all goroutines to finish
var wg sync.WaitGroup

func main() {
	fmt.Println("RUN PROGRAM")
	// ExecuteSpreadsheet()

	// CheckSyncron := time.NewTicker(2 * time.Second)
	UpdateSpreadsheet := time.NewTicker(30 * time.Second)
	// UpdateMasterToResign := time.NewTicker(2 * time.Minute)
	for {
		select {
		case <-UpdateSpreadsheet.C:
			// ExecuteSpreadsheet()
			wg.Add(2)
			go ExecuteSpreadsheet()
			go UpdateResignFromDatamaster()
			// fmt.Println("Execute Spreadsheet")
			// case <-UpdateMasterToResign.C:

			// ExecuteSpreadsheet()
			// fmt.Println("Execute UpdateMasterToResign")
			// case <-CheckSyncron.C:
			// ExecuteSpreadsheet()
			// fmt.Println("Check Sybc")
		}
	}
	wg.Wait()
}

func Keluar() {
	mu.Lock()
	fmt.Println("Keluar")
	mu.Unlock()
	wg.Done()
}

func Masuk() {
	mu.Lock()
	fmt.Println("Masuk")
	mu.Unlock()
	wg.Done()
}

func ExecuteSpreadsheet() {
	mu.Lock()
	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbresign, _ = models.ConnResign()
	defer dbresign.Close()

	resp, err := http.Get("https://script.googleusercontent.com/macros/echo?user_content_key=en9LPCpWxHIi6ypowj516ylITkOMcRm0eaorXRA-GMyKo-ng7B_4N3uJdiSRP8a2Q7gPMCZq1dbxs47rN-kyK-ATKocFrm8vm5_BxDlH2jW0nuo2oDemN9CCS2h10ox_1xSncGQajx_ryfhECjZEnF7XA4IZnfrhIsfLjiFuVgnfUn6GRRnJuLRSaK1nuhAc5gBTl6uWHFKtr26VJn7A0dwf6k5YrcAy1T9U3EAstbBKoSOdad6tfw&lib=MP2gm8nYgyptxZX7bL1YrGUbkDnjr7vlQ")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data map[string][][]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Sudah Bisa Convert
	for i := 0; i < len(data["content"]); i++ {
		var get_num_str, rep_nik, number_of_employees string
		get_num_str = fmt.Sprintf("%f", data["content"][i][1])
		rep_nik = strings.Replace(get_num_str, ".", "", -1)
		number_of_employees = rep_nik[0:10]

		var Count_employees int
		var Status_employee, Number_of_employees string

		// Cari Karyawan
		err = db.
			QueryRow("select COUNT(id), COALESCE(number_of_employees, 'NULL'), COALESCE(status_employee, 'NULL') FROM employees where number_of_employees = ? ", number_of_employees).
			Scan(&Count_employees, &Number_of_employees, &Status_employee)
		if err != nil {
			fmt.Println(err.Error(), "Cari Karyawan", number_of_employees)
			return
		}

		switch Count_employees {
		case 1: // Karyawan di temukan

			var resultemployee = models.Employee{}
			// Menampilkan data karyawan
			err = db.QueryRow("select COALESCE(name, ''), COALESCE(date_of_birth, '0000-00-00'), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') as date_out, status_employee from employees where number_of_employees = ? ", Number_of_employees).
				Scan(&resultemployee.Name, &resultemployee.Date_of_birth, &resultemployee.Hire_date, &resultemployee.Date_out, &resultemployee.Status_employee)
			if err != nil {
				fmt.Println(err.Error(), "Menampilkan data karyawan", Number_of_employees)
				return
			}

			var Count_resigns int
			// Cari data resign pada pada database hwi berdasarkan nik
			err = dbresign.QueryRow("select count(id) as count_resign from resigns where number_of_employees = ? ", Number_of_employees).
				Scan(&Count_resigns)
			if err != nil {
				fmt.Println(err.Error(), "Cari data resign pada pada database hwi berdasarkan nik", Number_of_employees)
				return
			}

			var Count_resign_submissions int
			var Count_status_resign_submissions string

			// Cari data pengajuan resign pada  database HWI
			err = dbresign.QueryRow("select count(id) as count_resign_submissions, COALESCE(status_resignsubmisssion, 'NULL') as  status_resignsubmisssion from resignation_submissions where number_of_employees = ? AND status_resignsubmisssion != 'cancel' ", Number_of_employees).
				Scan(&Count_resign_submissions, &Count_status_resign_submissions)
			if err != nil {
				fmt.Println(err.Error(), "Cari data pengajuan resign pada  database HWI", Number_of_employees)
				return
			}

			// ===================== RESULT ==========================
			var resignation_submission_id int
			if Count_resigns > 0 && Count_resign_submissions > 0 {
			} else if Count_resigns == 0 && Count_resign_submissions > 0 {
			} else if Count_resign_submissions == 0 {
				var typesubmissions, classificationsubmissions string
				typesubmissions = "false"
				classificationsubmissions = "Mengajukan permohonan resign setelah karyawan resign"

				switch Count_resigns {
				case 0:
					if Status_employee == "active" {
						typesubmissions = helper.DateSubmissionCompareRequest(fmt.Sprint(data["content"][i][0]), fmt.Sprint(data["content"][i][4]))
						if typesubmissions == "true" {
							classificationsubmissions = "Mengajukan permohonan resign sebelum karyawan resign"
						}
					}
				}

				var datasubmissions = map[string]interface{}{
					"number_of_employees":          Number_of_employees,
					"name":                         resultemployee.Name,
					"position":                     data["content"][i][5],
					"department":                   data["content"][i][9],
					"building":                     data["content"][i][9],
					"address":                      data["content"][i][3],
					"hire_date":                    resultemployee.Hire_date,
					"date_out":                     resultemployee.Date_out,
					"date_resignation_submissions": data["content"][i][4],
					"type":                         typesubmissions,
					"reason":                       data["content"][i][3],
					"detail_reason":                data["content"][i][6],
					"periode_of_service":           helper.Periode_of_serve(resultemployee.Hire_date, fmt.Sprint(data["content"][i][4])),
					"age":                          helper.Age(resultemployee.Date_of_birth),
					"suggestion":                   data["content"][i][7],
					"status_resignsubmisssion":     "wait",
					"using_media":                  "google",
					"classification":               classificationsubmissions,
					"print":                        0,
					"created_at":                   data["content"][i][0],
					"updated_at":                   data["content"][i][0],
				}
				repository.InsertResign("resignation_submissions", datasubmissions)
				err = dbresign.QueryRow("SELECT id FROM resignation_submissions WHERE number_of_employees = ?  ORDER BY id DESC", Number_of_employees).
					Scan(&resignation_submission_id)
				if err != nil {
					fmt.Println(err.Error(), "INSERT submission ", Number_of_employees)
					return
				}
				var datakuesioners = map[string]interface{}{
					"resignation_submission_id": resignation_submission_id,
					"number_of_employees":       Number_of_employees,
					"k1":                        cellfloattoint(data["content"][i][10]),
					"k2":                        cellfloattoint(data["content"][i][11]),
					"k3":                        cellfloattoint(data["content"][i][12]),
					"k4":                        cellfloattoint(data["content"][i][13]),
					"k5":                        cellfloattoint(data["content"][i][14]),
					"k6":                        cellfloattoint(data["content"][i][15]),
					"k7":                        cellfloattoint(data["content"][i][16]),
					"created_at":                data["content"][i][0],
					"updated_at":                data["content"][i][0],
				}
				repository.InsertResign("kuesioners", datakuesioners)
			}

		default:
		}
	}

	fmt.Println("Jumlah Data Di proses:", len(data["content"]))
	mu.Unlock()
	wg.Done()
}

func cellfloattoint(data interface{}) int {
	var get_num_str, rep_str, get_str string
	get_num_str = fmt.Sprintf("%f", data)
	rep_str = strings.Replace(get_num_str, ".", "", -1)
	get_str = rep_str[0:1]
	var num, _ = strconv.ParseInt(get_str, 10, 64)
	return int(num)
}

func UpdateResignFromDatamaster() {
	mu.Lock()
	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbresign, _ = models.ConnResign()
	defer dbresign.Close()

	var count int

	err = dbresign.QueryRow("SELECT COUNT(*) FROM resignation_submissions WHERE status_resignsubmisssion != 'cancel' AND (date_out IS NULL OR date_out = '0000-00-00') ").
		Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch count {
	case 0:
		break
	default:

		var id int
		var number_of_employees string
		rows, err := dbresign.Query("SELECT id, number_of_employees FROM resignation_submissions WHERE status_resignsubmisssion != 'cancel' AND (date_out IS NULL OR date_out = '0000-00-00') ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&id, &number_of_employees)
			wg.Add(1)
			go Checkempnull(id, number_of_employees)
		}

		break
	}

	fmt.Println("Data master Berhasil di proses => ")
	mu.Unlock()
	wg.Done()

}

func Checkempnull(id int, number_of_employees string) {
	mu.Lock()
	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbresign, _ = models.ConnResign()
	defer dbresign.Close()

	var date_out string
	err = db.QueryRow("SELECT COALESCE(date_out, '0000-00-00') FROM employees WHERE number_of_employees = ? ", number_of_employees).
		Scan(&date_out)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if date_out != "0000-00-00" {
		wg.Add(1)
		go ExecuteUpdate(id, number_of_employees, date_out)
	}
	mu.Unlock()
	wg.Done()
}

func ExecuteUpdate(id int, number_of_employees string, date_out string) {
	mu.Lock()
	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbresign, _ = models.ConnResign()
	defer dbresign.Close()

	_, err = dbresign.Exec("UPDATE resignation_submissions SET date_out = ? WHERE id = ?", date_out, id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = dbresign.Exec("UPDATE resigns SET date_out = ? WHERE number_of_employees = ? ", date_out, number_of_employees)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	mu.Unlock()
	wg.Done()
}
