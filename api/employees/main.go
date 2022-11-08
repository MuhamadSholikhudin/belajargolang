package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/theTardigrade/age"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	//penamaan Camel Cas untuk Import Package supaya bisa di pakai dari luar
	Number_of_employees string `json:"number_of_employees"` // `` cara membuat penamaan ulang pada golang pada saat di GET
	National_id         string `json:"national_id"`
	Name                string `json:"name"`
	Job_id              string `json:"job_id"`
	Department_id       string `json:"department_id"`
	Date_of_birth       string `json:"date_of_birth"`
	Hire_date           string `json:"hire_date"`
	Date_out            string `json:"date_out"`
	Place_of_birth      string `json:"place_of_birth"`
	Address_jalan       string `json:"address_jalan"`
	Address_rt          string `json:"address_rt"`
	Address_rw          string `json:"address_rw"`
	Address_village     string `json:"address_village"`
	Address_district    string `json:"address_district"`
	Address_city        string `json:"address_city"`
	Address_province    string `json:"address_province"`
	Status_employee     string `json:"status_employee"`
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
	Id                     string `json:"id"`
	Number_of_employees    string `json:"number_of_employees"`
	Name                   string `json:"name"`
	Position               string `json:"position"`
	Department             string `json:"department"`
	Hire_date              string `json:"hire_date"`
	Classification         string `json:"classification"`
	Date_out               string `json:"date_out"`
	Date_resignsubmissions string `json:"date_resignsubmissions"`
	Periode_of_service     int    `json:"periode_of_service"`
	Type                   string `json:"type"`
	Age                    int    `json:"age"`
	Status_resign          string `json:"status_resign"`
	Printed                string `json:"printed"`
	Created_at             string `json:"created_at"`
	Updated_at             string `json:"updated_at"`
}

type Letter struct {
	Id                  int    `json:"id"`
	Resign_id           int    `json:"resign_id"`
	Number_of_employees string `json:"number_of_employees"`
	Date                string `json:"date"`
	No                  string `json:"no"`
	Rom                 string `json:"rom"`
	Status              string `json:"status"`
	Action              string `json:"action"`
	Created_at          string `json:"created_at"`
	Updated_at          string `json:"updated_at"`
}

const (
	LINKFRONTEND string = "http://127.0.0.1:8000"

	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
	DDMMYYYY       = "2006-01-02"
)

func DMYhms() string {
	t := time.Now()
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println(err)
	}
	datetimenow := t.In(location).Format(DDMMYYYYhhmmss)
	return datetimenow
}

func DMY() string {
	t := time.Now()
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println(err)
	}
	datetimenow := t.In(location).Format(DDMMYYYY)
	return datetimenow
}

func StringMonth() string {
	var now = time.Now()
	var stringmonth string
	stringmonth = strconv.Itoa(int(now.Month()))
	var length = len([]rune(stringmonth))
	var value string = stringmonth
	if length == 1 {
		value = fmt.Sprintf("0%s", stringmonth)
	}
	return value
}

