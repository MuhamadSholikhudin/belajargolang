package controllers

import (
	"belajargolang/api/resign/config"
	"belajargolang/api/resign/helper"
	"belajargolang/api/resign/models"
	"belajargolang/api/resign/repository"
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
	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	var Resign_id int
	var ResignSel = models.Resign{}
	err = dbresign.QueryRow("SELECT id as resign_id, name, COALESCE(position, ''), COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') FROM resigns WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&Resign_id, &ResignSel.Name, &ResignSel.Position, &ResignSel.Department, &ResignSel.Hire_date, &ResignSel.Date_out)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Update Resign
	var data2 = map[string]interface{}{
		"status_resign": "acc",
	}
	where := fmt.Sprintf("id = %d ", Resign_id)
	repository.UpdateResign("resigns", data2, where)

	var CountCertifcateNumberOf_employees int

	err = dbresign.QueryRow("SELECT COUNT(*) FROM certificate_of_employments WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&CountCertifcateNumberOf_employees)
	if err != nil {
		fmt.Print(err.Error())
	}

	if CountCertifcateNumberOf_employees == 0 {
		var CountNoCertificateEmployee int
		var yearstring string
		yearstring = strconv.Itoa(time.Now().Year())
		err = dbresign.QueryRow("SELECT COALESCE(no_certificate_employee, 0) as no_certificate_employee FROM certificate_of_employments WHERE YEAR(date_certificate_employee) = ? AND MONTH(date_certificate_employee) = ? ORDER BY no_certificate_employee DESC", yearstring, helper.StringMonth()).
			Scan(&CountNoCertificateEmployee)
		if err != nil {
			fmt.Print(err.Error())
		}
		data := map[string]interface{}{
			"resign_id":                 Resign_id,
			"number_of_employees":       data.Number_of_employees,
			"date_certificate_employee": helper.DMY(),
			"no_certificate_employee":   (CountNoCertificateEmployee + 1),
			"rom":                       helper.Rom(helper.StringMonth()),
			"created_at":                helper.DMYhms(),
			"updated_at":                helper.DMYhms(),
		}
		repository.InsertResign("certificate_of_employments", data)
	}

	var certifictaeofemploment = models.Letter{}
	err = dbresign.QueryRow("SELECT id, resign_id, number_of_employees, date_certificate_employee, no_certificate_employee, rom, created_at, updated_at FROM certificate_of_employments WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&certifictaeofemploment.Id, &certifictaeofemploment.Resign_id, &certifictaeofemploment.Number_of_employees, &certifictaeofemploment.Date, &certifictaeofemploment.No, &certifictaeofemploment.Rom, &certifictaeofemploment.Created_at, &certifictaeofemploment.Updated_at)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.QueryRow("SELECT COALESCE(date_out, '0000-00-00') FROM employees WHERE number_of_employees = ?", data.Number_of_employees).
		Scan(&ResignSel.Date_out)
	if err != nil {
		fmt.Println(err.Error())
	}

	certificate := map[string]interface{}{
		"id":                  certifictaeofemploment.Id,
		"resign_id":           certifictaeofemploment.Resign_id,
		"number_of_employees": certifictaeofemploment.Number_of_employees,
		"name":                ResignSel.Name,
		"position":            ResignSel.Position,
		"department":          helper.Deptout(ResignSel.Department),
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

	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	var Resign_id int
	var ResignSel = models.Resign{}
	err = dbresign.QueryRow("SELECT id as resign_id, name, COALESCE(position, ''), COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') FROM resigns WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&Resign_id, &ResignSel.Name, &ResignSel.Position, &ResignSel.Department, &ResignSel.Hire_date, &ResignSel.Date_out)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Update Resign
	var data2 = map[string]interface{}{
		"status_resign": "acc",
	}
	where := fmt.Sprintf("id = %d ", Resign_id)
	repository.UpdateResign("resigns", data2, where)

	var CountExperienceNumberOf_employees int
	err = dbresign.QueryRow("SELECT COUNT(*) FROM work_experience_letters WHERE number_of_employees = ?", data.Number_of_employees).
		Scan(&CountExperienceNumberOf_employees)
	if err != nil {
		fmt.Print(err.Error())
	}

	if CountExperienceNumberOf_employees == 0 {
		var CountNoExperienceEmployee int
		var yearstring string
		yearstring = strconv.Itoa(time.Now().Year())
		err = dbresign.QueryRow("SELECT COALESCE(no_letter_experience, 0) as no_letter_experience FROM work_experience_letters WHERE YEAR(date_letter_experience) = ? AND MONTH(date_letter_experience) = ? ORDER BY no_letter_experience DESC", yearstring, helper.StringMonth()).
			Scan(&CountNoExperienceEmployee)
		if err != nil {
			fmt.Print(err.Error())
		}
		data := map[string]interface{}{
			"resign_id":              Resign_id,
			"number_of_employees":    data.Number_of_employees,
			"date_letter_experience": helper.DMY(),
			"no_letter_experience":   (CountNoExperienceEmployee + 1),
			"rom":                    helper.Rom(helper.StringMonth()),
			"created_at":             helper.DMYhms(),
			"updated_at":             helper.DMYhms(),
		}
		repository.InsertResign("work_experience_letters", data)
	}

	var workexperienceletter = models.Letter{}
	err = dbresign.QueryRow("SELECT id, resign_id, number_of_employees, date_letter_experience, no_letter_experience, rom, created_at, updated_at FROM work_experience_letters WHERE number_of_employees = ? ", data.Number_of_employees).
		Scan(&workexperienceletter.Id, &workexperienceletter.Resign_id, &workexperienceletter.Number_of_employees, &workexperienceletter.Date, &workexperienceletter.No, &workexperienceletter.Rom, &workexperienceletter.Created_at, &workexperienceletter.Updated_at)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.QueryRow("SELECT COALESCE(date_out, '0000-00-00') FROM employees WHERE number_of_employees = ?", data.Number_of_employees).
		Scan(&ResignSel.Date_out)
	if err != nil {
		fmt.Println(err.Error())
	}

	certificate := map[string]interface{}{
		"id":                  workexperienceletter.Id,
		"resign_id":           workexperienceletter.Resign_id,
		"number_of_employees": workexperienceletter.Number_of_employees,
		"name":                ResignSel.Name,
		"position":            ResignSel.Position,
		"department":          helper.Deptout(ResignSel.Department),
		"hire_date":           ResignSel.Hire_date,
		"date_out":            ResignSel.Date_out,
		"date":                workexperienceletter.Date,
		"no":                  workexperienceletter.No,
		"rom":                 workexperienceletter.Rom,
		"created_at":          workexperienceletter.Created_at,
		"updated_at":          workexperienceletter.Updated_at,
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
	sqlPaging = fmt.Sprintf("%s ORDER BY id DESC", sqlPaging)
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

	var sqlPaging string = "SELECT id, COALESCE(resign_id, 0) , COALESCE(number_of_employees, ''), COALESCE(date_letter_experience, '0000-00-00'), COALESCE(no_letter_experience, 0), COALESCE(rom, ''), COALESCE(created_at, ''), COALESCE(updated_at, '') FROM work_experience_letters"
	var sqlCount string = "SELECT COUNT(*) FROM work_experience_letters"
	var params string = ""

	number_of_employees, checkNumber_of_employees := q["number_of_employees"]
	if checkNumber_of_employees != false {
		justStringnumber_of_employees := strings.Join(number_of_employees, "")
		sqlPaging = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%' ORDER BY id DESC", sqlPaging, justStringnumber_of_employees)
		sqlCount = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%'", sqlCount, justStringnumber_of_employees)
		params = fmt.Sprintf("&%snumber_of_employees=%s", params, justStringnumber_of_employees)
	}
	sqlPaging = fmt.Sprintf("%s ORDER BY id DESC", sqlPaging)

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

	var db, _ = models.ConnHrd()
	defer db.Close()

	var dbresign, err = models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	var letter models.Letter
	err = dbresign.QueryRow("SELECT id, COALESCE(resign_id, 0) , COALESCE(number_of_employees, ''), COALESCE(date_letter_experience, '0000-00-00'), COALESCE(no_letter_experience, 0), COALESCE(rom, ''), COALESCE(created_at, ''), COALESCE(updated_at, '') FROM work_experience_letters WHERE number_of_employees = ? ", Number_of_employess).
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

	var Date_out string
	err = db.QueryRow("SELECT COALESCE(date_out, '0000-00-00') FROM employees WHERE number_of_employees = ?", Number_of_employess).
		Scan(&Date_out)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data := map[string]interface{}{
		"id":                  letter.Id,
		"name":                resign.Name,
		"number_of_employees": letter.Number_of_employees,
		"hire_date":           resign.Hire_date,
		"date_out":            Date_out,
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

	_, err = dbresign.Exec("UPDATE `work_experience_letters` SET `date_letter_experience`= ? ,`no_letter_experience`= ? ,`rom`=  ? ,`created_at`= ? ,`updated_at`= ?  WHERE number_of_employees = ? ", data.Date, data.No, data.Rom, data.Created_at, data.Updated_at, data.Number_of_employees)
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

func DeleteLetter(w http.ResponseWriter, r *http.Request) {
	// set header dengan type json
	w.Header().Set("Content-type", "application/json")

	// decode body jadi
	decoder := json.NewDecoder(r.Body)
	payload := struct {
		Number_of_employees string `json:"number_of_employees"`
		List                string `json:"list"`
	}{}
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Pilih tabel yang akan di hapus
	var table string
	if payload.List == "certificate" {
		table = "certificate_of_employments"
	} else {
		table = "work_experience_letters"
	}

	// Hapus Surat dari database
	where := fmt.Sprintf("number_of_employees = '%s' ", payload.Number_of_employees)
	repository.DeleteResign(table, where)

	// Tampung data yang mau di tamplikan
	result := map[string]interface{}{
		"code":    200,
		"data":    payload.Number_of_employees,
		"message": "Delete Success",
	}

	// encode ke json berbentuk byte
	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Respon json
	w.Write([]byte(resp))
}

func ExportLetter(w http.ResponseWriter, r *http.Request) {
	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	vars := mux.Vars(r)
	dataletter, _ := vars["dataletter"]

	var table, letter_Date, letter_No string

	if dataletter == "certificate_of_employments" {
		table = "certificate_of_employments"
		letter_Date = "date_certificate_employee"
		letter_No = "no_certificate_employee"
	} else {
		table = "work_experience_letters"
		letter_Date = "date_letter_experience"
		letter_No = "no_letter_experience"
	}

	query := fmt.Sprintf("select resign_id, COALESCE(%s, '0000-00-00'), COALESCE(%s, 0), COALESCE(rom, ''), COALESCE(created_at, '0000-00-00 00:00:00'), COALESCE(updated_at, '0000-00-00 00:00:00') from %s", letter_Date, letter_No, table)
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
	xlsx.SetCellValue(sheet1Name, "G1", "NOMOR SURAT")
	xlsx.SetCellValue(sheet1Name, "H1", "ROM")
	xlsx.SetCellValue(sheet1Name, "I1", "TAHUN")
	xlsx.SetCellValue(sheet1Name, "J1", "TYPE")
	xlsx.SetCellValue(sheet1Name, "K1", "STATUS SURAT")
	xlsx.SetCellValue(sheet1Name, "L1", "CLASSIFIKASI")
	xlsx.SetCellValue(sheet1Name, "M1", "UMUR")
	xlsx.SetCellValue(sheet1Name, "N1", "CREATEAD AT")
	xlsx.SetCellValue(sheet1Name, "O1", "UPDATED AT")

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
		go func(wg *sync.WaitGroup, no int, Resign_id int, Date string, No int, Rom string, Created_at string, Updated_at string) {
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
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", no), No)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", no), Rom)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("I%d", no), helper.YearMysql(Date))
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", no), Resign.Type)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", no), Resign.Status_resign)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("L%d", no), Resign.Classification)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("M%d", no), Resign.Age)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("N%d", no), Resign.Created_at)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("O%d", no), Resign.Updated_at)
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

func SearchLetter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", config.Url_web)
	w.Header().Add("Access-Control-Allow-Headers", "*")

	if r.Method == "POST" {
		dbresign, err := models.ConnResign()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer dbresign.Close()

		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Date_first       string `json:"date_first"`
			Date_last        string `json:"date_last"`
			Get_value_resign string `json:"get_value_resign"`
			Selectcoloumn    string `json:"selectcoloumn"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sqlsearch := fmt.Sprintf(`SELECT 
		resigns.number_of_employees, resigns.name, %s.date_%s FROM %s JOIN resigns ON %s.resign_id = resigns.id 
		WHERE 
		%s.date_%s BETWEEN '%s' AND '%s'`,
			payload.Selectcoloumn, payload.Get_value_resign, payload.Selectcoloumn, payload.Selectcoloumn, payload.Selectcoloumn, payload.Get_value_resign, payload.Date_first, payload.Date_last)

		rows, err := dbresign.Query(sqlsearch)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		var tr string = ""
		ind := 0
		for rows.Next() {
			var each = models.Resign{}
			var Date_search string
			var err = rows.Scan(&each.Number_of_employees, &each.Name, &Date_search)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			ind += 1
			tr = fmt.Sprintf(`
				%s<tr>
					<td>%d.</td>
					<td>%s</td>
					<td>%s</td>
					<td>%s</td>
					<td class="text-center">
						<span>
							<input class="form-check-input checkboxletter" value="%s" name="check[]" id="%s" type="checkbox" checked="checked">
						</span>
					</td>
					</tr>
					`, tr, ind, each.Number_of_employees, each.Name, Date_search, each.Number_of_employees, each.Number_of_employees)
		}
		tbody := fmt.Sprintf(`<div class="card">
				<div class="card-header">
				<div class="custom-control custom-checkbox">
					<input class="custom-control-input" type="checkbox" id="checklistallletter" checked="checked" value="checkall" onclick="CheckboxLetter();">
					<label for="checklistallletter" class="custom-control-label"> Checklist All</label>
				</div>
				</div>
				<div class="card-body p-0">
					<table class="table table-sm">
						<thead>
						<tr>
							<th>NO</th>
							<th>NIK</th>
							<th>Nama</th>
							<th>Tanggal</th>
							<th style="width: 10px;">Check</th>
						</tr>
						</thead>
						<tbody> %s %s`, tr, `</tbody>
				</table>
			</div>
		</div>`)
		resp, err := json.Marshal(tbody)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		w.Write([]byte(resp))
		return
	}
	message := http.StatusBadRequest
	resp, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Write([]byte(resp))
	return
}
func ProcessAccLetter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", config.Url_web)
	w.Header().Add("Access-Control-Allow-Headers", "*")

	if r.Method == "POST" {
		dbresign, err := models.ConnResign()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer dbresign.Close()

		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Data []string `json:"data"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(payload.Data) == 0 {
			message := fmt.Sprint("Tidak Ada Karyawan yang di Acc")
			fmt.Println(message)
			resp, err := json.Marshal(message)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			w.Write([]byte(resp))
			return
		}

		var data []interface{}
		for i := 0; i < len(payload.Data); i++ {
			data = append(data, DataLetter(payload.Data[i], "q", "2"))
		}
		result := map[string]interface{}{
			"data":    data,
			"code":    200,
			"message": "Successfully",
		}
		resp, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		w.Write([]byte(resp))
		return
	}
	message := http.StatusBadRequest
	resp, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Write([]byte(resp))
	return
}

