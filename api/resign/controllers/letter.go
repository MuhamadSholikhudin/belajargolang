package controllers

import (
	"belajargolang/api/resign/helper"
	"belajargolang/api/resign/models"
	"bufio"
	"bytes"
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
	"github.com/gorilla/mux"
)

func PostCertifcate(w http.ResponseWriter, r *http.Request) {

	//untuk membuat json pertama kita harus set Header
	w.Header().Set("Content-Type", "application/json")

	data := models.Letter{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var Resign_id int
	var ResignSel = models.Resign{}
	err = db.QueryRow("SELECT id as resign_id, name, COALESCE(position, ''), COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') FROM resigns WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&Resign_id, &ResignSel.Name, &ResignSel.Position, &ResignSel.Department, &ResignSel.Hire_date, &ResignSel.Date_out)
	if err != nil {
		fmt.Print(err.Error())
	}

	var CountCertifcateNumberOf_employees int

	err = db.QueryRow("SELECT COUNT(*) FROM certificate_of_employments WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&CountCertifcateNumberOf_employees)
	if err != nil {
		fmt.Print(err.Error())
	}

	if CountCertifcateNumberOf_employees == 0 {
		var CountCertificateByDate, CountNoCertificateEmployee int
		var yearstring string
		yearstring = strconv.Itoa(time.Now().Year())
		err = db.QueryRow("SELECT COUNT(id) as CountCertificateByDate, COALESCE(no_certificate_employee, 0) as no_certificate_employee FROM certificate_of_employments WHERE YEAR(date_certificate_employee) = ? AND MONTH(date_certificate_employee) = ? ORDER BY date_certificate_employee DESC", yearstring, helper.StringMonth()).
			Scan(&CountCertificateByDate, &CountNoCertificateEmployee)
		if err != nil {
			fmt.Print(err.Error())
		}
		// create certificate
		_, err := db.Exec("INSERT INTO certificate_of_employments (resign_id, number_of_employees, date_certificate_employee, no_certificate_employee, rom, created_at, updated_at) VALUES (?,?,?,?,?,?,?)", Resign_id, data.Number_of_employees, helper.DMY(), (CountNoCertificateEmployee + 1), helper.Rom(helper.StringMonth()), helper.DMYhms(), helper.DMYhms())
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	var certifictaeofemploment = models.Letter{}
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

	data := models.Letter{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var Resign_id int
	var ResignSel = models.Resign{}
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
		err = db.QueryRow("SELECT COUNT(id) as CountCertificateByDate, COALESCE(no_letter_experience, 0) as no_letter_experience FROM work_experience_letters WHERE YEAR(date_certificate_employee) = ? AND MONTH(date_certificate_employee) = ? ORDER BY date_certificate_employee DESC", yearstring, helper.StringMonth()).
			Scan(&CountCertificateByDate, &CountNoCertificateEmployee)
		if err != nil {
			fmt.Print(err.Error())
		}
		// create certificate
		_, err := db.Exec("INSERT INTO work_experience_letters (resign_id, number_of_employees, date_letter_exprerience, no_letter_experience, rom, created_at, updated_at) VALUES (?,?,?,?,?,?,?)", Resign_id, data.Number_of_employees, helper.DMY(), (CountNoCertificateEmployee + 1), helper.Rom(helper.StringMonth()), helper.DMYhms(), helper.DMYhms())
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	var certifictaeofemploment = models.Letter{}

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

func GetParklaringCertificates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnResign()
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
			{"id": "NULL", "number_of_employees": "NULL", "name": "NULL", "hire_date": "NULL", "date_out": "NULL", "position": "NULL", "department": "NULL", "date_letter": "NULL", "no_letter": "NULL", "rom": "NULL", "created_at": "NULL", "updated_at": "NULL"},
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
		var each = models.Letter{}
		var err = rows.
			Scan(&each.Id, &each.Resign_id, &each.Number_of_employees, &each.Date, &each.No, &each.Rom, &each.Created_at, &each.Updated_at)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var Resign models.Resign
		err = db.QueryRow("SELECT COALESCE(name, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), COALESCE(position, ''), COALESCE(department, '') FROM resigns WHERE id = ? ", each.Resign_id).
			Scan(&Resign.Name, &Resign.Hire_date, &Resign.Date_out, &Resign.Position, &Resign.Department)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var data = map[string]interface{}{"id": each.Id, "number_of_employees": each.Number_of_employees, "name": Resign.Name, "hire_date": Resign.Hire_date, "date_out": Resign.Date_out, "position": Resign.Position, "department": Resign.Department, "date_letter": each.Date, "no_letter": each.No, "rom": each.Rom, "created_at": each.Created_at, "update_at": each.Updated_at}

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

	var dbresign, err = models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	var letter models.Letter

	err = dbresign.QueryRow("SELECT id, COALESCE(resign_id, 0) , COALESCE(number_of_employees, ''), COALESCE(date_certificate_employee, '0000-00-00'), COALESCE(no_certificate_employee, 0), COALESCE(rom, ''), COALESCE(created_at, ''), COALESCE(updated_at, '') FROM certificate_of_employments WHERE number_of_employees = ? ", Number_of_employess).
		Scan(&letter.Id, &letter.Resign_id, &letter.Number_of_employees, &letter.Date, &letter.No, &letter.Rom, &letter.Created_at, &letter.Updated_at)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var resign models.Resign

	err = dbresign.QueryRow("SELECT COALESCE(name, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), COALESCE(position, ''), COALESCE(department, '') FROM resigns WHERE id = ? ", letter.Resign_id).
		Scan(&resign.Name, &resign.Hire_date, &resign.Date_out, &resign.Position, &resign.Department)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data := map[string]interface{}{
		"id":                  letter.Id,
		"name":                resign.Name,
		"number_of_employees": letter.Number_of_employees,
		"hire_date":           resign.Hire_date,
		"date_out":            resign.Date_out,
		"position":            resign.Position,
		"department":          resign.Department,
		"date_letter":         letter.Date,
		"no_letter":           letter.No,
		"rom":                 letter.Rom,
		"created_at":          letter.Created_at,
		"updated_at":          letter.Updated_at,
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

	data := models.Letter{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	_, err = dbresign.Exec("UPDATE `certificate_of_employments` SET `date_certificate_employee`= ? ,`no_certificate_employee`= ? ,`rom`=  ? ,`created_at`= ? ,`updated_at`= ?  WHERE number_of_employees = ? ", data.Date, data.No, data.Rom, data.Created_at, data.Updated_at, data.Number_of_employees)
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

func GetParklaringExperiences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var count_work_experience_letters int

	err = db.QueryRow("SELECT COUNT(id) as count_work_experience_letters FROM work_experience_letters ").
		Scan(&count_work_experience_letters)
	if count_work_experience_letters == 0 {
		var datanull = []map[string]string{
			{"id": "NULL", "number_of_employees": "NULL", "name": "NULL", "hire_date": "NULL", "date_out": "NULL", "position": "NULL", "department": "NULL", "date_letter": "NULL", "no_letter": "NULL", "rom": "NULL", "created_at": "NULL", "updated_at": "NULL"},
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

	var sqlPaging string = "SELECT id, COALESCE(resign_id, 0) , COALESCE(number_of_employees, ''), COALESCE(date_letter_exprerience, '0000-00-00'), COALESCE(no_letter_experience, 0), COALESCE(rom, ''), COALESCE(created_at, ''), COALESCE(updated_at, '') FROM work_experience_letters"
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
			{"id": "NULL", "number_of_employees": "NULL", "name": "NULL", "hire_date": "NULL", "date_out": "NULL", "position": "NULL", "department": "NULL", "date_letter": "NULL", "no_letter": "NULL", "rom": "NULL", "created_at": "NULL", "updated_at": "NULL"},
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
		var each = models.Letter{}
		var err = rows.
			Scan(&each.Id, &each.Resign_id, &each.Number_of_employees, &each.Date, &each.No, &each.Rom, &each.Created_at, &each.Updated_at)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var Resign models.Resign
		err = db.QueryRow("SELECT COALESCE(name, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), COALESCE(position, ''), COALESCE(department, '') FROM resigns WHERE id = ? ", each.Resign_id).
			Scan(&Resign.Name, &Resign.Hire_date, &Resign.Date_out, &Resign.Position, &Resign.Department)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var data = map[string]interface{}{"id": each.Id, "number_of_employees": each.Number_of_employees, "name": Resign.Name, "hire_date": Resign.Hire_date, "date_out": Resign.Date_out, "position": Resign.Position, "department": Resign.Department, "date_letter": each.Date, "no_letter": each.No, "rom": each.Rom, "created_at": each.Created_at, "update_at": each.Updated_at}
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

func GetEditParklaringExperience(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	Number_of_employess, _ := strconv.Atoi(vars["number_of_employees"])

	var dbresign, err = models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	var letter models.Letter

	err = dbresign.QueryRow("SELECT id, COALESCE(resign_id, 0) , COALESCE(number_of_employees, ''), COALESCE(date_letter_exprerience, '0000-00-00'), COALESCE(no_letter_experience, 0), COALESCE(rom, ''), COALESCE(created_at, ''), COALESCE(updated_at, '') FROM work_experience_letters WHERE number_of_employees = ? ", Number_of_employess).
		Scan(&letter.Id, &letter.Resign_id, &letter.Number_of_employees, &letter.Date, &letter.No, &letter.Rom, &letter.Created_at, &letter.Updated_at)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var resign models.Resign

	err = dbresign.QueryRow("SELECT COALESCE(name, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), COALESCE(position, ''), COALESCE(department, '') FROM resigns WHERE id = ? ", letter.Resign_id).
		Scan(&resign.Name, &resign.Hire_date, &resign.Date_out, &resign.Position, &resign.Department)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data := map[string]interface{}{
		"id":                  letter.Id,
		"name":                resign.Name,
		"number_of_employees": letter.Number_of_employees,
		"hire_date":           resign.Hire_date,
		"date_out":            resign.Date_out,
		"position":            resign.Position,
		"department":          resign.Department,
		"date_letter":         letter.Date,
		"no_letter":           letter.No,
		"rom":                 letter.Rom,
		"created_at":          letter.Created_at,
		"updated_at":          letter.Updated_at,
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

func GetUpdateParklaringExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data := models.Letter{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	_, err = dbresign.Exec("UPDATE `work_experience_letters` SET `date_letter_exprerience`= ? ,`no_letter_experience`= ? ,`rom`=  ? ,`created_at`= ? ,`updated_at`= ?  WHERE number_of_employees = ? ", data.Date, data.No, data.Rom, data.Created_at, data.Updated_at, data.Number_of_employees)
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

func ExportLetter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	dataletter, _ := vars["dataletter"]

	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	var table, letter_Date, letter_No string

	if dataletter == "certificate_of_employments" {
		table = "certificate_of_employments"
		letter_Date = "date_certificate_employee"
		letter_No = "no_certificate_employee"
	} else {
		table = "work_experience_letters"
		letter_Date = "date_letter_exprerience"
		letter_No = "no_letter_experience"
	}

	query := fmt.Sprintf("select resign_id, COALESCE(%s, '0000-00-00'), COALESCE(%s, 0), COALESCE(rom, ''), COALESCE(created_at, '0000-00-00 00:00:00'), COALESCE(updated_at, '0000-00-00 00:00:00') from %s", letter_Date, letter_No, table)

	// rows, err := dbresign.Query("select resign_id, COALESCE(?, '0000-00-00'), COALESCE(?, 0), COALESCE(rom, ''), COALESCE(created_at, '0000-00-00 00:00:00'), COALESCE(updated_at, '0000-00-00 00:00:00') from ?", letter_Date, letter_No, dataletter)
	rows, err := dbresign.Query(query)
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
	xlsx.SetCellValue(sheet1Name, "E1", "HIRE DATE =DATE(LEFT(E2,4), MID(E2,6,2), RIGHT(E2,2))")
	xlsx.SetCellValue(sheet1Name, "F1", "DATE OUT =DATE(LEFT(F2,4), MID(F2,6,2), RIGHT(F2,2))")
	xlsx.SetCellValue(sheet1Name, "G1", "TYPE")
	xlsx.SetCellValue(sheet1Name, "H1", "NOMOR SURAT")
	xlsx.SetCellValue(sheet1Name, "I1", "STATUS SURAT")
	xlsx.SetCellValue(sheet1Name, "J1", "CLASSIFIKASI")
	xlsx.SetCellValue(sheet1Name, "K1", "UMUR")
	xlsx.SetCellValue(sheet1Name, "L1", "CREATEAD AT")
	xlsx.SetCellValue(sheet1Name, "M1", "UPDATED AT")

	var wg sync.WaitGroup

	no := 1
	for rows.Next() {
		var letter models.Letter
		var err = rows.Scan(&letter.Resign_id, &letter.Date, &letter.No, &letter.Rom, &letter.Created_at, &letter.Updated_at)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		no = (no + 1)
		wg.Add(1)
		go func(wg *sync.WaitGroup, no int, Resign_id int, Date string, No string, Rom string, Created_at string, Updated_at string) {
			defer wg.Done()
			var Resign models.Resign

			err = dbresign.QueryRow("SELECT COALESCE(number_of_employees, ''),	COALESCE(name, ''),	COALESCE(position, ''),	COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'),	COALESCE(date_out, '0000-00-00'),	COALESCE(date_resignsubmissions, '0000-00-00'),	COALESCE(type, ''),	COALESCE(age, 0),	COALESCE(status_resign, ''),	COALESCE(classification, ''),	COALESCE(created_at, '0000-00-00 00:00:00'),	COALESCE(updated_at, '0000-00-00 00:00:00')	from resigns where id = ?", Resign_id).
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
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", no), helper.NomorLetter(No, "/SKK_HR/HWI/", Rom, Date))
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("I%d", no), Resign.Status_resign)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", no), Resign.Classification)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", no), Resign.Age)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("L%d", no), Resign.Created_at)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("M%d", no), Resign.Updated_at)

		}(&wg, no, letter.Resign_id, letter.Date, letter.No, letter.Rom, letter.Created_at, letter.Updated_at)
	}

	wg.Wait()

	var b bytes.Buffer
	writr := bufio.NewWriter(&b)
	xlsx.Write(writr)
	writr.Flush()
	fileContents := b.Bytes()
	fileSize := strconv.Itoa(len(fileContents))

	attachment := fmt.Sprintf("attachment;filename=%s.xlsx", table)

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-disposition", attachment)
	w.Header().Set("Content-Length", fileSize)

	t := bytes.NewReader(b.Bytes())
	io.Copy(w, t)

	fmt.Fprintln(w, "Download Sukses")
}
