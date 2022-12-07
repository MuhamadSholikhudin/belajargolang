package controllers

import (
	"belajargolang/api/resign/helper"
	"belajargolang/api/resign/models"
	"belajargolang/api/resign/repository"
	"bufio"
	"bytes"
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

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
)

func Submissions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnResign()
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
		sqlPaging = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%' ORDER BY id DESC", sqlPaging, justStringnumber_of_employees)
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

	var submission []models.Resignation_submission

	for rows.Next() {
		var each = models.Resignation_submission{}
		var err = rows.Scan(&each.Number_of_employees, &each.Name, &each.Position, &each.Department, &each.Building, &each.Hire_date, &each.Date_out, &each.Date_resignation_submissions, &each.Type, &each.Reason, &each.Detail_reason, &each.Periode_of_service, &each.Age, &each.Suggestion, &each.Status_resignsubmisssion, &each.Using_media, &each.Classification, &each.Created_at, &each.Updated_at)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		submission = append(submission, each)
	}

	links := map[string]interface{}{
		"first": fmt.Sprintf("/resigns_submission?page=%s%s", first, params),
		"last":  fmt.Sprintf("/resigns_submission?page=%s%s", last, params),
		"next":  fmt.Sprintf("/resigns_submission?page=%s%s", next, params),
		"prev":  fmt.Sprintf("/resigns_submission?page=%s%s", prev, params),
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

	db, err := models.ConnResign()
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

	var result []models.Resignation_submission

	for rows.Next() {
		var each = models.Resignation_submission{}
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

	db, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var dbresign, _ = models.ConnResign()
	defer dbresign.Close()

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
			err = db.QueryRow("select name, COALESCE(date_of_birth, '0000-00-00'), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00') as date_out, status_employee from employees where number_of_employees = ? ", Number_of_employees).
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
				_, err = dbresign.Exec("INSERT INTO `resignation_submissions` (`number_of_employees`, `name`, `position`, `department`, `building`, `address`, `hire_date`, `date_out`, `date_resignation_submissions`, `type`, `reason`, `detail_reason`, `periode_of_service`, `age`, `suggestion`, `status_resignsubmisssion`, `using_media`, `classification`, `print`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )", Number_of_employees, resultemployee.Name, record[5], record[9], record[9], record[3], resultemployee.Hire_date, resultemployee.Date_out, record[4], "false", record[3], record[6], helper.Periode_of_serve(resultemployee.Hire_date, record[4]), helper.Age(resultemployee.Date_of_birth), record[7], "wait", "google", "Mengajukan permohonan resign setelah karyawan resign", 0, record[0], record[0])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				err = dbresign.QueryRow("SELECT id FROM resignation_submissions WHERE number_of_employees = ? AND created_at = ? ", Number_of_employees, record[0]).
					Scan(&resignation_submission_id)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				_, err = dbresign.Exec("INSERT INTO `kuesioners`( `resignation_submission_id`, `number_of_employees`, `k1`, `k2`, `k3`, `k4`, `k5`, `k6`, `k7`, `created_at`, `updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ", resignation_submission_id, Number_of_employees, record[10], record[11], record[12], record[13], record[14], record[15], record[16], record[0], record[0])
				if err != nil {
					fmt.Println(err.Error())
					return
				}

			} else if Count_resigns == 0 && Count_resign_submissions == 0 {
				// TIdak resign dan belum mengajukan resign maka boleh mengajukan resign
				_, err = dbresign.Exec("INSERT INTO `resignation_submissions` (`number_of_employees`, `name`, `position`, `department`, `building`, `address`, `hire_date`, `date_out`, `date_resignation_submissions`, `type`, `reason`, `detail_reason`, `periode_of_service`, `age`, `suggestion`, `status_resignsubmisssion`, `using_media`, `classification`, `print`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,? )", Number_of_employees, resultemployee.Name, record[5], record[9], record[9], record[3], resultemployee.Hire_date, resultemployee.Date_out, record[4], "true", record[3], record[6], helper.Periode_of_serve(resultemployee.Hire_date, record[4]), helper.Age(resultemployee.Date_of_birth), record[7], "wait", "google", "Mengajukan permohonan resign setelah karyawan resign", 0, record[0], record[0])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				err = dbresign.QueryRow("SELECT id FROM resignation_submissions WHERE number_of_employees = ? AND created_at = ? ", Number_of_employees, record[0]).
					Scan(&resignation_submission_id)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				_, err = dbresign.Exec("INSERT INTO `kuesioners`(`resignation_submission_id`, `number_of_employees`, `k1`, `k2`, `k3`, `k4`, `k5`, `k6`, `k7`, `created_at`, `updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ", resignation_submission_id, Number_of_employees, record[10], record[11], record[12], record[13], record[14], record[15], record[16], record[0], record[0])
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

	var dbresign, err = models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	var Submission models.Resignation_submission

	err = dbresign.QueryRow("SELECT number_of_employees, COALESCE(name, ''), COALESCE(position, ''), COALESCE(department, ''), COALESCE(building, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), COALESCE(date_resignation_submissions, '0000-00-00'), COALESCE(type, ''), COALESCE(reason, ''), COALESCE(detail_reason, ''), COALESCE(suggestion, ''), COALESCE(periode_of_service, 0), COALESCE(status_resignsubmisssion, ''), COALESCE(age, 0), COALESCE(using_media, ''), COALESCE(classification, ''), COALESCE(created_at, '0000-00-00 00:00:00'), COALESCE(updated_at, '0000-00-00 00:00:00')  FROM resignation_submissions WHERE number_of_employees = ? ", Number_of_employess).
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

	data := models.Resignation_submission{}
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

	_, err = dbresign.Exec("UPDATE `resignation_submissions` SET `name`= ? ,`position`= ? ,`department`=  ? , `hire_date`= ? ,`date_out`= ? ,`date_resignation_submissions`= ? ,`type`= ? ,`reason`= ? ,`detail_reason`= ? ,`periode_of_service`= ? ,`age`= ? ,`suggestion`= ? ,`status_resignsubmisssion`= ? ,`using_media`= ? ,`classification`= ? ,`created_at`= ? ,`updated_at`= ?  WHERE number_of_employees = ? ", data.Name, data.Position, data.Department, data.Hire_date, data.Date_out, data.Date_resignation_submissions, data.Type, data.Reason, data.Detail_reason, helper.Periode_of_serve(data.Hire_date, data.Date_resignation_submissions), data.Age, data.Suggestion, data.Status_resignsubmisssion, data.Using_media, data.Classification, data.Created_at, data.Updated_at, data.Number_of_employees)
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

func GetResignSubmissionStatus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	// nik, _ := strconv.Atoi(vars["number_of_employees"])
	// status_api, _ := strconv.Atoi(vars["status_resign"])
	number_of_employees := vars["number_of_employees"]
	status_resign := vars["status_resign"]

	var db, err = models.ConnHrd()
	defer db.Close()

	var dbresign, _ = models.ConnResign()
	defer dbresign.Close()

	var statushttp string

	if status_resign == "acc" {

		// Update data resign status acc
		_, err = dbresign.Exec("UPDATE `resignation_submissions` SET `status_resignsubmisssion`= ? WHERE number_of_employees = ? ", status_resign, number_of_employees)
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
		_, err = dbresign.Exec(cancel_resign)
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

func PostStatus(w http.ResponseWriter, r *http.Request) {
	//Untuk membuat json pertama kita harus set Header
	w.Header().Set("Content-Type", "application/json")

	//Mendecode requset body langsung menjadi json
	data := models.Resign{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	ProsessSubmission(data.Number_of_employees, data.Status_resign)

	/*
		var status_employee, status_submission string
		switch data.Status_resign {
		case "acc":
			status_employee = "notactive"
			status_submission = "acc"
		case "cancel_submission":
			var Scanstatus_employee string
			err = db.QueryRow("SELECT status_employee FROM employees WHERE number_of_employees = ?", data.Number_of_employees).Scan(&Scanstatus_employee)
			status_employee = Scanstatus_employee
			status_submission = "cancel"
		case "cancel_and_active":
			status_employee = "active"
			status_submission = "cancel"
		}

		// Update data resign status acc
		_, err = dbresign.Exec("UPDATE `resignation_submissions` SET `status_resignsubmisssion`= ? WHERE number_of_employees = ? ORDER BY id DESC", status_submission, data.Number_of_employees)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Update employees
		_, err = db.Exec("UPDATE `employees` SET `status_employee`= ? WHERE number_of_employees = ? ", status_employee, data.Number_of_employees)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		switch data.Status_resign {
		case "acc":
			// Select resign submissions
			var Submission = models.Resignation_submission{}
			err = dbresign.QueryRow("SELECT number_of_employees,	name,	position,	department,	building,	hire_date,	COALESCE(date_out, '0000-00-00') as date_out, COALESCE(date_resignation_submissions, '0000-00-00'),	type, reason, detail_reason, periode_of_service, age, suggestion,	status_resignsubmisssion,	using_media, classification FROM resignation_submissions WHERE number_of_employees = ? ORDER BY id DESC", data.Number_of_employees).
				Scan(&Submission.Number_of_employees, &Submission.Name, &Submission.Position, &Submission.Department, &Submission.Building, &Submission.Hire_date, &Submission.Date_out, &Submission.Date_resignation_submissions, &Submission.Type, &Submission.Reason, &Submission.Detail_reason, &Submission.Periode_of_service, &Submission.Age, &Submission.Suggestion, &Submission.Status_resignsubmisssion, &Submission.Using_media, &Submission.Classification)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var Count_resign_id int
			err = dbresign.QueryRow("SELECT COUNT(id) FROM resigns WHERE number_of_employees = ?", data.Number_of_employees).
				Scan(&Count_resign_id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if Count_resign_id == 0 {
				//Insert data resigns
				_, err = dbresign.Exec("INSERT INTO `resigns`(	`number_of_employees`,`name`, `position`, `department`, `hire_date`, `classification`, `date_out`, `date_resignsubmissions`, `periode_of_service`, `type`, `age`, `status_resign`, `printed`, `created_at`, `updated_at`) VALUES (?,	?,	?,	?,	?,	?,	?,	?,	?,	?,	?, ? , ?, ?, ?)", Submission.Number_of_employees, Submission.Name, Submission.Position, Submission.Department, Submission.Hire_date, helper.CekDateSubmission(data.Number_of_employees), Submission.Date_resignation_submissions, Submission.Date_resignation_submissions, Submission.Periode_of_service, Submission.Type, Submission.Age, data.Status_resign, 0, helper.DMYhms(), helper.DMYhms())
				if err != nil {
					fmt.Println(err.Error())
					return
				}
			} else {
				//Insert data resigns
				_, err = dbresign.Exec("UPDATE resigns SET status_resign = ? WHERE number_of_employees = ?", data.Status_resign, data.Number_of_employees)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
			}
		case "cancel_and_active":
			_, err = dbresign.Exec("DELETE FROM resigns WHERE number_of_employees = ?", data.Number_of_employees)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			_, err = dbresign.Exec("DELETE FROM certificate_of_employments WHERE number_of_employees = ?", data.Number_of_employees)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			_, err = dbresign.Exec("DELETE FROM work_experience_letters WHERE number_of_employees = ?", data.Number_of_employees)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	*/

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

func ExportSubmission(w http.ResponseWriter, r *http.Request) {
	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	rows, err := dbresign.Query("select id, number_of_employees from resignation_submissions where id > ?", 0)
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

	xlsx.SetCellValue(sheet1Name, "S1", "1. Saya terampil menyelesaikan target pekerjaan")
	xlsx.SetCellValue(sheet1Name, "T1", "2. Atasan menggunakan kata-kata/sikap yang wajar dalam bekerja")
	xlsx.SetCellValue(sheet1Name, "U1", "3. Rekan kerja saya membantu kesulitan saya dalam menyelesaikan pekerjaan")
	xlsx.SetCellValue(sheet1Name, "V1", "4. Jarak perusahaan dengan tempat tinggal tidak menjadi masalah bagi saya")
	xlsx.SetCellValue(sheet1Name, "W1", "5. Jam kerja (termasuk shift malam) tidak masalah bagi saya")
	xlsx.SetCellValue(sheet1Name, "X1", "6. Saya berkeinginan kembali ke perusahaan (PT HWI) suatu saat nanti")
	xlsx.SetCellValue(sheet1Name, "Y1", "7. Keluarga (termasuk menikah, mengurus keluarga) bukanlah alasan bagi saya untuk meninggalkan perusahaan ini")
	xlsx.SetCellValue(sheet1Name, "Z1", "")

	var wg sync.WaitGroup

	no := 1

	for rows.Next() {
		var NIK string
		var Submission_id int
		var err = rows.Scan(&Submission_id, &NIK)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		no += 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, Submission_id int, message string, no int) {
			defer wg.Done()
			var Submission models.Resignation_submission

			err = dbresign.QueryRow("SELECT COALESCE(number_of_employees, ''),	COALESCE(name, ''),	COALESCE(position, ''),	COALESCE(department, ''),	COALESCE(building, ''),	COALESCE(hire_date, '0000-00-00'),	COALESCE(date_out, '0000-00-00'),	COALESCE(date_resignation_submissions, '0000-00-00'),	COALESCE(type, ''),	COALESCE(reason, ''),	COALESCE(detail_reason, ''),	COALESCE(periode_of_service, ''),	COALESCE(age, 0),	COALESCE(suggestion, ''),	COALESCE(status_resignsubmisssion, ''),	COALESCE(using_media, ''),	COALESCE(classification, ''),	COALESCE(created_at, '0000-00-00 00:00:00'),	COALESCE(updated_at, '0000-00-00 00:00:00')	from resignation_submissions where number_of_employees = ?", message).
				Scan(&Submission.Number_of_employees, &Submission.Name, &Submission.Position, &Submission.Department, &Submission.Building, &Submission.Hire_date, &Submission.Date_out, &Submission.Date_resignation_submissions, &Submission.Type, &Submission.Reason, &Submission.Detail_reason, &Submission.Periode_of_service, &Submission.Age, &Submission.Suggestion, &Submission.Status_resignsubmisssion, &Submission.Using_media, &Submission.Classification, &Submission.Created_at, &Submission.Updated_at)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var k1, k2, k3, k4, k5, k6, k7 int

			err = dbresign.QueryRow("SELECT k1, k2, k3, k4, k5, k6, k7 FROM kuesioners WHERE resignation_submission_id = ?", Submission_id).
				Scan(&k1, &k2, &k3, &k4, &k5, &k6, &k7)
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

			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("N%d", no), k1)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("O%d", no), k2)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("P%d", no), k3)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("Q%d", no), k4)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("R%d", no), k5)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("S%d", no), k6)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("T%d", no), k7)

		}(&wg, Submission_id, NIK, no)
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

