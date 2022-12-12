package controllers

import (
	"belajargolang/api/resign/helper"
	"belajargolang/api/resign/models"
	"belajargolang/api/resign/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	result := map[string]string{
		"data": "Connection Succesfully",
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dbhwi, err = models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	var Count_Submission, Count_Resign, Count_Certificate, Count_Experience, Count_Kuesioner int

	err = dbhwi.QueryRow("SELECT COUNT(id) as id from resignation_submissions").Scan(&Count_Submission)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dbhwi.QueryRow("SELECT COUNT(id) as id from kuesioners").Scan(&Count_Kuesioner)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dbhwi.QueryRow("SELECT COUNT(id) as id from resigns").Scan(&Count_Resign)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dbhwi.QueryRow("SELECT COUNT(id) as id from certificate_of_employments").Scan(&Count_Certificate)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dbhwi.QueryRow("SELECT COUNT(id) as id from work_experience_letters").Scan(&Count_Experience)
	if err != nil {
		fmt.Println(err.Error())
	}

	data := map[string]int{
		"submission":         Count_Submission,
		"kuesioner":          Count_Kuesioner,
		"resign":             Count_Resign,
		"letter_certificate": Count_Certificate,
		"letter_experience":  Count_Experience,
	}

	result := map[string]interface{}{
		"data":    data,
		"code":    200,
		"message": "Successfully",
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resp))
	return
}

func EmployeeAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	dbhwi, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	var message string

	if r.Method == "POST" {
		var data models.Employee
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if data.Status_employee == "active" {
			//Menghapus data resign
			where := fmt.Sprintf("number_of_employees = '%s' ", data.Number_of_employees)
			repository.DeleteResign("work_experience_letters", where)
			repository.DeleteResign("certificate_of_employments", where)
			repository.DeleteResign("resigns", where)
			message = fmt.Sprintf("Data resign %s berhasil di hapus", data.Number_of_employees)

		} else if data.Status_employee == "notactive" {
			//menambah data resign
			var Count_id int
			var Employee models.Employee

			err = db.QueryRow("SELECT COUNT(id) as id , COALESCE(status_employee, '') as status_employee, COALESCE(name, ''), COALESCE(hire_date, ''), COALESCE(date_of_birth, ''), COALESCE(date_out, '0000-00-00'),COALESCE(job_id, 25),COALESCE(department_id, 116), COALESCE(address_jalan, ''), COALESCE(address_rt, ''), COALESCE(address_rw, ''), COALESCE(address_village, ''), COALESCE(address_district, ''), COALESCE(address_city, ''), COALESCE(address_province, '') FROM employees WHERE number_of_employees = ? ", data.Number_of_employees).
				Scan(&Count_id, &Employee.Status_employee, &Employee.Name, &Employee.Hire_date, &Employee.Date_of_birth, &Employee.Date_out, &Employee.Job_id, &Employee.Department_id, &Employee.Address_jalan, &Employee.Address_rt, &Employee.Address_rw, &Employee.Address_village, &Employee.Address_district, &Employee.Address_city, &Employee.Address_province)
			if err != nil {
				fmt.Println(err.Error())
			}
			switch Count_id {
			case 1:
				var Count_resign int
				queryrow_resign := fmt.Sprintf("SELECT COUNT(id) as id FROM resigns WHERE number_of_employees = %s", data.Number_of_employees)
				err = dbhwi.QueryRow(queryrow_resign).Scan(&Count_resign)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				if Count_resign == 0 {
					/*
						_, err = dbhwi.Exec("INSERT INTO resigns (`number_of_employees`, `name`, `position`, `department`, `hire_date`, `classification`, `date_out`, `date_resignsubmissions`, `periode_of_service`, `type`, `age`, `status_resign`,  `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
							data.Number_of_employees, Employee.Name, JobDepartment(data.Number_of_employees)[0], JobDepartment(data.Number_of_employees)[1], Employee.Hire_date, helper.CekDateSubmission(data.Number_of_employees), data.Date_out, nil, helper.Periode_of_serve(Employee.Hire_date, data.Date_out), helper.TypeResign(data.Number_of_employees, data.Date_out)["type"], helper.Age(Employee.Date_of_birth), helper.TypeResign(data.Number_of_employees, data.Date_out)["status"], helper.DMYhms(), helper.DMYhms())
						if err != nil {
							fmt.Println(err.Error())
							return
						}
					*/
					var data = map[string]interface{}{
						"number_of_employees":    data.Number_of_employees,
						"name":                   Employee.Name,
						"position":               JobDepartment(data.Number_of_employees)[0],
						"department":             JobDepartment(data.Number_of_employees)[1],
						"hire_date":              Employee.Hire_date,
						"classification":         helper.TypeResign(data.Number_of_employees, data.Date_out)["classification"],
						"date_out":               data.Date_out,
						"date_resignsubmissions": data.Date_out,
						"periode_of_service":     helper.Periode_of_serve(Employee.Hire_date, data.Date_out),
						"type":                   helper.TypeResign(data.Number_of_employees, data.Date_out)["type"],
						"age":                    helper.Periode_of_serve(Employee.Hire_date, data.Date_out),
						"status_resign":          helper.TypeResign(data.Number_of_employees, data.Date_out)["status"],
						"printed":                0,
						"created_at":             helper.DMYhms(),
						"updated_at":             helper.DMYhms(),
					}
					repository.InsertResign("resigns", data)
				} else if Count_resign > 0 {
					/*
						_, err = dbhwi.Exec("UPDATE resigns SET number_of_employees = ? , name = ?, position = ?, department = ?, hire_date = ?, classification = ?, date_out = ?, date_resignsubmissions = ?, periode_of_service = ?, type = ?, age = ?, status_resign = ?, created_at = ?, updated_at = ? WHERE number_of_employees = ?", data.Number_of_employees, Employee.Name, JobDepartment(data.Number_of_employees)[0], JobDepartment(data.Number_of_employees)[1], Employee.Hire_date, helper.CekDateSubmission(data.Number_of_employees), data.Date_out, nil, helper.Periode_of_serve(Employee.Hire_date, data.Date_out), helper.TypeResign(data.Number_of_employees, data.Date_out)["type"], helper.Age(Employee.Date_of_birth), helper.TypeResign(data.Number_of_employees, data.Date_out)["status"], helper.DMYhms(), helper.DMYhms(), data.Number_of_employees)
						if err != nil {
							fmt.Println(err.Error())
							return
						}
					*/
					var data1 = map[string]interface{}{
						"number_of_employees":    data.Number_of_employees,
						"name":                   Employee.Name,
						"position":               JobDepartment(data.Number_of_employees)[0],
						"department":             JobDepartment(data.Number_of_employees)[1],
						"hire_date":              Employee.Hire_date,
						"classification":         helper.TypeResign(data.Number_of_employees, data.Date_out)["classification"],
						"date_out":               data.Date_out,
						"date_resignsubmissions": data.Date_out,
						"periode_of_service":     helper.Periode_of_serve(Employee.Hire_date, data.Date_out),
						"type":                   helper.TypeResign(data.Number_of_employees, data.Date_out)["type"],
						"age":                    helper.Periode_of_serve(Employee.Hire_date, data.Date_out),
						"status_resign":          helper.TypeResign(data.Number_of_employees, data.Date_out)["status"],
						"printed":                0,
						"created_at":             helper.DMYhms(),
						"updated_at":             helper.DMYhms(),
					}
					where := fmt.Sprintf("number_of_employees = '%s' ", data.Number_of_employees)
					repository.UpdateResign("resigns", data1, where)
				}
				message = fmt.Sprintf("Data %s Berhasil di resign kan", data.Number_of_employees)
				break
			}
		}

		result := map[string]interface{}{
			"code":    200,
			"message": message,
		}

		resp, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.Write([]byte(resp))
		return
	}

}

func Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nik, _ := strconv.Atoi(vars["number_of_employees"])
	ktp, _ := strconv.Atoi(vars["national_id"])

	var db, err = models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = models.Employee{}
	number_of_employees := nik
	national_id := ktp
	err = db.
		QueryRow("select number_of_employees, national_id from employees where number_of_employees = ? AND national_id = ?", number_of_employees, national_id).
		Scan(&result.Number_of_employees, &result.National_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))
}