func Rom(stringmonth string) string {
	var Rom string
	if stringmonth == "01" {
		Rom = "I"
	}
	if stringmonth == "02" {
		Rom = "II"
	}
	if stringmonth == "03" {
		Rom = "III"
	}
	if stringmonth == "04" {
		Rom = "IV"
	}
	if stringmonth == "05" {
		Rom = "V"
	}
	if stringmonth == "06" {
		Rom = "VI"
	}
	if stringmonth == "07" {
		Rom = "VII"
	}
	if stringmonth == "08" {
		Rom = "VIII"
	}
	if stringmonth == "09" {
		Rom = "IX"
	}
	if stringmonth == "10" {
		Rom = "X"
	}
	if stringmonth == "11" {
		Rom = "XI"
	}
	if stringmonth == "12" {
		Rom = "XII"
	}
	return Rom
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
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

	fmt.Println(StringMonth())

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

	var resultemployee = Employee{}

	number_of_employees := nik
	national_id := ktp

	var Count_resigns int

	// Menampilkan data karyawan
	err = db.QueryRow("select name, date_of_birth, hire_date, COALESCE(date_out, '0000-00-00') as date_out, status_employee from employees where number_of_employees = ? AND national_id = ?", number_of_employees, national_id).
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

	//Jika sudah resign dan sudah mengajukan dan pegajuannya tidak cancel maka tidak dapat mengajukan lagi
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

	//Jika belum resign dan sudah mengajukan dan pegajuannya tidak cancel maka tidak dapat mengajukan lagi
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

	//Sudah resign tapi belum mengajukan resign maka boleh mengajukan
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

	// TIdak resign dan belum mengajukan resign maka boleh mengajukan resign
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

		var datanull = []map[string]string{
			{"number_of_employees": "NULL", "name": "NULL", "created_at": "NULL", "date_resignation_submissions": "NULL", "position": "NULL", "department": "NULL", "status_resignsubmisssion": "NULL"},
		}

		result := map[string]interface{}{
			"code":  404,
			"meta":  "NULL",
			"data":  datanull,
			"links": "NULL",
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(resp))
		return
	}

	u, err := url.Parse(r.RequestURI)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	var sqlPaging string = "SELECT number_of_employees, name, COALESCE(position, ''), COALESCE(department, ''), COALESCE(building, ''), COALESCE(hire_date, ''), COALESCE(date_out, '') as date_out, COALESCE(date_resignation_submissions, ''), COALESCE(type, ''), COALESCE(reason, ''), COALESCE(detail_reason, ''), COALESCE(periode_of_service, 0), COALESCE(age, 0), COALESCE(suggestion, ''), COALESCE(status_resignsubmisssion, ''), COALESCE(using_media, ''), COALESCE(classification, ''), COALESCE(created_at, ''), COALESCE(updated_at, '') FROM resignation_submissions"
	var sqlCount string = "SELECT COUNT(*) FROM resignation_submissions"

	var params string = ""

	number_of_employees, checkNumber_of_employees := q["number_of_employees"]
	if checkNumber_of_employees != false {
		justStringnumber_of_employees := strings.Join(number_of_employees, "")
		sqlPaging = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%' order by created_at desc", sqlPaging, justStringnumber_of_employees)
		sqlCount = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%'", sqlCount, justStringnumber_of_employees)
		params = fmt.Sprintf("&%snumber_of_employees=%s", params, justStringnumber_of_employees)
	}

	var total int64

	db.QueryRow(sqlCount).Scan(&total)

	if total == 0 {
		var datanull = []map[string]string{
			{"number_of_employees": "NULL", "name": "NULL", "created_at": "NULL", "date_resignation_submissions": "NULL", "position": "NULL", "department": "NULL", "status_resignsubmisssion": "NULL"},
		}

		result := map[string]interface{}{
			"code":  404,
			"meta":  "NULL",
			"data":  datanull,
			"links": "NULL",
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(resp))
		return
	}

	var totalminbyperpage int64 = total - ((total / 10) * 10)

	var lastPage int64
	if totalminbyperpage == 0 {
		lastPage = (total / 10)
	} else {
		lastPage = ((total / 10) + 1)
	}

	page, _ := strconv.Atoi("1")
	cpage, checkPage := q["page"]
	if checkPage != false {
		spage, _ := strconv.Atoi(strings.Join(cpage, ""))
		page = spage
	}

	var first, last, next, prev string
	first, last, next, prev = "", "", "", ""

	first = "1"
	last = strconv.Itoa(int(lastPage))

	next = strconv.Itoa(int(page + 1))
	if int(page+1) >= int(lastPage) {
		next = strconv.Itoa(int(lastPage))
	}

	prev = strconv.Itoa(int(page - 1))
	if int(page) == 1 {
		prev = strconv.Itoa(int(page))
	}

	perPage, _ := strconv.Atoi("10")
	sqlPaging = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlPaging, perPage, (page-1)*perPage)

	rows, err := db.Query(sqlPaging)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var submission []ResignSubmission

	for rows.Next() {
		var each = ResignSubmission{}
		var err = rows.Scan(&each.Number_of_employees, &each.Name, &each.Position, &each.Department, &each.Building, &each.Hire_date, &each.Date_out, &each.Date_resignation_submissions, &each.Type, &each.Reason, &each.Detail_reason, &each.Periode_of_service, &each.Age, &each.Suggestion, &each.Status_resignsubmisssion, &each.Using_media, &each.Classification, &each.Created_at, &each.Updated_at)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		submission = append(submission, each)
	}

	links := map[string]interface{}{
		"first": fmt.Sprintf("%s/resigns_submission?page=%s%s", LINKFRONTEND, first, params),
		"last":  fmt.Sprintf("%s/resigns_submission?page=%s%s", LINKFRONTEND, last, params),
		"next":  fmt.Sprintf("%s/resigns_submission?page=%s%s", LINKFRONTEND, next, params),
		"prev":  fmt.Sprintf("%s/resigns_submission?page=%s%s", LINKFRONTEND, prev, params),
	}

	informationpages := map[string]interface{}{
		"currentPage": page,
		"from":        ((page - 1) * 10) + 1,
		"lastPage":    lastPage,
		"perPage":     10,
		"to":          ((page - 1) * 10) + len(submission),
		"total":       total,
	}

	pages := map[string]interface{}{
		"page": informationpages,
	}

	result := map[string]interface{}{
		"code":  200,
		"meta":  pages,
		"data":  submission,
		"links": links,
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

func UploadSubmission(w http.ResponseWriter, r *http.Request) {

	db, err := Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbhwi, _ = ConnHwi()
	defer dbhwi.Close()

	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadedFile, _, err := r.FormFile("file")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	reader := csv.NewReader(uploadedFile)
	records, _ := reader.ReadAll()

	var notification []string
	notification = append(notification, "")
	var code int = 200

	for _, record := range records {

		var Count_employees int
		var Number_of_employees string
		var Status_employee string

		err = db.
			QueryRow("select COUNT(id), COALESCE(number_of_employees, 'NULL'), COALESCE(status_employee, 'NULL') FROM employees where number_of_employees = ? ", record[1]).
			Scan(&Count_employees, &Number_of_employees, &Status_employee)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		switch Count_employees {
		case 1: // Karyawan di temukan

			var resultemployee = Employee{}

			// Menampilkan data karyawan
			err = db.QueryRow("select name, COALESCE(date_of_birth, '0000-00-00'), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') as date_out, status_employee from employees where number_of_employees = ? ", Number_of_employees).
				Scan(&resultemployee.Name, &resultemployee.Date_of_birth, &resultemployee.Hire_date, &resultemployee.Date_out, &resultemployee.Status_employee)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var Count_resigns int
			// Cari data resign pada pada database hwi berdasarkan nik
			err = dbhwi.QueryRow("select count(id) as count_resign from resigns where number_of_employees = ? ", Number_of_employees).
				Scan(&Count_resigns)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var Count_resign_submissions int
			var Count_status_resign_submissions string

			// Cari data pengajuan resign pada  database HWI
			err = dbhwi.QueryRow("select count(id) as count_resign_submissions, COALESCE(status_resignsubmisssion, 'NULL') as  status_resignsubmisssion from resignation_submissions where number_of_employees = ? ", Number_of_employees).
				Scan(&Count_resign_submissions, &Count_status_resign_submissions)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// const day, month, year = 2, 1, 1999
			// date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			// dateAge := age.Calculate(date)

			// fmt.Println(dateAge)

			// ===================== RESULT

			if Count_resigns > 0 && Count_resign_submissions > 0 && Count_status_resign_submissions != "cancel" {
				//Jika sudah resign dan sudah mengajukan dan pegajuannya tidak cancel maka tidak dapat mengajukan lagi

				//fmt.Println("Pengajuan ", Number_of_employees, " tidak dapat di simpan karena sudah mengajukan dan sudah resign")

				each := fmt.Sprintf("NIK ini %s tidak dapat mengajukan resign karena sudah mengajukan resign dan status karyawan sudah resign. </br>", Number_of_employees)
				notification = append(notification, each)
				code = 400

			} else if Count_resigns == 0 && Count_resign_submissions > 0 && Count_status_resign_submissions != "cancel" {
				//Jika belum resign dan sudah mengajukan dan pegajuannya tidak cancel maka tidak dapat mengajukan lagi

				//fmt.Println("Pengajuan ", Number_of_employees, " tidak dapat di simpan karena sudah mengajukan dan statusnya menunggu")

				each := fmt.Sprintf("NIK ini %s tidak dapat mengajukan resign karena sudah resign dengan status pengajuan wait. </br>", Number_of_employees)
				notification = append(notification, each)
				code = 400

			} else if Count_resigns == 1 && Count_resign_submissions == 0 {
				//Sudah resign tapi belum mengajukan resign maka boleh mengajukan

				_, err = dbhwi.Exec("INSERT INTO `resignation_submissions` (`number_of_employees`, `name`, `position`, `department`, `building`, `hire_date`, `date_out`, `date_resignation_submissions`, `type`, `reason`, `detail_reason`, `periode_of_service`, `age`, `suggestion`, `status_resignsubmisssion`, `using_media`, `classification`, `print`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )", Number_of_employees, resultemployee.Name, record[5], record[9], record[9], resultemployee.Hire_date, resultemployee.Date_out, record[4], "false", record[3], record[6], Periode_of_serve(resultemployee.Hire_date, record[4]), Age(resultemployee.Date_of_birth), record[7], "wait", "google", "Mengajukan permohonan resign setelah karyawan resign", 0, record[0], record[0])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				//fmt.Println("INSERT DATA Submission type false")

				_, err = dbhwi.Exec("INSERT INTO `kuesioners`( `number_of_employees`, `k1`, `k2`, `k3`, `k4`, `k5`, `k6`, `k7`, `created_at`, `updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?) ", Number_of_employees, record[10], record[11], record[12], record[13], record[14], record[15], record[16], record[0], record[0])
				if err != nil {
					fmt.Println(err.Error())
					return
				}

			} else if Count_resigns == 0 && Count_resign_submissions == 0 {
				// TIdak resign dan belum mengajukan resign maka boleh mengajukan resign

				_, err = dbhwi.Exec("INSERT INTO `resignation_submissions` (`number_of_employees`, `name`, `position`, `department`, `building`, `hire_date`, `date_out`, `date_resignation_submissions`, `type`, `reason`, `detail_reason`, `periode_of_service`, `age`, `suggestion`, `status_resignsubmisssion`, `using_media`, `classification`, `print`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )", Number_of_employees, resultemployee.Name, record[5], record[9], record[9], resultemployee.Hire_date, resultemployee.Date_out, record[4], "true", record[3], record[6], Periode_of_serve(resultemployee.Hire_date, record[4]), Age(resultemployee.Date_of_birth), record[7], "wait", "google", "Mengajukan permohonan resign setelah karyawan resign", 0, record[0], record[0])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				//fmt.Println("INSERT DATA Submission type true")

				_, err = dbhwi.Exec("INSERT INTO `kuesioners`( `number_of_employees`, `k1`, `k2`, `k3`, `k4`, `k5`, `k6`, `k7`, `created_at`, `updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?) ", Number_of_employees, record[10], record[11], record[12], record[13], record[14], record[15], record[16], record[0], record[0])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
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

func GetEditSubmission(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	Number_of_employess, _ := strconv.Atoi(vars["number_of_employees"])

	var dbhwi, err = ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	var Submission ResignSubmission

	err = dbhwi.QueryRow("SELECT number_of_employees, COALESCE(name, ''), COALESCE(position, ''), COALESCE(department, ''), COALESCE(building, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), COALESCE(date_resignation_submissions, '0000-00-00'), COALESCE(type, ''), COALESCE(reason, ''), COALESCE(detail_reason, ''), COALESCE(suggestion, ''), COALESCE(periode_of_service, 0), COALESCE(status_resignsubmisssion, ''), COALESCE(age, 0), COALESCE(using_media, ''), COALESCE(classification, ''), COALESCE(created_at, '0000-00-00 00:00:00'), COALESCE(updated_at, '0000-00-00 00:00:00')  FROM resignation_submissions WHERE number_of_employees = ? ", Number_of_employess).
		Scan(&Submission.Number_of_employees, &Submission.Name, &Submission.Position, &Submission.Department, &Submission.Building, &Submission.Hire_date, &Submission.Date_out, &Submission.Date_resignation_submissions, &Submission.Type, &Submission.Reason, &Submission.Detail_reason, &Submission.Suggestion, &Submission.Periode_of_service, &Submission.Status_resignsubmisssion, &Submission.Age, &Submission.Using_media, &Submission.Classification, &Submission.Created_at, &Submission.Updated_at)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := map[string]interface{}{
		"data": Submission,
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))

}

func GetUpdateSubmission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data := ResignSubmission{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dbhwi, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	_, err = dbhwi.Exec("UPDATE `resignation_submissions` SET `name`= ? ,`position`= ? ,`department`=  ? , `hire_date`= ? ,`date_out`= ? ,`date_resignation_submissions`= ? ,`type`= ? ,`reason`= ? ,`detail_reason`= ? ,`periode_of_service`= ? ,`age`= ? ,`suggestion`= ? ,`status_resignsubmisssion`= ? ,`using_media`= ? ,`classification`= ? ,`created_at`= ? ,`updated_at`= ?  WHERE number_of_employees = ? ", data.Name, data.Position, data.Department, data.Hire_date, data.Date_out, data.Date_resignation_submissions, data.Type, data.Reason, data.Detail_reason, Periode_of_serve(data.Hire_date, data.Date_resignation_submissions), data.Age, data.Suggestion, data.Status_resignsubmisssion, data.Using_media, data.Classification, data.Created_at, data.Updated_at, data.Number_of_employees)
	if err != nil {

		result := map[string]interface{}{
			"code":    400,
			"message": "Update Loss",
		}
		resp, _ := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Println(400)
		w.Write([]byte(resp))
		return
	}

	result := map[string]interface{}{
		"code":    200,
		"data":    data,
		"message": "Update Success",
	}

	resp, _ := json.Marshal(result)
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
	_, err = dbhwi.Exec("INSERT INTO `resigns`(	`number_of_employees`,`name`, `hire_date`, `classification`, `date_out`, `date_resignsubmissions`, `periode_of_service`, `type`, `age`, `status_resign`, `printed`, `created_at`, `updated_at`) VALUES (?,	?,	?,	?,	?,	?,	?,	?,	?,	?,	?, ? , ?)", &Submission.Number_of_employees, &Submission.Name, &Submission.Hire_date, &Submission.Classification, &Submission.Date_resignation_submissions, &Submission.Date_resignation_submissions, &Submission.Periode_of_service, &Submission.Type, &Submission.Age, &Submission.Status_resignsubmisssion, 0, DMYhms(), DMYhms())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Insert Data Resign !")

	var resign_id int

	err = dbhwi.QueryRow("SELECT id FROM resigns WHERE created_at = ?", DMYhms()).
		Scan(&resign_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// insert data certificat atau description work
	if Submission.Type == "true" && Submission.Periode_of_service > 365 && data.Status_resign == "acc" {

		_, err = dbhwi.Exec("INSERT INTO `certificate_of_employments`(`resign_id`, `date_certificate_employee`, `no_certificate_employee`, `created_at`, `updated_at`) VALUES (?,?,?,?,?)", resign_id, DMY(), 1, DMYhms(), DMYhms())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Insert Data Certificate !")

	} else if Submission.Type == "false" && data.Status_resign == "acc" {

		_, err = dbhwi.Exec("INSERT INTO `work_experience_letters`(`resign_id`, `date_letter_exprerience`, `no_letter_experience`, `created_at`, `updated_at`) VALUES (?,?,?,?,?)", resign_id, DMY(), 1, DMYhms(), DMYhms())
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
		var datanull = []map[string]string{
			{"id": "NULL", "number_of_employees": "NULL", "name": "NULL", "hire_date": "NULL", "date_out": "NULL", "date_resignsubmissions": "NULL", "position": "NULL", "department": "NULL", "type": "NULL", "age": "0", "status_resign": "NULL", "printed": "NULL", "created_at": "NULL", "updated_at": "NULL"},
		}

		result := map[string]interface{}{
			"code":  404,
			"meta":  "NULL",
			"data":  datanull,
			"links": "NULL",
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(resp))
		return
	}
	u, err := url.Parse(r.RequestURI)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	var sqlPaging string = "SELECT id, number_of_employees, COALESCE(name, ''), COALESCE(hire_date, ''), COALESCE(classification, ''), COALESCE(date_out, ''), COALESCE(date_resignsubmissions, ''), COALESCE(periode_of_service, 0), COALESCE(position, ''), COALESCE(department, ''), COALESCE(type, ''), COALESCE(age, ''), COALESCE(status_resign, ''), COALESCE(printed, ''), COALESCE(created_at, ''), COALESCE(updated_at, '') FROM resigns"
	var sqlCount string = "SELECT COUNT(*) FROM resigns"
	var params string = ""

	number_of_employees, checkNumber_of_employees := q["number_of_employees"]
	if checkNumber_of_employees != false {
		justStringnumber_of_employees := strings.Join(number_of_employees, "")
		sqlPaging = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%'", sqlPaging, justStringnumber_of_employees)
		sqlCount = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%'", sqlCount, justStringnumber_of_employees)
		params = fmt.Sprintf("&%snumber_of_employees=%s", params, justStringnumber_of_employees)
	}

	var total int64
	db.QueryRow(sqlCount).Scan(&total)
	if total == 0 {
		var datanull = []map[string]string{
			{"id": "NULL", "number_of_employees": "NULL", "name": "NULL", "hire_date": "NULL", "date_out": "NULL", "date_resignsubmissions": "NULL", "position": "NULL", "department": "NULL", "type": "NULL", "age": "0", "status_resign": "NULL", "printed": "NULL", "created_at": "NULL", "updated_at": "NULL"},
		}

		result := map[string]interface{}{
			"code":  404,
			"meta":  "NULL",
			"data":  datanull,
			"links": "NULL",
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(resp))
		return
	}

	var totalminbyperpage, lastPage int64
	totalminbyperpage = total - ((total / 10) * 10)

	if totalminbyperpage == 0 {
		lastPage = (total / 10)
	} else {
		lastPage = ((total / 10) + 1)
	}

	page, _ := strconv.Atoi("1")
	cpage, checkPage := q["page"]
	if checkPage != false {
		spage, _ := strconv.Atoi(strings.Join(cpage, ""))
		page = spage
	}

	var first, last, next, prev string
	first, last, next, prev = "", "", "", ""

	first = "1"
	last = strconv.Itoa(int(lastPage))

	next = strconv.Itoa(int(page + 1))
	if int(page+1) >= int(lastPage) {
		next = strconv.Itoa(int(lastPage))
	}

	prev = strconv.Itoa(int(page - 1))
	if int(page) == 1 {
		prev = strconv.Itoa(int(page))
	}

	perPage, _ := strconv.Atoi("10")
	sqlPaging = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlPaging, perPage, (page-1)*perPage)

	rows, err := db.Query(sqlPaging)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var resign []Resign

	for rows.Next() {
		var each = Resign{}
		var err = rows.Scan(&each.Id, &each.Number_of_employees, &each.Name, &each.Hire_date, &each.Classification, &each.Date_out, &each.Date_resignsubmissions, &each.Periode_of_service, &each.Position, &each.Department, &each.Type, &each.Age, &each.Status_resign, &each.Printed, &each.Created_at, &each.Updated_at)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		resign = append(resign, each)
	}

	links := map[string]interface{}{
		"first": fmt.Sprintf("%s/resigns_resign?page=%s%s", LINKFRONTEND, first, params),
		"last":  fmt.Sprintf("%s/resigns_resign?page=%s%s", LINKFRONTEND, last, params),
		"next":  fmt.Sprintf("%s/resigns_resign?page=%s%s", LINKFRONTEND, next, params),
		"prev":  fmt.Sprintf("%s/resigns_resign?page=%s%s", LINKFRONTEND, prev, params),
	}

	informationpages := map[string]interface{}{
		"currentPage": page,
		"from":        ((page - 1) * 10) + 1,
		"lastPage":    lastPage,
		"perPage":     10,
		"to":          ((page - 1) * 10) + len(resign),
		"total":       total,
	}

	pages := map[string]interface{}{
		"page": informationpages,
	}

	result := map[string]interface{}{
		"code":  200,
		"meta":  pages,
		"data":  resign,
		"links": links,
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte(resp))

}

func GetEditResign(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	Number_of_employess, _ := strconv.Atoi(vars["number_of_employees"])

	var dbhwi, err = ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	var Resign Resign

	err = dbhwi.QueryRow("SELECT number_of_employees, COALESCE(name, ''), COALESCE(position, ''), COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), COALESCE(date_resignsubmissions, '0000-00-00'), COALESCE(type, ''), COALESCE(periode_of_service, 0), COALESCE(status_resign, ''), COALESCE(age, 0), COALESCE(classification, ''), COALESCE(created_at, '0000-00-00 00:00:00'), COALESCE(updated_at, '0000-00-00 00:00:00')  FROM resigns WHERE number_of_employees = ? ", Number_of_employess).
		Scan(&Resign.Number_of_employees, &Resign.Name, &Resign.Position, &Resign.Department, &Resign.Hire_date, &Resign.Date_out, &Resign.Date_resignsubmissions, &Resign.Type, &Resign.Periode_of_service, &Resign.Status_resign, &Resign.Age, &Resign.Classification, &Resign.Created_at, &Resign.Updated_at)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := map[string]interface{}{
		"data": Resign,
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))

}

func GetUpdateResign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data := Resign{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dbhwi, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	_, err = dbhwi.Exec("UPDATE `resigns` SET `name`= ? ,`position`= ? ,`department`=  ? , `hire_date`= ? ,`date_out`= ? ,`date_resignsubmissions`= ? ,`type`= ? , `periode_of_service`= ? ,`age`= ? ,`status_resign`= ? , `classification`= ? ,`created_at`= ? ,`updated_at`= ?  WHERE number_of_employees = ? ", data.Name, data.Position, data.Department, data.Hire_date, data.Date_out, data.Date_resignsubmissions, data.Type, Periode_of_serve(data.Hire_date, data.Date_resignsubmissions), data.Age, data.Status_resign, data.Classification, data.Created_at, data.Updated_at, data.Number_of_employees)
	if err != nil {
		fmt.Println(err.Error())
		result := map[string]interface{}{
			"code":    400,
			"message": "Update Loss",
		}
		resp, _ := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}

	result := map[string]interface{}{
		"code":    200,
		"data":    data,
		"message": "Update Success",
	}

	resp, _ := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))
}

func PostCertifcate(w http.ResponseWriter, r *http.Request) {

	//untuk membuat json pertama kita harus set Header
	w.Header().Set("Content-Type", "application/json")

	data := Letter{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var Resign_id int
	var ResignSel = Resign{}
	err = db.QueryRow("SELECT id as resign_id, name, COALESCE(position, ''), COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') FROM resigns WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&Resign_id, &ResignSel.Name, &ResignSel.Position, &ResignSel.Department, &ResignSel.Hire_date, &ResignSel.Date_out)
	if err != nil {
		fmt.Print(err.Error())
	}

	var CountCertifcateNumberOf_employees int

	err = db.QueryRow("SELECT COUNT(*) FROM certificate_of_employments WHERE number_of_employees = ?", data.Number_of_employees).
		Scan(&CountCertifcateNumberOf_employees)
	if err != nil {
		fmt.Print(err.Error())
	}

	if CountCertifcateNumberOf_employees == 0 {
		var CountCertificateByDate, CountNoCertificateEmployee int
		var yearstring string
		yearstring = strconv.Itoa(time.Now().Year())
		err = db.QueryRow("SELECT COUNT(id) as CountCertificateByDate, COALESCE(no_certificate_employee, 0) as no_certificate_employee FROM certificate_of_employments WHERE YEAR(date_certificate_employee) = ? AND MONTH(date_certificate_employee) = ? ORDER BY date_certificate_employee DESC", yearstring, StringMonth()).
			Scan(&CountCertificateByDate, &CountNoCertificateEmployee)
		if err != nil {
			fmt.Print(err.Error())
		}
		// create certificate
		_, err := db.Exec("INSERT INTO certificate_of_employments (resign_id, number_of_employees, date_certificate_employee, no_certificate_employee, rom, created_at, updated_at) VALUES (?,?,?,?,?,?,?)", Resign_id, data.Number_of_employees, DMY(), (CountNoCertificateEmployee + 1), Rom(StringMonth()), DMYhms(), DMYhms())
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	var certifictaeofemploment = Letter{}

	err = db.QueryRow("SELECT id, resign_id, number_of_employees, date_certificate_employee, no_certificate_employee, rom, created_at, updated_at FROM certificate_of_employments WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&certifictaeofemploment.Id, &certifictaeofemploment.Resign_id, &certifictaeofemploment.Number_of_employees, &certifictaeofemploment.Date, &certifictaeofemploment.No, &certifictaeofemploment.Rom, &certifictaeofemploment.Created_at, &certifictaeofemploment.Updated_at)

	if err != nil {
		fmt.Println(err.Error())
	}

	certificate := map[string]interface{}{
		"id":                  certifictaeofemploment.Id,
		"resign_id":           certifictaeofemploment.Resign_id,
		"number_of_employees": certifictaeofemploment.Number_of_employees,
		"name":                ResignSel.Name,
		"position":            ResignSel.Position,
		"department":          ResignSel.Department,
		"hire_date":           ResignSel.Hire_date,
		"date_out":            ResignSel.Date_out,
		"date":                certifictaeofemploment.Date,
		"no":                  certifictaeofemploment.No,
		"rom":                 certifictaeofemploment.Rom,
		"created_at":          certifictaeofemploment.Created_at,
		"updated_at":          certifictaeofemploment.Updated_at,
	}

	result := map[string]interface{}{
		"code":    200,
		"data":    certificate,
		"message": "Succesfully",
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resp))
}

func PostExperience(w http.ResponseWriter, r *http.Request) {

	//untuk membuat json pertama kita harus set Header
	w.Header().Set("Content-Type", "application/json")

	data := Letter{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var Resign_id int
	var ResignSel = Resign{}
	err = db.QueryRow("SELECT id as resign_id, name, COALESCE(position, ''), COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') FROM resigns WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&Resign_id, &ResignSel.Name, &ResignSel.Position, &ResignSel.Department, &ResignSel.Hire_date, &ResignSel.Date_out)
	if err != nil {
		fmt.Print(err.Error())
	}

	var CountCertifcateNumberOf_employees int

	err = db.QueryRow("SELECT COUNT(*) FROM work_experience_letters WHERE number_of_employees = ?", data.Number_of_employees).
		Scan(&CountCertifcateNumberOf_employees)
	if err != nil {
		fmt.Print(err.Error())
	}

	if CountCertifcateNumberOf_employees == 0 {
		var CountCertificateByDate, CountNoCertificateEmployee int
		var yearstring string
		yearstring = strconv.Itoa(time.Now().Year())
		err = db.QueryRow("SELECT COUNT(id) as CountCertificateByDate, COALESCE(no_letter_experience, 0) as no_letter_experience FROM work_experience_letters WHERE YEAR(date_certificate_employee) = ? AND MONTH(date_certificate_employee) = ? ORDER BY date_certificate_employee DESC", yearstring, StringMonth()).
			Scan(&CountCertificateByDate, &CountNoCertificateEmployee)
		if err != nil {
			fmt.Print(err.Error())
		}
		// create certificate
		_, err := db.Exec("INSERT INTO work_experience_letters (resign_id, number_of_employees, date_letter_exprerience, no_letter_experience, rom, created_at, updated_at) VALUES (?,?,?,?,?,?,?)", Resign_id, data.Number_of_employees, DMY(), (CountNoCertificateEmployee + 1), Rom(StringMonth()), DMYhms(), DMYhms())
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	var certifictaeofemploment = Letter{}

	err = db.QueryRow("SELECT id, resign_id, number_of_employees, date_letter_exprerience, no_letter_experience, rom, created_at, updated_at FROM work_experience_letters WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&certifictaeofemploment.Id, &certifictaeofemploment.Resign_id, &certifictaeofemploment.Number_of_employees, &certifictaeofemploment.Date, &certifictaeofemploment.No, &certifictaeofemploment.Rom, &certifictaeofemploment.Created_at, &certifictaeofemploment.Updated_at)

	if err != nil {
		fmt.Println(err.Error())
	}

	certificate := map[string]interface{}{
		"id":                  certifictaeofemploment.Id,
		"resign_id":           certifictaeofemploment.Resign_id,
		"number_of_employees": certifictaeofemploment.Number_of_employees,
		"name":                ResignSel.Name,
		"position":            ResignSel.Position,
		"department":          ResignSel.Department,
		"hire_date":           ResignSel.Hire_date,
		"date_out":            ResignSel.Date_out,
		"date":                certifictaeofemploment.Date,
		"no":                  certifictaeofemploment.No,
		"rom":                 certifictaeofemploment.Rom,
		"created_at":          certifictaeofemploment.Created_at,
		"updated_at":          certifictaeofemploment.Updated_at,
	}

	result := map[string]interface{}{
		"code":    200,
		"data":    certificate,
		"message": "Succesfully",
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resp))
}

func UploadResigns(w http.ResponseWriter, r *http.Request) {

	db, err := Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbhwi, _ = ConnHwi()
	defer dbhwi.Close()

	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadedFile, _, err := r.FormFile("file")

	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	reader := csv.NewReader(uploadedFile)

	records, _ := reader.ReadAll()

	var notification []string
	notification = append(notification, "")
	var code int = 200
	var Employee = Employee{}
	var Count_id int
	for _, record := range records {

		err = db.QueryRow("SELECT COUNT(id) as id , COALESCE(status_employee, '') as status_employee, COALESCE(name, ''), COALESCE(hire_date, ''), COALESCE(date_of_birth, ''), COALESCE(date_out, '0000-00-00'),COALESCE(job_id, 25),COALESCE(department_id, 116), COALESCE(address_jalan, ''), COALESCE(address_rt, ''), COALESCE(address_rw, ''), COALESCE(address_village, ''), COALESCE(address_district, ''), COALESCE(address_city, ''), COALESCE(address_province, '') FROM employees WHERE number_of_employees = ? ", record[0]).
			Scan(&Count_id, &Employee.Status_employee, &Employee.Name, &Employee.Hire_date, &Employee.Date_of_birth, &Employee.Date_out, &Employee.Job_id, &Employee.Department_id, &Employee.Address_jalan, &Employee.Address_rt, &Employee.Address_rw, &Employee.Address_village, &Employee.Address_district, &Employee.Address_city, &Employee.Address_province)

		if err != nil {
			fmt.Println(err.Error())
		}

		// fmt.Println("Ini Count_id ", Count_id, Employee.Status_employee, Employee.Name, Employee.Hire_date, Employee.Date_of_birth, Employee.Date_out, Employee.Address_jalan, Employee.Address_rt, Employee.Address_rw, Employee.Address_village, Employee.Address_district, Employee.Address_city, Employee.Address_province)
		if Count_id > 0 {
			// 	fmt.Println("Status ini => ", Employee.Status_employee)

			switch Employee.Status_employee {
			case "active":
				_, err = db.Exec("UPDATE employees SET date_out = '?' , status_employee = '?', exit_statement = '?' WHERE number_of_employees = '?' ", record[2], "notactive", record[3], record[0])
				if err != nil {
					fmt.Println(err.Error())
				}
			case "notactive":
				queryupdate := fmt.Sprintf("UPDATE employees SET date_out = '%s' , status_employee = '%s', exit_statement = '%s' WHERE number_of_employees = '%s' ", "0000-00-00", "active", record[3], record[0])
				_, err = db.Exec(queryupdate)
				if err != nil {
					fmt.Println(err.Error())
				}
			default:
				fmt.Println("Tidak Melakukan Transaksi update data karyawan")
			}

			/*
				if Employee.Status_employee == "active" {
					queryupdate := fmt.Sprintf("UPDATE employees SET date_out = '%s' , status_employee = '%s', exit_statement = '%s' WHERE number_of_employees = '%s' ", record[2], "notactive", record[3], record[0])
					_, err = db.Exec("UPDATE employees SET date_out = '?' , status_employee = '?', exit_statement = '?' WHERE number_of_employees = '?' ", record[2], "notactive", record[3], record[0])
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Println("UPDATE SUKSES")
					fmt.Println(queryupdate)
				} else if Employee.Status_employee == "notactive" {
					fmt.Println("Status ini notactive => aslinya ", Employee.Status_employee)
					queryupdate := fmt.Sprintf("UPDATE employees SET date_out = '%s' , status_employee = '%s', exit_statement = '%s' WHERE number_of_employees = '%s' ", "0000-00-00", "active", record[3], record[0])
					_, err = db.Exec(queryupdate)
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Println(queryupdate)
				}
			*/
		}

		var Count_idresigns = 0
		err = dbhwi.QueryRow("SELECT COUNT(id) as id FROM resigns WHERE number_of_employees = ? ", record[0]).
			Scan(&Count_idresigns)
		if err != nil {
			fmt.Println(err.Error())
		}

		if record[0] == "number_of_employees" {

		} else if Count_idresigns < 1 && Count_id > 0 {
			var Job_level, Department string
			err = db.QueryRow("SELECT job_level FROM jobs WHERE id = ?", Employee.Job_id).Scan(&Job_level)
			if err != nil {
				fmt.Println(err.Error())
			}
			err = db.QueryRow("SELECT department FROM departments WHERE id = ?", Employee.Department_id).Scan(&Department)
			if err != nil {
				fmt.Println(err.Error())
			}
			_, err = dbhwi.Exec("INSERT INTO `resigns`(	`number_of_employees`,`name`, `position`, `department`, `hire_date`, `classification`, `date_out`, `date_resignsubmissions`, `periode_of_service`, `type`, `age`, `status_resign`, `printed`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?,	?,	?,	?,	?,	?,	?,	?,	?,	?, ? , ?)", record[0], Employee.Name, Job_level, Department, Employee.Hire_date, CekDateSubmission(record[0]), record[2], nil, Periode_of_serve(Employee.Hire_date, record[2]), "false", Age(Employee.Date_of_birth), "resign", 0, DMYhms(), DMYhms())
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			each := fmt.Sprintf("NIK %s Tidak dapat resign karena sudah resign </br>", record[0])
			code = 404
			notification = append(notification, each)
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

func GetParklaringCertificates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var count_certificate_of_employments int

	err = db.QueryRow("SELECT COUNT(id) as count_certificate_of_employments FROM certificate_of_employments ").
		Scan(&count_certificate_of_employments)
	if count_certificate_of_employments == 0 {
		var datanull = []map[string]string{
			{"id": "NULL", "number_of_employees": "NULL", "name": "NULL", "hire_date": "NULL", "date_out": "NULL", "position": "NULL", "department": "NULL", "date_certificate_employee": "NULL", "no_certificate_employee": "NULL", "rom": "NULL", "created_at": "NULL", "updated_at": "NULL"},
		}

		result := map[string]interface{}{
			"code":  404,
			"meta":  "NULL",
			"data":  datanull,
			"links": "NULL",
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(resp))
		return
	}

	u, err := url.Parse(r.RequestURI)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	var sqlPaging string = "SELECT id, COALESCE(resign_id, 0) , COALESCE(number_of_employees, ''), COALESCE(date_certificate_employee, '0000-00-00'), COALESCE(no_certificate_employee, 0), COALESCE(rom, ''), COALESCE(created_at, ''), COALESCE(updated_at, '') FROM certificate_of_employments"
	var sqlCount string = "SELECT COUNT(*) FROM certificate_of_employments"
	var params string = ""

	number_of_employees, checkNumber_of_employees := q["number_of_employees"]
	if checkNumber_of_employees != false {
		justStringnumber_of_employees := strings.Join(number_of_employees, "")
		sqlPaging = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%'", sqlPaging, justStringnumber_of_employees)
		sqlCount = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%'", sqlCount, justStringnumber_of_employees)
		params = fmt.Sprintf("&%snumber_of_employees=%s", params, justStringnumber_of_employees)
	}

	var total int64
	db.QueryRow(sqlCount).Scan(&total)
	if total == 0 {
		var datanull = []map[string]string{
			{"id": "NULL", "number_of_employees": "NULL", "name": "NULL", "hire_date": "NULL", "date_out": "NULL", "position": "NULL", "department": "NULL", "date_certificate_employee": "NULL", "no_certificate_employee": "NULL", "rom": "NULL", "created_at": "NULL", "updated_at": "NULL"},
		}

		result := map[string]interface{}{
			"code":  404,
			"meta":  "NULL",
			"data":  datanull,
			"links": "NULL",
		}

		resp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write([]byte(resp))
		return
	}

	var totalminbyperpage, lastPage int64
	totalminbyperpage = total - ((total / 10) * 10)

	if totalminbyperpage == 0 {
		lastPage = (total / 10)
	} else {
		lastPage = ((total / 10) + 1)
	}

	page, _ := strconv.Atoi("1")
	cpage, checkPage := q["page"]
	if checkPage != false {
		spage, _ := strconv.Atoi(strings.Join(cpage, ""))
		page = spage
	}

	var first, last, next, prev string
	first, last, next, prev = "", "", "", ""

	first = "1"
	last = strconv.Itoa(int(lastPage))

	next = strconv.Itoa(int(page + 1))
	if int(page+1) >= int(lastPage) {
		next = strconv.Itoa(int(lastPage))
	}

	prev = strconv.Itoa(int(page - 1))
	if int(page) == 1 {
		prev = strconv.Itoa(int(page))
	}

	perPage, _ := strconv.Atoi("10")
	sqlPaging = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlPaging, perPage, (page-1)*perPage)

	rows, err := db.Query(sqlPaging)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var letter []map[string]interface{}

	for rows.Next() {
		var each = Letter{}
		var err = rows.Scan(&each.Id, &each.Resign_id, &each.Number_of_employees, &each.Date, &each.No, &each.Rom, &each.Created_at, &each.Updated_at)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var Resign Resign
		err = db.QueryRow("SELECT name, hire_date, date_out, position, department FROM resigns WHERE id = ? ", each.Resign_id).
			Scan(&Resign.Name, &Resign.Hire_date, &Resign.Date_out, &Resign.Position, &Resign.Department)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var data = map[string]interface{}{"id": each.Id, "number_of_employees": each.Number_of_employees, "name": Resign.Name, "hire_date": Resign.Hire_date, "date_out": Resign.Date_out, "position": Resign.Position, "department": Resign.Department, "date_certificate_employee": each.Date, "no": each.No, "rom": each.Rom, "created_at": each.Created_at, "update_at": each.Updated_at}

		letter = append(letter, data)
	}

	links := map[string]interface{}{
		"first": fmt.Sprintf("page=%s%s", first, params),
		"last":  fmt.Sprintf("page=%s%s", last, params),
		"next":  fmt.Sprintf("page=%s%s", next, params),
		"prev":  fmt.Sprintf("page=%s%s", prev, params),
	}

	informationpages := map[string]interface{}{
		"currentPage": page,
		"from":        ((page - 1) * 10) + 1,
		"lastPage":    lastPage,
		"perPage":     10,
		"to":          ((page - 1) * 10) + len(letter),
		"total":       total,
	}

	pages := map[string]interface{}{
		"page": informationpages,
	}

	result := map[string]interface{}{
		"code":  200,
		"meta":  pages,
		"data":  letter,
		"links": links,
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte(resp))

}

func GetEditParklaringCertificate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	Number_of_employess, _ := strconv.Atoi(vars["number_of_employees"])

	var dbhwi, err = ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	var letter Letter

	err = dbhwi.QueryRow("SELECT id, COALESCE(resign_id, 0) , COALESCE(number_of_employees, ''), COALESCE(date_certificate_employee, '0000-00-00'), COALESCE(no_certificate_employee, 0), COALESCE(rom, ''), COALESCE(created_at, ''), COALESCE(updated_at, '') FROM certificate_of_employments WHERE number_of_employees = ? ", Number_of_employess).
		Scan(&letter.Id, &letter.Resign_id, &letter.Number_of_employees, &letter.Date, &letter.No, &letter.Rom, &letter.Created_at, &letter.Updated_at)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var resign Resign

	err = dbhwi.QueryRow("SELECT COALESCE(name, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), COALESCE(position, ''), COALESCE(department, '') FROM resigns WHERE id = ? ", letter.Resign_id).
		Scan(&resign.Name, &resign.Hire_date, &resign.Date_out, &resign.Position, &resign.Department)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data := map[string]interface{}{
		"id":                        letter.Id,
		"name":                      resign.Name,
		"number_of_employees":       letter.Number_of_employees,
		"hire_date":                 resign.Hire_date,
		"date_out":                  resign.Date_out,
		"position":                  resign.Position,
		"department":                resign.Department,
		"date_certificate_employee": letter.Date,
		"no_certificate_employee":   letter.No,
		"rom":                       letter.Rom,
		"created_at":                letter.Created_at,
		"updated_at":                letter.Updated_at,
	}

	result := map[string]interface{}{
		"code":    200,
		"data":    data,
		"message": "Successfully",
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))

}

func GetUpdateParklaringCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data := Letter{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dbhwi, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	_, err = dbhwi.Exec("UPDATE `certificate_of_employments` SET `date_certificate_employee`= ? ,`no_certificate_employee`= ? ,`rom`=  ? ,`created_at`= ? ,`updated_at`= ?  WHERE number_of_employees = ? ", data.Date, data.No, data.Rom, data.Created_at, data.Updated_at, data.Number_of_employees)
	if err != nil {
		fmt.Println(err.Error())
		result := map[string]interface{}{
			"code":    400,
			"message": "Update Loss",
		}
		resp, _ := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		return
	}

	result := map[string]interface{}{
		"code":    200,
		"data":    data,
		"message": "Update Success",
	}

	resp, _ := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))
}

func ExportSubmission(w http.ResponseWriter, r *http.Request) {
	dbhwi, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	rows, err := dbhwi.Query("select number_of_employees from resignation_submissions where id > ?", 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	xlsx := excelize.NewFile()
	sheet1Name := "Sheet1"

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	err = xlsx.AutoFilter(sheet1Name, "A1", "Q1", "")
	if err != nil {
		fmt.Println(err.Error())
	}
	xlsx.SetCellValue(sheet1Name, "A1", "NIK")
	xlsx.SetCellValue(sheet1Name, "B1", "NAME")
	xlsx.SetCellValue(sheet1Name, "C1", "POSISI")
	xlsx.SetCellValue(sheet1Name, "D1", "DEPARTMENT")
	xlsx.SetCellValue(sheet1Name, "E1", "GEDUNG")
	xlsx.SetCellValue(sheet1Name, "F1", "HIRE DATE =DATE(LEFT(F2,4), MID(F2,6,2), RIGHT(F2,2))")
	xlsx.SetCellValue(sheet1Name, "G1", "DATE OUT =DATE(LEFT(G2,4), MID(G2,6,2), RIGHT(G2,2))")
	xlsx.SetCellValue(sheet1Name, "H1", "TYPE")
	xlsx.SetCellValue(sheet1Name, "I1", "ALASAM")
	xlsx.SetCellValue(sheet1Name, "J1", "ALASAN TAMBAHAN")
	xlsx.SetCellValue(sheet1Name, "K1", "UMUR")
	xlsx.SetCellValue(sheet1Name, "L1", "SARAN")
	xlsx.SetCellValue(sheet1Name, "M1", "STATUS RESIGN")
	xlsx.SetCellValue(sheet1Name, "N1", "USING MEDIA")
	xlsx.SetCellValue(sheet1Name, "O1", "CLASSIFIKASI")
	xlsx.SetCellValue(sheet1Name, "P1", "CREATEAD AT")
	xlsx.SetCellValue(sheet1Name, "Q1", "UPDATED AT")
	xlsx.SetCellValue(sheet1Name, "R1", "TGL PERMOHONAN =DATE(LEFT(R2,4), MID(R2,6,2), RIGHT(R2,2))")

	var wg sync.WaitGroup

	no := 1

	for rows.Next() {
		var NIK string
		var err = rows.Scan(&NIK)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		no += 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, message string, no int) {
			defer wg.Done()
			var Submission ResignSubmission

			err = dbhwi.QueryRow("SELECT COALESCE(number_of_employees, ''),	COALESCE(name, ''),	COALESCE(position, ''),	COALESCE(department, ''),	COALESCE(building, ''),	COALESCE(hire_date, '0000-00-00'),	COALESCE(date_out, '0000-00-00'),	COALESCE(date_resignation_submissions, '0000-00-00'),	COALESCE(type, ''),	COALESCE(reason, ''),	COALESCE(detail_reason, ''),	COALESCE(periode_of_service, ''),	COALESCE(age, 0),	COALESCE(suggestion, ''),	COALESCE(status_resignsubmisssion, ''),	COALESCE(using_media, ''),	COALESCE(classification, ''),	COALESCE(created_at, '0000-00-00 00:00:00'),	COALESCE(updated_at, '0000-00-00 00:00:00')	from resignation_submissions where number_of_employees = ?", message).
				Scan(&Submission.Number_of_employees, &Submission.Name, &Submission.Position, &Submission.Department, &Submission.Building, &Submission.Hire_date, &Submission.Date_out, &Submission.Date_resignation_submissions, &Submission.Type, &Submission.Reason, &Submission.Detail_reason, &Submission.Periode_of_service, &Submission.Age, &Submission.Suggestion, &Submission.Status_resignsubmisssion, &Submission.Using_media, &Submission.Classification, &Submission.Created_at, &Submission.Updated_at)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), Submission.Number_of_employees)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", no), Submission.Name)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", no), Submission.Position)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", no), Submission.Department)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", no), Submission.Building)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", no), Submission.Hire_date)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", no), Submission.Date_out)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", no), Submission.Type)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("I%d", no), Submission.Reason)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", no), Submission.Detail_reason)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", no), Submission.Age)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("L%d", no), Submission.Suggestion)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("M%d", no), Submission.Status_resignsubmisssion)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("N%d", no), Submission.Using_media)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("O%d", no), Submission.Classification)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("P%d", no), Submission.Created_at)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("Q%d", no), Submission.Updated_at)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("R%d", no), Submission.Date_resignation_submissions)
		}(&wg, NIK, no)
	}

	wg.Wait()

	var b bytes.Buffer
	writr := bufio.NewWriter(&b)
	xlsx.Write(writr)
	writr.Flush()
	fileContents := b.Bytes()
	fileSize := strconv.Itoa(len(fileContents))

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-disposition", "attachment;filename=Data_Pengajuan_Design.xlsx")
	w.Header().Set("Content-Length", fileSize)

	t := bytes.NewReader(b.Bytes())
	io.Copy(w, t)

	fmt.Fprintln(w, "Download Sukses")
}

