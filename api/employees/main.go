package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	//penamaan Camel Cas untuk Import Package supaya bisa di pakai dari luar
	Number_of_employees string `json:"number_of_employees"` // `` cara membuat penamaan ulang pada golang pada saat di GET
	National_id         string `json:"national_id"`
}

type Dateemployee struct {
	Name            string `json:"name"`
	Date_of_birth   string `json:"date_of_birth"`
	Hire_date       string `json:"hire_date"`
	Date_out        string `json:"date_out"`
	Status_employee string `json:"status_employee"`
}

type Count struct {
	sum int
}

type Job struct {
	Job_level string `json:"job_level"`
}

type Department struct {
	Department string `json:"department"`
}

type ResignSubmission struct {
	Number_of_employees          string `json:"number_of_employees"`
	Name                         string `json:"name"`
	Position                     string `json:"position"`
	Department                   string `json:"department"`
	Building                     string `json:"building"`
	Hire_date                    string `json:"hire_date"`
	Date_out                     string `json:"date_out"`
	Date_resignation_submissions string `json:"date_resignation_submissions"`
	Type                         string `json:"type"`
	Reason                       string `json:"reason"`
	Detail_reason                string `json:"detail_reason"`
	Suggestion                   string `json:"suggestion"`
	Periode_of_service           int    `json:"periode_of_service"`
	Status_resignsubmisssion     string `json:"status_resignsubmisssion"`
	Age                          int    `json:"age"`
	Using_media                  string `json:"using_media"`
	Classification               string `json:"classification"`
	Created_at                   string `json:"created_at"`
	Updated_at                   string `json:"updated_at"`
}

type ResignAcc struct {
	Number_of_employees string `json:"number_of_employees"`
	Status_resign       string `json:"status_resign"`
}

type Resign struct {
	Id string `json:"number_of_employees"`
Number_of_employees string `json:"number_of_employees"`
Name string `json:"number_of_employees"`
Hire_date date 
Classification string `json:"number_of_employees"` 
Date_out date 
Date_resignsubmissions string `json:"number_of_employees"` 
Periode_of_service string `json:"number_of_employees"` 
Type string `json:"number_of_employees"`
Age string `json:"number_of_employees"` 
Status_resign string `json:"number_of_employees"`
Printed string `json:"number_of_employees"` 
Created_at string `json:"number_of_employees"` 
Updated_at string `json:"number_of_employees"`
}

func Conn() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrdit")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnHwi() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hwi")
	if err != nil {
		return nil, err
	}
	return db, nil
}

var datas []Employee

/*

	func Index(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		db, err := Conn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		rows, err := db.Query("select number_of_employees, COALESCE(national_id, 'NULL') as national_id from employees where id > ?", 0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		var result []Employee

		for rows.Next() {
			var each = Employee{}
			var err = rows.Scan(&each.Number_of_employees, &each.National_id)

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

*/

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

func Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	nik, _ := strconv.Atoi(vars["number_of_employees"])
	ktp, _ := strconv.Atoi(vars["national_id"])

	var db, err = Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = Employee{}
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

	var db, err = Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = Count{}

	number_of_employees := nik
	national_id := ktp

	err = db.
		QueryRow("select count(id) as sum from employees where number_of_employees = ? AND national_id = ?", number_of_employees, national_id).
		Scan(&result.sum)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var statushttp string

	if result.sum > 0 {
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

	var db, err = Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var Place_of_birth string
	var Date_of_birth string
	var Address_jalan string
	var Address_rt string
	var Address_rw string
	var Address_village string
	var Address_district string
	var Address_city string
	var Address_province string

	err = db.
		QueryRow("select  COALESCE(place_of_birth, '') as place_of_birth, COALESCE(date_of_birth, '') as date_of_birth, COALESCE(address_jalan, '') as address_jalan, COALESCE(address_rt, '') as address_rt, COALESCE(address_rw, '') as address_rw, COALESCE(address_village, '') as address_village, COALESCE(address_district, '') as address_district, COALESCE(address_city, '') as address_city, COALESCE(address_province, '') as address_province from employees where number_of_employees = ? ", numner_of_employees).
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

	var db, err = Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbhwi, _ = ConnHwi()
	defer dbhwi.Close()

	var resultemployee = Dateemployee{}

	number_of_employees := nik
	national_id := ktp

	var Count_resigns int

	err = db.
		QueryRow("select name, date_of_birth, hire_date, COALESCE(date_out, '0000-00-00') as date_out, status_employee from employees where number_of_employees = ? AND national_id = ?", number_of_employees, national_id).
		Scan(&resultemployee.Name, &resultemployee.Date_of_birth, &resultemployee.Hire_date, &resultemployee.Date_out, &resultemployee.Status_employee)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Cari data resign pada pada database hwi berdasarkan nik
	err = dbhwi.QueryRow("select count(id) as count_resign from resigns where number_of_employees = ?", number_of_employees).
		Scan(&Count_resigns)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var Count_resign_submissions int
	var Count_status_resign_submissions string

	// Cari data pengajuan resign pada  database HWI
	err = dbhwi.QueryRow("select count(id) as count_resign_submissions, COALESCE(status_resignsubmisssion, 'NULL') as  status_resignsubmisssion from resignation_submissions where number_of_employees = ?", number_of_employees).
		Scan(&Count_resign_submissions, &Count_status_resign_submissions)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if Count_resigns > 0 && Count_resign_submissions > 0 && Count_status_resign_submissions != "cancel" {
		result := map[string]interface{}{
			"status":      405, //tidak diijinkan mengajukan
			"information": "sudah resign dan sudah mengajukan resign sehingga tidak dapat mengajukan resign lagi untuk mengambil parklaring anda dapat langsung ke hrd",
			"employee":    resultemployee,
		}
		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}

	if Count_resigns == 0 && Count_resign_submissions > 0 && Count_status_resign_submissions != "cancel" {
		result := map[string]interface{}{
			"status":      404, //tidak diijinkan mengajukan
			"information": "sudah mengajukan resign sehingga tidak dapat mengajukan resign lagi untuk mengambil parklaring anda dapat langsung ke hrd",
			"employee":    resultemployee,
		}
		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}

	//Sudah resign tapi belum mengajukan resign
	if Count_resigns == 1 && Count_resign_submissions == 0 {
		result := map[string]interface{}{
			"status":      202, //boleh mengajukan resign walau sudah resign
			"information": "sudah resign dan belum mengajukan resign sehingga dapat mengajukan resign untuk mengambil parklaring anda dapat diambil 2 minggu dari tanggal pengajuan ini",
			"employee":    resultemployee,
		}
		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}

	// TIdak resign dan belum mengajukan resign
	if Count_resigns == 0 && Count_resign_submissions == 0 {
		result := map[string]interface{}{
			"status":      200, //boleh mengajukan resign karena belum resign
			"information": "silahkan isi data anda dengan benar untuk pengajuan resign",
			"employee":    resultemployee,
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
		"employee":    resultemployee,
	}
	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))

}