func GetKaryawan(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nik, _ := strconv.Atoi(vars["number_of_employees"])
	ktp, _ := strconv.Atoi(vars["national_id"])

	var db, err = models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var Count_id int

	number_of_employees := nik
	national_id := ktp

	err = db.
		QueryRow("select count(id) as sum from employees where number_of_employees = ? AND national_id = ?", number_of_employees, national_id).
		Scan(&Count_id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var statushttp string

	if Count_id > 0 {
		statushttp = "200"
	} else {
		statushttp = "405"
	}

	response := map[string]interface{}{
		"status": statushttp,
	}

	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resp))
}

func GetAlamat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	numner_of_employees, _ := strconv.Atoi(vars["number_of_employees"])

	var db, err = models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var Place_of_birth,
		Date_of_birth,
		Address_jalan,
		Address_rt,
		Address_rw,
		Address_village,
		Address_district,
		Address_city,
		Address_province string

	err = db.
		QueryRow("SELECT COALESCE(place_of_birth, '') as place_of_birth, COALESCE(date_of_birth, '') as date_of_birth, COALESCE(address_jalan, '') as address_jalan, COALESCE(address_rt, '') as address_rt, COALESCE(address_rw, '') as address_rw, COALESCE(address_village, '') as address_village, COALESCE(address_district, '') as address_district, COALESCE(address_city, '') as address_city, COALESCE(address_province, '') as address_province FROM employees WHERE number_of_employees = ? ", numner_of_employees).
		Scan(&Place_of_birth, &Date_of_birth, &Address_jalan, &Address_rt, &Address_rw, &Address_village, &Address_district, &Address_city, &Address_province)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := map[string]interface{}{
		"place_of_birth":   Place_of_birth,
		"date_of_birth":    Date_of_birth,
		"address_jalan":    Address_jalan,
		"address_rt":       Address_rt,
		"address_rw":       Address_rw,
		"address_village":  Address_village,
		"address_district": Address_district,
		"address_city":     Address_city,
		"address_province": Address_province,
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))
}

func GetResign(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nik, _ := strconv.Atoi(vars["number_of_employees"])
	ktp, _ := strconv.Atoi(vars["national_id"])

	var db, err = models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbhwi, _ = models.ConnResign()
	defer dbhwi.Close()

	var resultemployee = models.Employee{}

	number_of_employees := nik
	national_id := ktp

	var Count_resigns int

	// Menampilkan data karyawan
	err = db.QueryRow("SELECT name, date_of_birth, hire_date, COALESCE(date_out, '0000-00-00') as date_out, status_employee FROM employees WHERE number_of_employees = ? AND national_id = ?", number_of_employees, national_id).
		Scan(&resultemployee.Name, &resultemployee.Date_of_birth, &resultemployee.Hire_date, &resultemployee.Date_out, &resultemployee.Status_employee)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Cari data resign pada pada database hwi berdasarkan nik
	err = dbhwi.QueryRow("SELECT count(id) as count_resign FROM resigns WHERE number_of_employees = ?", number_of_employees).
		Scan(&Count_resigns)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var Count_resign_submissions int
	var Count_status_resign_submissions string

	// Cari data pengajuan resign pada  database HWI
	err = dbhwi.QueryRow("SELECT count(id) as count_resign_submissions, COALESCE(status_resignsubmisssion, 'NULL') as  status_resignsubmisssion FROM resignation_submissions WHERE number_of_employees = ?", number_of_employees).
		Scan(&Count_resign_submissions, &Count_status_resign_submissions)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Card_id := strconv.Itoa(number_of_employees)

	dataemployee := map[string]interface{}{
		"status_employee": resultemployee.Status_employee,
		"date_out":        resultemployee.Date_out,
		"date_of_birth":   resultemployee.Date_of_birth,
		"hire_date":       resultemployee.Hire_date,
		"name":            resultemployee.Name,
		"job_level":       JobDepartment(Card_id)[0],
		"department":      JobDepartment(Card_id)[1],
	}

	//Jika sudah resign dan sudah mengajukan dan pegajuannya tidak cancel maka tidak dapat mengajukan lagi
	if Count_resigns > 0 && Count_resign_submissions > 0 && Count_status_resign_submissions != "cancel" {
		result := map[string]interface{}{
			"status":      405, //tidak diijinkan mengajukan
			"information": "sudah resign dan sudah mengajukan resign sehingga tidak dapat mengajukan resign lagi untuk mengambil parklaring anda dapat langsung ke hrd",
			"employee":    dataemployee,
		}
		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}

	//Jika belum resign dan sudah mengajukan dan pegajuannya tidak cancel maka tidak dapat mengajukan lagi
	if Count_resigns == 0 && Count_resign_submissions > 0 && Count_status_resign_submissions != "cancel" {
		result := map[string]interface{}{
			"status":      404, //tidak diijinkan mengajukan
			"information": "sudah mengajukan resign sehingga tidak dapat mengajukan resign lagi untuk mengambil parklaring anda dapat langsung ke hrd",
			"employee":    dataemployee,
		}
		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}

	//Sudah resign tapi belum mengajukan resign maka boleh mengajukan
	if (Count_resigns == 1 && Count_resign_submissions == 0) || resultemployee.Status_employee == "notactive" {
		result := map[string]interface{}{
			"status":      202, //boleh mengajukan resign walau sudah resign
			"information": "sudah resign dan belum mengajukan resign sehingga dapat mengajukan resign untuk mengambil parklaring anda dapat diambil 2 minggu dari tanggal pengajuan ini",
			"employee":    dataemployee,
		}
		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}

	// TIdak resign dan belum mengajukan resign maka boleh mengajukan resign
	if Count_resigns == 0 && Count_resign_submissions == 0 {
		result := map[string]interface{}{
			"status":      200, //boleh mengajukan resign karena belum resign
			"information": "silahkan isi data anda dengan benar untuk pengajuan resign",
			"employee":    dataemployee,
		}
		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}

	result := map[string]interface{}{
		"status":      200, //boleh mengajukan resign karena belum resign
		"information": "silahkan isi data anda dengan benar untuk pengajuan resign",
		"employee":    dataemployee,
	}
	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))

}