func ExportResign(w http.ResponseWriter, r *http.Request) {
	dbhwi, err := ConnHwi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhwi.Close()

	rows, err := dbhwi.Query("select number_of_employees from resigns")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	xlsx := excelize.NewFile()
	sheet1Name := "Sheet1"

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	err = xlsx.AutoFilter(sheet1Name, "A1", "Q1", "")
	if err != nil {
		fmt.Println(err.Error())
	}

	xlsx.SetCellValue(sheet1Name, "A1", "NIK")
	xlsx.SetCellValue(sheet1Name, "B1", "NAME")
	xlsx.SetCellValue(sheet1Name, "C1", "POSISI")
	xlsx.SetCellValue(sheet1Name, "D1", "DEPARTMENT")
	xlsx.SetCellValue(sheet1Name, "E1", "GEDUNG")
	xlsx.SetCellValue(sheet1Name, "F1", "HIRE DATE =DATE(LEFT(F2,4), MID(F2,6,2), RIGHT(F2,2))")
	xlsx.SetCellValue(sheet1Name, "G1", "DATE OUT =DATE(LEFT(G2,4), MID(G2,6,2), RIGHT(G2,2))")
	xlsx.SetCellValue(sheet1Name, "H1", "TYPE")
	xlsx.SetCellValue(sheet1Name, "I1", "ALASAM")
	xlsx.SetCellValue(sheet1Name, "J1", "ALASAN TAMBAHAN")
	xlsx.SetCellValue(sheet1Name, "K1", "UMUR")
	xlsx.SetCellValue(sheet1Name, "L1", "SARAN")
	xlsx.SetCellValue(sheet1Name, "M1", "STATUS RESIGN")
	xlsx.SetCellValue(sheet1Name, "N1", "USING MEDIA")
	xlsx.SetCellValue(sheet1Name, "O1", "CLASSIFIKASI")
	xlsx.SetCellValue(sheet1Name, "P1", "CREATEAD AT")
	xlsx.SetCellValue(sheet1Name, "Q1", "UPDATED AT")
	xlsx.SetCellValue(sheet1Name, "R1", "TGL PERMOHONAN =DATE(LEFT(R2,4), MID(R2,6,2), RIGHT(R2,2))")

	var wg sync.WaitGroup

	no := 1

	for rows.Next() {
		var NIK string
		var err = rows.Scan(&NIK)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		no += 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, message string, no int) {
			defer wg.Done()
			var Resign Resign

			err = dbhwi.QueryRow("SELECT COALESCE(number_of_employees, ''),	COALESCE(name, ''),	COALESCE(position, ''),	COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'),	COALESCE(date_out, '0000-00-00'),	COALESCE(date_resignsubmissions, '0000-00-00'),	COALESCE(type, ''),	COALESCE(age, 0),	COALESCE(status_resign, ''),	COALESCE(classification, ''),	COALESCE(created_at, '0000-00-00 00:00:00'),	COALESCE(updated_at, '0000-00-00 00:00:00')	from resigns where number_of_employees = ?", message).
				Scan(&Resign.Number_of_employees, &Resign.Name, &Resign.Position, &Resign.Department, &Resign.Hire_date, &Resign.Date_out, &Resign.Date_resignsubmissions, &Resign.Type, &Resign.Age, &Resign.Status_resign, &Resign.Classification, &Resign.Created_at, &Resign.Updated_at)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), Resign.Number_of_employees)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", no), Resign.Name)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", no), Resign.Position)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", no), Resign.Department)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", no), Resign.Hire_date)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", no), Resign.Date_out)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", no), Resign.Type)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", no), Resign.Age)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("I%d", no), Resign.Status_resign)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", no), Resign.Classification)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", no), Resign.Created_at)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("L%d", no), Resign.Updated_at)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("M%d", no), Resign.Date_resignsubmissions)
		}(&wg, NIK, no)
	}

	wg.Wait()

	var b bytes.Buffer
	writr := bufio.NewWriter(&b)
	xlsx.Write(writr)
	writr.Flush()
	fileContents := b.Bytes()
	fileSize := strconv.Itoa(len(fileContents))

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-disposition", "attachment;filename=Data_resign.xlsx")
	w.Header().Set("Content-Length", fileSize)

	t := bytes.NewReader(b.Bytes())
	io.Copy(w, t)

	fmt.Fprintln(w, "Download Sukses")
}

