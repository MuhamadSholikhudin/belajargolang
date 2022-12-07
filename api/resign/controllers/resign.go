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
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
)

func Resigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := models.ConnResign()
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
		sqlPaging = fmt.Sprintf("%s WHERE number_of_employees LIKE '%%%s%%' ORDER BY id DESC", sqlPaging, justStringnumber_of_employees)
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

	var resign []models.Resign

	for rows.Next() {
		var each = models.Resign{}
		var err = rows.Scan(&each.Id, &each.Number_of_employees, &each.Name, &each.Hire_date, &each.Classification, &each.Date_out, &each.Date_resignsubmissions, &each.Periode_of_service, &each.Position, &each.Department, &each.Type, &each.Age, &each.Status_resign, &each.Printed, &each.Created_at, &each.Updated_at)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		resign = append(resign, each)
	}

	links := map[string]interface{}{
		"first": fmt.Sprintf("/resigns_resign?page=%s%s", first, params),
		"last":  fmt.Sprintf("/resigns_resign?page=%s%s", last, params),
		"next":  fmt.Sprintf("/resigns_resign?page=%s%s", next, params),
		"prev":  fmt.Sprintf("/resigns_resign?page=%s%s", prev, params),
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

	var dbresign, err = models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	var Resign models.Resign

	err = dbresign.QueryRow("SELECT number_of_employees, COALESCE(name, ''), COALESCE(position, ''), COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), COALESCE(date_resignsubmissions, '0000-00-00'), COALESCE(type, ''), COALESCE(periode_of_service, 0), COALESCE(status_resign, ''), COALESCE(age, 0), COALESCE(classification, ''), COALESCE(created_at, '0000-00-00 00:00:00'), COALESCE(updated_at, '0000-00-00 00:00:00')  FROM resigns WHERE number_of_employees = ? ", Number_of_employess).
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

func PutResign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	data := models.Resign{}
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

	_, err = dbresign.Exec("UPDATE `resigns` SET `name`= ? ,`position`= ? ,`department`=  ? , `hire_date`= ? ,`date_out`= ? ,`date_resignsubmissions`= ? ,`type`= ? , `periode_of_service`= ? ,`age`= ? ,`status_resign`= ? , `classification`= ? ,`created_at`= ? ,`updated_at`= ?  WHERE number_of_employees = ? ", data.Name, data.Position, data.Department, data.Hire_date, data.Date_out, data.Date_resignsubmissions, data.Type, helper.Periode_of_serve(data.Hire_date, data.Date_out), data.Age, data.Status_resign, data.Classification, data.Created_at, data.Updated_at, data.Number_of_employees)
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

func UploadResigns(w http.ResponseWriter, r *http.Request) {

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
	var Employee = models.Employee{}
	var Count_id int
	for _, record := range records {
		err = db.QueryRow("SELECT COUNT(id) as id , COALESCE(status_employee, '') as status_employee, COALESCE(name, ''), COALESCE(hire_date, ''), COALESCE(date_of_birth, ''), COALESCE(date_out, '0000-00-00'),COALESCE(job_id, 25),COALESCE(department_id, 116), COALESCE(address_jalan, ''), COALESCE(address_rt, ''), COALESCE(address_rw, ''), COALESCE(address_village, ''), COALESCE(address_district, ''), COALESCE(address_city, ''), COALESCE(address_province, '') FROM employees WHERE number_of_employees = ? ", record[0]).
			Scan(&Count_id, &Employee.Status_employee, &Employee.Name, &Employee.Hire_date, &Employee.Date_of_birth, &Employee.Date_out, &Employee.Job_id, &Employee.Department_id, &Employee.Address_jalan, &Employee.Address_rt, &Employee.Address_rw, &Employee.Address_village, &Employee.Address_district, &Employee.Address_city, &Employee.Address_province)
		if err != nil {
			fmt.Println(err.Error())
		}
		if Count_id > 0 {
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
		}

		var Count_idresigns = 0
		err = dbresign.QueryRow("SELECT COUNT(id) as id FROM resigns WHERE number_of_employees = ? ", record[0]).
			Scan(&Count_idresigns)
		if err != nil {
			fmt.Println(err.Error())
		}
		if record[0] == "number_of_employees" {
		} else if Count_idresigns < 1 && Count_id > 0 {
			_, err = dbresign.Exec("INSERT INTO `resigns`(	`number_of_employees`,`name`, `position`, `department`, `hire_date`, `classification`, `date_out`, `date_resignsubmissions`, `periode_of_service`, `type`, `age`, `status_resign`, `printed`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?,	?,	?,	?,	?,	?,	?,	?,	?,	?, ? , ?)", record[0], Employee.Name, JobDepartment(record[0])[0], JobDepartment(record[0])[1], Employee.Hire_date, helper.CekDateSubmission(record[0]), record[2], nil, helper.Periode_of_serve(Employee.Hire_date, record[2]), helper.TypeResign(record[0], record[2])["type"], helper.Age(Employee.Date_of_birth), "resign", 0, helper.DMYhms(), helper.DMYhms())
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

func ExportResign(w http.ResponseWriter, r *http.Request) {
	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	rows, err := dbresign.Query("select number_of_employees from resigns")
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
	xlsx.SetCellValue(sheet1Name, "E1", "HIRE DATE =DATE(LEFT(F2,4), MID(F2,6,2), RIGHT(F2,2))")
	xlsx.SetCellValue(sheet1Name, "F1", "DATE OUT =DATE(LEFT(F2,4), MID(F2,6,2), RIGHT(F2,2))")
	xlsx.SetCellValue(sheet1Name, "G1", "TGL PERMOHONAN =DATE(LEFT(G2,4), MID(G2,6,2), RIGHT(G2,2))")
	xlsx.SetCellValue(sheet1Name, "H1", "TYPE")
	xlsx.SetCellValue(sheet1Name, "I1", "UMUR")
	xlsx.SetCellValue(sheet1Name, "J1", "STATUS RESIGN")
	xlsx.SetCellValue(sheet1Name, "K1", "CLASSIFIKASI")
	xlsx.SetCellValue(sheet1Name, "L1", "CREATEAD AT")
	xlsx.SetCellValue(sheet1Name, "M1", "UPDATED AT")

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
			var Resign models.Resign

			err = dbresign.QueryRow("SELECT COALESCE(number_of_employees, ''),	COALESCE(name, ''),	COALESCE(position, ''),	COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'),	COALESCE(date_out, '0000-00-00'),	COALESCE(date_resignsubmissions, '0000-00-00'),	COALESCE(type, ''),	COALESCE(age, 0),	COALESCE(status_resign, ''),	COALESCE(classification, ''),	COALESCE(created_at, '0000-00-00 00:00:00'),	COALESCE(updated_at, '0000-00-00 00:00:00')	from resigns where number_of_employees = ?", message).
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
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", no), Resign.Date_resignsubmissions)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", no), Resign.Type)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("I%d", no), Resign.Age)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", no), Resign.Status_resign)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", no), Resign.Classification)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("L%d", no), Resign.Created_at)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("M%d", no), Resign.Updated_at)
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