func PrintLetter(w http.ResponseWriter, r *http.Request) {
	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()
	decoder := json.NewDecoder(r.Body)
	payload := struct {
		Number_of_employees []string `json:"number_of_employees"`
		Table               string   `json:"table"`
		Column              string   `json:"column"`
	}{}
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data []interface{}
	for i := 0; i < len(payload.Number_of_employees); i++ {
		data = append(data, DataLetter(payload.Number_of_employees[i], payload.Table, payload.Column))
	}
	result := map[string]interface{}{
		"data":    data,
		"code":    200,
		"message": "Successfully",
	}
	resp, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Write([]byte(resp))
	return
}

func DataLetter(number_of_employees string, table string, column string) map[string]interface{} {

	dbresign, _ := models.ConnResign()
	defer dbresign.Close()

	db, _ := models.ConnHrd()
	defer db.Close()

	var Resign_id int
	var ResignSel = models.Resign{}
	_ = dbresign.QueryRow("SELECT id as resign_id, name, COALESCE(position, ''), COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') FROM resigns WHERE number_of_employees = ? ", number_of_employees).
		Scan(&Resign_id, &ResignSel.Name, &ResignSel.Position, &ResignSel.Department, &ResignSel.Hire_date, &ResignSel.Date_out)

	var letters = models.Letter{}
	queryletter := fmt.Sprintf(`SELECT id, resign_id, number_of_employees, date_%s, no_%s, rom, created_at, updated_at FROM %s WHERE number_of_employees = '%s' `, column, column, table, number_of_employees)
	_ = dbresign.QueryRow(queryletter).
		Scan(&letters.Id, &letters.Resign_id, &letters.Number_of_employees, &letters.Date, &letters.No, &letters.Rom, &letters.Created_at, &letters.Updated_at)

	_ = db.QueryRow("SELECT COALESCE(date_out, '0000-00-00') FROM employees WHERE number_of_employees = ?", number_of_employees).
		Scan(&letters.Date)

	dataletter := map[string]interface{}{
		"number_of_employees": number_of_employees,
		"name":                ResignSel.Name,
		"position":            ResignSel.Position,
		"department":          helper.Deptout(ResignSel.Department),
		"hire_date":           ResignSel.Hire_date,
		"date_out":            ResignSel.Date_out,
		"date":                letters.Date,
		"no":                  letters.No,
		"rom":                 letters.Rom,
		"created_at":          letters.Created_at,
		"updated_at":          letters.Updated_at,
	}
	return dataletter
}