func Age(DateString string) int {

	var s string
	s = DateString
	yearDate, _ := strconv.Atoi(string(s[0:4]))
	monthDate, _ := strconv.Atoi(string(s[5:7]))
	dayDate, _ := strconv.Atoi(string(s[8:10]))

	month := time.Month(monthDate)

	date := time.Date(yearDate, month, dayDate, 0, 0, 0, 0, time.UTC)
	dateAge := age.Calculate(date)

	return dateAge
}

func Periode_of_serve(DateString string, DateString2 string) int {

	var s string
	s = DateString
	yearDate, _ := strconv.Atoi(string(s[0:4]))
	monthDate, _ := strconv.Atoi(string(s[5:7]))
	dayDate, _ := strconv.Atoi(string(s[8:10]))

	var s2 string
	// s2 = strings.Join(DateString2, "")
	s2 = DateString2
	yearDate2, _ := strconv.Atoi(string(s2[0:4]))
	monthDate2, _ := strconv.Atoi(string(s2[5:7]))
	dayDate2, _ := strconv.Atoi(string(s2[8:10]))

	t1 := Date(yearDate, monthDate, dayDate)
	t2 := Date(yearDate2, monthDate2, dayDate2)
	days := t2.Sub(t1).Hours() / 24
	return int(days)
}