func Searchresign(w http.ResponseWriter, r *http.Request) {
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
			Date_out      string `json:"date_out"`
			Selectcoloumn string `json:"selectcoloumn"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sqlsearch := fmt.Sprintf("SELECT number_of_employees, name, %s, status_resign FROM resigns WHERE %s LIKE '%%%s%%' AND status_resign = 'wait' ", payload.Selectcoloumn, payload.Selectcoloumn, payload.Date_out)
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
			var err = rows.Scan(&each.Number_of_employees, &each.Name, &Date_search, &each.Status_resign)
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
							<input class="form-check-input checkboxresign" id="%s" type="checkbox" checked="checked">
						</span>
					</td>
					</tr>
					`, tr, ind, each.Number_of_employees, each.Name, Date_search, each.Number_of_employees)
		}

		tbody := fmt.Sprintf(`<div class="card">
				<div class="card-header">
				<div class="custom-control custom-checkbox">
					<input class="custom-control-input" type="checkbox" id="checklistallresign" checked="checked" value="checkall" onclick="CheckboxResign();">
					<label for="checklistallresign" class="custom-control-label"> Checklist All</label>
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

func ProcessAccResign(w http.ResponseWriter, r *http.Request) {
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
			AccResign(payload.Data[i], "acc")
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

func AccResign(number_of_employees string, status_resign string) {

	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	// Update Resign
	var data = map[string]interface{}{
		"status_resign": status_resign,
	}
	where := fmt.Sprintf("number_of_employees = '%s' AND status_resign = 'wait' ", number_of_employees)
	repository.UpdateResign("resigns", data, where)

	var Resign_id int
	var resign models.Resign
	err = dbresign.QueryRow("SELECT id as resign_id, name, COALESCE(position, ''), COALESCE(department, ''), COALESCE(hire_date, '0000-00-00'), COALESCE(date_out, '0000-00-00'), periode_of_service, type FROM resigns WHERE number_of_employees = ? ", number_of_employees).
		Scan(&Resign_id, &resign.Name, &resign.Position, &resign.Department, &resign.Hire_date, &resign.Date_out, &resign.Periode_of_service, &resign.Type)
	if err != nil {
		fmt.Print(err.Error())
	}

	var yearstring string
	yearstring = strconv.Itoa(time.Now().Year())

	var CountCertificateByDate, CountNoCertificateEmployee int
	err = dbresign.QueryRow("SELECT COUNT(id) as CountCertificateByDate, COALESCE(no_certificate_employee, 0) as no_certificate_employee FROM certificate_of_employments WHERE YEAR(date_certificate_employee) = ? AND MONTH(date_certificate_employee) = ? ORDER BY date_certificate_employee DESC", yearstring, helper.StringMonth()).
		Scan(&CountCertificateByDate, &CountNoCertificateEmployee)
	if err != nil {
		fmt.Print(err.Error())
	}

	var CountExperienceByDate, CountNoExperienceEmployee int
	err = dbresign.QueryRow("SELECT COUNT(id) as CountExperienceByDate, COALESCE(no_letter_experience, 0) as no_letter_experience FROM work_experience_letters WHERE YEAR(date_letter_exprerience) = ? AND MONTH(date_letter_exprerience) = ? ORDER BY date_letter_exprerience DESC", yearstring, helper.StringMonth()).
		Scan(&CountExperienceByDate, &CountNoExperienceEmployee)
	if err != nil {
		fmt.Print(err.Error())
	}

	var CountCertificate, CountExperience int
	CountCertificate = repository.CountResign("certificate_of_employments", fmt.Sprintf("number_of_employees = '%s' ", number_of_employees))
	CountExperience = repository.CountResign("work_experience_letters", fmt.Sprintf("number_of_employees = '%s' ", number_of_employees))

	if CountCertificate < 1 && resign.Periode_of_service >= 365 && resign.Type == "true" {
		var data = map[string]interface{}{
			"resign_id":                 Resign_id,
			"number_of_employees":       number_of_employees,
			"date_certificate_employee": helper.DMY(),
			"no_certificate_employee":   (CountNoCertificateEmployee + 1),
			"rom":                       helper.Rom(helper.StringMonth()),
			"created_at":                helper.DMYhms(),
			"updated_at":                helper.DMYhms(),
		}
		repository.InsertResign("certificate_of_employments", data)
	} else if CountExperience == 0 {
		var data = map[string]interface{}{
			"resign_id":               Resign_id,
			"number_of_employees":     number_of_employees,
			"date_letter_exprerience": helper.DMY(),
			"no_letter_experience":    (CountNoExperienceEmployee + 1),
			"rom":                     helper.Rom(helper.StringMonth()),
			"created_at":              helper.DMYhms(),
			"updated_at":              helper.DMYhms(),
		}
		repository.InsertResign("work_experience_letters", data)
	}
}