func GetJobs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT job_level FROM jobs WHERE level > ? and job_level != ? ORDER BY level DESC ", 8, "NONE")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []models.Job

	for rows.Next() {
		var each = models.Job{}
		var err = rows.Scan(&each.Job_level)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resp))
}

func GetDepartments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT department FROM departments WHERE department != 'HR GA SEA' and 	department != 'PLANT E' and 	department != 'MATL PUR' and 	department != 'CHANGE' and 	department != 'PLANT C' and 	department != 'PLANT H' and 	department != 'PLANT D' and 	department != 'PLANT E' and 	department != 'BUSINESS' and 	department != 'PLANT A AND B' and 	department != 'CHEMICAL ADVISOR' and department != 'EXECUTIVE SR. DIR HWI 2' and department != 'NONE' ORDER BY department ASC")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []models.Department

	for rows.Next() {
		var each = models.Department{}
		var err = rows.Scan(&each.Department)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resp))
}

func GetGedungs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := []string{
		"GEDUNG A", "GEDUNG B", "GEDUNG C", "GEDUNG D", "GEDUNG E", "GEDUNG F", "GEDUNG G", "GEDUNG H", "LAMINATING", "GUDANG SETTING", "WAREHOUSE (MATERIAL)", "SABLON", "EMBOSS", "TRAINING CENTER", "MAIN OFFICE", "EPTE (BEACUKAI)", "POS SECURITY", "KANTOR SERIKAT", "MESS / DRIVER",
	}
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))

}
func JobDepartment(number_of_employees string) []string {
	var db, err = models.ConnHrd()
	defer db.Close()

	Employee := models.Employee{}
	err = db.QueryRow("SELECT job_id, department_id FROM employees WHERE number_of_employees = ?", number_of_employees).
		Scan(&Employee.Job_id, &Employee.Department_id)
	if err != nil {
		fmt.Println(err.Error())
	}
	var Job_level, Department string
	err = db.QueryRow("SELECT job_level FROM jobs WHERE id = ?", Employee.Job_id).Scan(&Job_level)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.QueryRow("SELECT department FROM departments WHERE id = ?", Employee.Department_id).Scan(&Department)
	if err != nil {
		fmt.Println(err.Error())
	}
	var output = []string{Job_level, Department}
	return output
}