func SearchSubmission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "http://127.0.0.1:8000")
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
			Date_resignation_submission string `json:"date_resignation_submissions"`
			Selectdatesubmission        string `json:"selectdatesubmission"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sqlsearch := fmt.Sprintf("SELECT number_of_employees, name, %s, status_resignsubmisssion FROM resignation_submissions WHERE %s LIKE '%%%s%%' AND status_resignsubmisssion = 'wait' ", payload.Selectdatesubmission, payload.Selectdatesubmission, payload.Date_resignation_submission)
		rows, err := dbresign.Query(sqlsearch)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		var tr string = ""
		ind := 0
		for rows.Next() {
			var each = models.Resignation_submission{}
			var Date_search string
			var err = rows.Scan(&each.Number_of_employees, &each.Name, &Date_search, &each.Status_resignsubmisssion)
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
							<input class="form-check-input checkboxsubmission" id="%s" type="checkbox"
								checked="checked">
						</span>
					</td>
					</tr>
					`, tr, ind, each.Number_of_employees, each.Name, Date_search, each.Number_of_employees)
		}

		tbody := fmt.Sprintf(`<div class="card">
				<div class="card-header">
				<div class="custom-control custom-checkbox">
					<input class="custom-control-input" type="checkbox" id="checklistallsubmission" checked="checked" value="checkall" onclick="CheckboxSubmission();">
					<label for="checklistallsubmission" class="custom-control-label"> Checklist All</label>
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

	// http.Error(w, "Only accept POST request", http.StatusBadRequest)
	message := http.StatusBadRequest

	resp, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Write([]byte(resp))
	return
}

func ProcessAccSubmission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "http://127.0.0.1:8000")
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
			message := fmt.Sprint(" Tidak Ada Karyawan yang di Acc")
			fmt.Println(message)
			resp, err := json.Marshal(message)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			w.Write([]byte(resp))
			return
		}

		for i := 0; i < len(payload.Data); i++ {
			ProsessSubmission(payload.Data[i], "acc")
		}
		message := fmt.Sprint(len(payload.Data), " Karyawan Berhasil di Acc")
		resp, err := json.Marshal(message)
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

func ProsessSubmission(number_of_employees string, status_resign string) {

	dbhrd, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhrd.Close()

	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	switch status_resign {
	case "acc":
		// Select resign submissions
		var Submission = models.Resignation_submission{}
		err = dbresign.QueryRow("SELECT number_of_employees,	name,	position,	department,	building,	hire_date,	COALESCE(date_out, '0000-00-00') as date_out, COALESCE(date_resignation_submissions, '0000-00-00'),	type, reason, detail_reason, periode_of_service, age, suggestion,	status_resignsubmisssion,	using_media, classification FROM resignation_submissions WHERE number_of_employees = ? ORDER BY id DESC", number_of_employees).
			Scan(&Submission.Number_of_employees, &Submission.Name, &Submission.Position, &Submission.Department, &Submission.Building, &Submission.Hire_date, &Submission.Date_out, &Submission.Date_resignation_submissions, &Submission.Type, &Submission.Reason, &Submission.Detail_reason, &Submission.Periode_of_service, &Submission.Age, &Submission.Suggestion, &Submission.Status_resignsubmisssion, &Submission.Using_media, &Submission.Classification)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var Count_resign_id int
		err = dbresign.QueryRow("SELECT COUNT(id) FROM resigns WHERE number_of_employees = ?", number_of_employees).
			Scan(&Count_resign_id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if Count_resign_id == 0 {
			var data = map[string]interface{}{
				"number_of_employees":    Submission.Number_of_employees,
				"name":                   Submission.Name,
				"position":               Submission.Position,
				"department":             Submission.Department,
				"hire_date":              Submission.Hire_date,
				"classification":         helper.TypeResign(number_of_employees, Submission.Date_resignation_submissions)["classification"],
				"date_out":               Submission.Date_resignation_submissions,
				"date_resignsubmissions": Submission.Date_resignation_submissions,
				"periode_of_service":     Submission.Periode_of_service,
				"type":                   helper.TypeResign(number_of_employees, Submission.Date_resignation_submissions)["type"],
				"age":                    Submission.Age,
				"status_resign":          "wait",
				"printed":                0,
				"created_at":             helper.DMYhms(),
				"updated_at":             helper.DMYhms(),
			}
			repository.InsertResign("resigns", data)

			var data1 = map[string]interface{}{
				"date_out":        Submission.Date_resignation_submissions,
				"status_employee": "notactive",
				"exit_statement":  Submission.Reason,
			}
			where1 := fmt.Sprintf("number_of_employees = '%s' ", number_of_employees)
			repository.UpdateHrd("employees", data1, where1)

			var data2 = map[string]interface{}{
				"status_resignsubmisssion": status_resign,
				"date_out":                 Submission.Date_resignation_submissions,
			}
			where2 := fmt.Sprintf("number_of_employees = '%s' AND status_resignsubmisssion = 'wait' ", number_of_employees)
			repository.UpdateResign("resignation_submissions", data2, where2)

		} else {

			var date_out string
			err = dbhrd.QueryRow("SELECT COALESCE(date_out, '0000-00-00') FROM employees WHERE number_of_employees = ?", number_of_employees).
				Scan(&date_out)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var data = map[string]interface{}{
				"status_resignsubmisssion": status_resign,
				"date_out":                 date_out,
			}
			where := fmt.Sprintf("number_of_employees = '%s' AND status_resignsubmisssion = 'wait' ", number_of_employees)
			repository.UpdateResign("resignation_submissions", data, where)

			var data2 = map[string]interface{}{
				"type":                   helper.TypeResign(number_of_employees, date_out)["type"],
				"classification":         helper.TypeResign(number_of_employees, date_out)["classification"],
				"status_resign":          helper.TypeResign(number_of_employees, date_out)["status"],
				"date_resignsubmissions": Submission.Date_resignation_submissions,
			}
			where2 := fmt.Sprintf("number_of_employees = '%s' ", number_of_employees)
			repository.UpdateResign("resigns", data2, where2)
		}

	case "cancel_and_active":

		where := fmt.Sprintf("number_of_employees = '%s' ", number_of_employees)
		repository.DeleteResign("work_experience_letters", where)
		repository.DeleteResign("certificate_of_employments", where)
		repository.DeleteResign("resigns", where)

		var data1 = map[string]interface{}{
			"date_out":        nil,
			"status_employee": "active",
			"exit_statement":  nil,
		}
		where1 := fmt.Sprintf("number_of_employees = '%s' ", number_of_employees)
		repository.UpdateHrd("employees", data1, where1)

		var data2 = map[string]interface{}{
			"status_resignsubmisssion": "wait",
		}
		where2 := fmt.Sprintf("number_of_employees = '%s' AND status_resignsubmisssion = 'acc' ", number_of_employees)
		repository.UpdateResign("resignation_submissions", data2, where2)

	case "cancel_submission":

		var data = map[string]interface{}{
			"status_resignsubmisssion": "cancel",
		}
		where := fmt.Sprintf("number_of_employees = '%s' AND status_resignsubmisssion = 'wait' ", number_of_employees)
		repository.UpdateResign("resignation_submissions", data, where)

		var data2 = map[string]interface{}{
			"date_out":        nil,
			"status_employee": "active",
			"exit_statement":  nil,
		}
		where2 := fmt.Sprintf("number_of_employees = '%s' ", number_of_employees)
		repository.UpdateHrd("employees", data2, where2)
	}

}