func GetJobs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db, err := Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select job_level from jobs where level > ? and job_level != ? order by level desc ", 8, "NONE")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Job

	for rows.Next() {
		var each = Job{}
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

	db, err := Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT department FROM departments WHERE department != 'HR GA SEA' and 	department != 'PLANT E' and 	department != 'MATL PUR' and 	department != 'CHANGE' and 	department != 'PLANT C' and 	department != 'PLANT H' and 	department != 'PLANT D' and 	department != 'PLANT E' and 	department != 'BUSINESS' and 	department != 'PLANT A AND B' and 	department != 'CHEMICAL ADVISOR' and department != 'EXECUTIVE SR. DIR HWI 2' and department != 'NONE' order by department asc")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Department

	for rows.Next() {
		var each = Department{}
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

func GetResignSubmission(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var count_submission int
	err = db.QueryRow("SELECT COUNT(id) as count_submission FROM resignation_submissions").
		Scan(&count_submission)

	if count_submission == 0 {

		var result = []map[string]string{
			{"number_of_employees": "NULL", "name": "NULL", "created_at": "NULL", "date_resignation_submissions": "NULL", "position": "NULL", "department": "NULL", "status_resignsubmisssion": "NULL"},
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(resp))
		return
	}

	rows, err := db.Query("SELECT number_of_employees, name, position, department, building, hire_date, COALESCE(date_out, '0000-00-00') as date_out, COALESCE(date_resignation_submissions, '0000-00-00'), COALESCE(type, ''), COALESCE(reason, ''), COALESCE(detail_reason, ''), COALESCE(periode_of_service, 0), COALESCE(age, 0), COALESCE(suggestion, ''), COALESCE(status_resignsubmisssion, ''), COALESCE(using_media, ''), created_at, updated_at FROM resignation_submissions order by created_at desc")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []ResignSubmission

	for rows.Next() {
		var each = ResignSubmission{}
		var err = rows.Scan(&each.Number_of_employees, &each.Name, &each.Position, &each.Department, &each.Building, &each.Hire_date, &each.Date_out, &each.Date_resignation_submissions, &each.Type, &each.Reason, &each.Detail_reason, &each.Periode_of_service, &each.Age, &each.Suggestion, &each.Status_resignsubmisssion, &each.Using_media, &each.Created_at, &each.Updated_at)

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

func GetResignSubmissionSearch(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	s, _ := vars["search"]

	search := s

	where := " "
	where = fmt.Sprintf(" %s WHERE name LIKE '%%%s%%' OR number_of_employees LIKE '%%%s%%' ", where, search, search)

	db, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT number_of_employees, name, position, department, building, hire_date, date_out, date_resignation_submissions, type, reason, detail_reason, periode_of_service, age, suggestion, status_resignsubmisssion, using_media, created_at, updated_at FROM resignation_submissions %s order by created_at desc", where)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []ResignSubmission

	for rows.Next() {
		var each = ResignSubmission{}
		var err = rows.Scan(&each.Number_of_employees, &each.Name, &each.Position, &each.Department, &each.Building, &each.Hire_date, &each.Date_out, &each.Date_resignation_submissions, &each.Type, &each.Reason, &each.Detail_reason, &each.Periode_of_service, &each.Age, &each.Suggestion, &each.Status_resignsubmisssion, &each.Using_media, &each.Created_at, &each.Updated_at)

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

func Post(w http.ResponseWriter, r *http.Request) {
	//untuk membuat json pertama kita harus set Header
	w.Header().Set("Content-Type", "application/json")

	// data := r.Body
	//mendecode requset body langsung menjadi json
	data := Employee{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(data)

	db, err := Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = Employee{}

	err = db.
		QueryRow("select number_of_employees, national_id FROM employees where number_of_employees = ? and national_id = ?", data.Number_of_employees, data.National_id).
		Scan(&result.Number_of_employees, &result.National_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resp))

}

func GetResignSubmissionStatus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	// nik, _ := strconv.Atoi(vars["number_of_employees"])
	// status_api, _ := strconv.Atoi(vars["status_resign"])
	number_of_employees := vars["number_of_employees"]
	status_resign := vars["status_resign"]

	var db, err = Conn()
	defer db.Close()

	var dbhwi, _ = ConnHwi()
	defer dbhwi.Close()

	var statushttp string

	if status_resign == "acc" {

		// Update data resign status acc
		_, err = dbhwi.Exec("UPDATE `resignation_submissions` SET `status_resignsubmisssion`= ? WHERE number_of_employees = ? ", status_resign, number_of_employees)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Acc Resign !")

		// Update data employees
		_, err = db.Exec("update employees set `status_employee` = ? where number_of_employees = ? ", "notactive", number_of_employees)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Update Employee Resign !")

		// // Insert data parklaring
		// insert_parklaring := fmt.Sprintf("insert into  set status_employee = %s where number_of_employees = %s ", "notactive", number_of_employees)

		// _, err = db.Exec(insert_parklaring)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }
		// fmt.Println("Insert Parklaring Resign !")

	} else if status_resign == "cancel" {

		// Update data resign status acc
		cancel_resign := fmt.Sprintf("update resign_submissions set status_resignsubmissions = %s where number_of_employees = %s ", status_resign, number_of_employees)
		_, err = dbhwi.Exec(cancel_resign)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Update Employee Resign !")

		// Update data employees
		active_employee := fmt.Sprintf("update employee set status_employee = %s where number_of_employees = %s ", "notactive", number_of_employees)
		_, err = db.Exec(active_employee)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Update Employee Resign !")
	}

	statushttp = status_resign
	response := map[string]interface{}{
		"status": statushttp,
	}

	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resp))
}

func PostAcc(w http.ResponseWriter, r *http.Request) {

	currentTime := time.Now()

	timestampnow := currentTime.Format("2006-01-02")
	datenow := currentTime.Format("2006-01-02")

	//untuk membuat json pertama kita harus set Header
	w.Header().Set("Content-Type", "application/json")

	// data := r.Body
	//mendecode requset body langsung menjadi json
	data := ResignAcc{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(data)

	db, err := Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbhwi, _ = ConnHwi()
	defer dbhwi.Close()

	//acc submission
	// Update data resign status acc
	_, err = dbhwi.Exec("UPDATE `resignation_submissions` SET `status_resignsubmisssion`= ? WHERE number_of_employees = ? ", data.Status_resign, data.Number_of_employees)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Acc Resign !")

	var status_employee string
	if data.Status_resign == "acc" {
		status_employee = "notactive"
	} else {
		status_employee = "active"
	}

	//edit employees
	_, err = db.Exec("UPDATE `employees` SET `status_employee`= ? WHERE number_of_employees = ? ", status_employee, data.Number_of_employees)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Acc Submission !")

	//select resign submissions

	var Submission = ResignSubmission{}

	err = dbhwi.QueryRow("SELECT number_of_employees,	name,	position,	department,	building,	hire_date,	COALESCE(date_out, '0000-00-00') as date_out,	date_resignation_submissions,	type,	reason,	detail_reason,	periode_of_service,	age,	suggestion,	status_resignsubmisssion,	using_media,	classification FROM resignation_submissions WHERE number_of_employees = ?", data.Number_of_employees).
		Scan(&Submission.Number_of_employees, &Submission.Name, &Submission.Position, &Submission.Department, &Submission.Building, &Submission.Hire_date, &Submission.Date_out, &Submission.Date_resignation_submissions, &Submission.Type, &Submission.Reason, &Submission.Detail_reason, &Submission.Suggestion, &Submission.Periode_of_service, &Submission.Age, &Submission.Status_resignsubmisssion, &Submission.Using_media, &Submission.Classification)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//insert data resigns
	_, err = dbhwi.Exec("INSERT INTO `resigns`(	`number_of_employees`,`name`, `hire_date`, `classification`, `date_out`, `date_resignsubmissions`, `periode_of_service`, `type`, `age`, `status_resign`, `printed`, `created_at`, `updated_at`) VALUES (?,	?,	?,	?,	?,	?,	?,	?,	?,	?,	?, ? , ?)", &Submission.Number_of_employees, &Submission.Name, &Submission.Hire_date, &Submission.Classification, &Submission.Date_out, &Submission.Date_resignation_submissions, &Submission.Periode_of_service, &Submission.Type, &Submission.Age, &Submission.Status_resignsubmisssion, 0, timestampnow, timestampnow)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Insert Data Resign !")

	var resign_id int

	err = dbhwi.QueryRow("SELECT id FROM resigns WHERE created_at = ?", timestampnow).
		Scan(&resign_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// insert data certificat atau description work
	if Submission.Type == "true" && Submission.Periode_of_service > 365 && data.Status_resign == "acc" {

		_, err = dbhwi.Exec("INSERT INTO `certificate_of_employments`(`resign_id`, `date_certificate_employee`, `no_certificate_employee`, `created_at`, `updated_at`) VALUES (?,?,?,?,?)", resign_id, datenow, 1, timestampnow, timestampnow)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Insert Data Certificate !")

	} else if Submission.Type == "false" && data.Status_resign == "acc" {

		_, err = dbhwi.Exec("INSERT INTO `work_experience_letters`(`resign_id`, `date_letter_exprerience`, `no_letter_experience`, `created_at`, `updated_at`) VALUES (?,?,?,?,?)", resign_id, datenow, 1, timestampnow, timestampnow)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Insert Data Experience !")

	}

	response := map[string]interface{}{
		"status_resign": data.Status_resign,
		"status_code":   200,
	}

	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resp))

}

func GetResigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var count_resign int

	err = db.QueryRow("SELECT COUNT(id) as count_resign FROM resigns ").
		Scan(&count_resign)
	if count_resign == 0 {
		var result = []map[string]string{
			{"number_of_employees": "NULL", "name": "NULL", "created_at": "NULL", "date_resignation_submissions": "NULL", "position": "NULL", "department": "NULL", "status_resignsubmisssion": "NULL"},
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(resp))
		return
	}

	rows, err := db.Query("SELECT * FROM resign")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Resign

	for rows.Next(){
		var each Resign{}
		var err = rows.Scan(&each)
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

/*
	func Update(w http.ResponseWriter, r *http.Request) {
		//untuk membuat json pertama kita harus set Header
		w.Header().Set("Content-Type", "application/json")

		//mendecode requset body langsung menjadi json
		data := Employee{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		db, err := Conn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		_, err = db.Exec("update employees set name = ? where id = ?", data.Name, data.Id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("insert success!")

		response := map[string]interface{}{
			"status": "Oke",
		}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	func Delete(w http.ResponseWriter, r *http.Request) {
		//untuk membuat json pertama kita harus set Header
		w.Header().Set("Content-Type", "application/json")

		//mendecode requset body langsung menjadi json
		data := Employee{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		db, err := Conn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		_, err = db.Exec("delete from user where id = ?", data.Id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Delete success!")

		response := map[string]interface{}{
			"status": "Oke",
		}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
*/

func main() {

	r := mux.NewRouter()

	//Submissions
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/employees/{number_of_employees}", Get).Methods("GET")
	r.HandleFunc("/resign/{number_of_employees}/{national_id}", GetKaryawan).Methods("GET")
	r.HandleFunc("/resigndate/{number_of_employees}/{national_id}", GetResign).Methods("GET")
	r.HandleFunc("/resignjobs", GetJobs).Methods("GET")
	r.HandleFunc("/resigndepartments", GetDepartments).Methods("GET")
	r.HandleFunc("/resignbuildings", GetGedungs).Methods("GET")
	r.HandleFunc("/resignalamat/{number_of_employees}", GetAlamat).Methods("GET")

	r.HandleFunc("/resignsubmissions/{search}", GetResignSubmissionSearch).Methods("GET")
	r.HandleFunc("/resignsubmissions/{number_of_employees}/{status_resign}", GetResignSubmissionStatus).Methods("GET")
	r.HandleFunc("/resignsubmissions", GetResignSubmission).Methods("GET")

	r.HandleFunc("/resign", Post).Methods("POST")
	r.HandleFunc("/resignacc", PostAcc).Methods("POST")

	//Resigns
	r.HandleFunc("/resigns", GetResigns).Methods("GET")

	// r.HandleFunc("/user/{id}", Update).Methods("PUT")
	// r.HandleFunc("/user/{id}", Delete).Methods("DELETE")

	fmt.Println("LIsten on Port 10.10.42.6:8880")
	http.ListenAndServe("10.10.42.6:8880", r)

}