func CekDateSubmission(Number_of_employees string) string {
	var dbhwi, err = ConnHwi()
	defer dbhwi.Close()

	var Count_id int
	var Submission ResignSubmission

	err = dbhwi.QueryRow("SELECT COUNT(id) as id, COALESCE(date_resignation_submissions, '0000-00-00') FROM resignation_submissions WHERE number_of_employees = ? AND status_resignsubmisssion = 'wait'  ", Number_of_employees).
		Scan(&Count_id, &Submission.Date_resignation_submissions)
	if err != nil {
		fmt.Println(err.Error())
	}

	var output string
	switch Count_id {
	case 1:
		var s string
		s = Submission.Date_resignation_submissions
		yearDate, _ := strconv.Atoi(string(s[0:4]))
		monthDate, _ := strconv.Atoi(string(s[5:7]))
		dayDate, _ := strconv.Atoi(string(s[8:10]))

		month := time.Month(monthDate)

		theTime := time.Date(yearDate, month, dayDate, 0, 0, 0, 0, time.Local)

		after := theTime.AddDate(0, 0, 14)

		var stringafter string
		stringafter = after.Format("2006-01-02")

		var currentTime string
		currentTime = time.Now().Format("2006-01-02")

		status_type := Periode_of_serve(currentTime, stringafter)

		if status_type <= 0 {
			output = "Sudah mengajukan resign dan resign sesuai procedure"
		} else {
			output = "Mengajukan resign tetapi resign sebelum waktunya"
		}
	default:
		output = "Resign dahulu sebelum mengajukan resign"
	}

	return output
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

	r.HandleFunc("/resignsubmissions", GetResignSubmission).Methods("GET")
	r.HandleFunc("/resignsubmissions/{search}", GetResignSubmissionSearch).Methods("GET")
	r.HandleFunc("/resignsubmissions/{number_of_employees}/{status_resign}", GetResignSubmissionStatus).Methods("GET")
	r.HandleFunc("/resignsubmission_upload", UploadSubmission).Methods("POST")
	r.HandleFunc("/resignsubmission_edit/{number_of_employees}", GetEditSubmission).Methods("GET")
	r.HandleFunc("/resignsubmission_update", GetUpdateSubmission).Methods("POST")

	r.HandleFunc("/resign", Post).Methods("POST")
	r.HandleFunc("/resignacc", PostAcc).Methods("POST")
	r.HandleFunc("/ExportSubmission", ExportSubmission).Methods("GET")

	//Resigns
	r.HandleFunc("/resigns", GetResigns).Methods("GET")
	r.HandleFunc("/resigns/upload", UploadResigns).Methods("POST")
	r.HandleFunc("/resigns_edit/{number_of_employees}", GetEditResign).Methods("GET")
	r.HandleFunc("/resigns_update", GetUpdateResign).Methods("POST")
	r.HandleFunc("/resigns/makecertificate", PostCertifcate).Methods("POST")
	r.HandleFunc("/resigns/makeexperience", PostExperience).Methods("POST")
	r.HandleFunc("/ExportResign", ExportResign).Methods("GET")

	//Parklaring
	r.HandleFunc("/parklarings_certificate", GetParklaringCertificates).Methods("GET")
	r.HandleFunc("/parklarings_certificateedit/{number_of_employees}", GetEditParklaringCertificate).Methods("GET")
	r.HandleFunc("/parklarings_certificateupdate", GetUpdateParklaringCertificate).Methods("POST")

	// r.HandleFunc("/user/{id}", Update).Methods("PUT")
	// r.HandleFunc("/user/{id}", Delete).Methods("DELETE")

	fmt.Println("LIsten on Port 10.10.42.6:8880")
	http.ListenAndServe(":8880", r)

}
