package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	employee_id int
	activity    string
	cell        string
	bagian      string
	remark      string
}

type Promotions struct {
	old_job_level        int
	new_job_level        int
	start_date_job_level string
}

type Mutations struct {
	old_department        int
	new_department        int
	start_date_department string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrd")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func routeSubmitImportExcel(w http.ResponseWriter, r *http.Request) {
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
	xlsx, err := excelize.OpenReader(uploadedFile)
	sheet1Name := "PromotionMutation"

	for index := range xlsx.GetRows(sheet1Name) {
		tambah := index + 1
		// employee_id
		//nocell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", tambah))

		// number_of_employees
		number_of_employeescell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", tambah))

		// name
		namecell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", tambah))

		// old_job_level
		old_job_levelcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", tambah))

		// new_Job_level
		new_job_levelcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", tambah))

		// start_date_job_level
		start_date_job_levelcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", tambah))

		// old_department
		old_departmentcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", tambah))

		// new_department
		new_departmentcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("H%d", tambah))

		//start_date_department
		start_date_departmentcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", tambah))

		//bagian
		bagiancell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("J%d", tambah))

		// cell
		cellcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", tambah))

		//remark
		remarkcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("L%d", tambah))

		var tve string
		var vale, err = DateExcelToGo(start_date_job_levelcell)
		tve = TimeToStringYYYYMMDD(vale)
		if err != nil {
			tve = "string"
		}
		// fmt.Println(a, b, c, d, tve, f, g, h, i, j, k, l, m)
		fmt.Println(number_of_employeescell, namecell, old_job_levelcell, new_job_levelcell, tve, old_departmentcell, new_departmentcell, start_date_departmentcell, bagiancell, cellcell, remarkcell)
	}
}

// func CheckEmp(number_of_empoyees string) int {

// }

func TimeToStringYYYYMMDD(t time.Time) string {
	location, _ := time.LoadLocation("Asia/Bangkok")
	t.In(location).Format("2006-01-02")
	var tts string
	tts = t.In(location).Format("2006-01-02")
	return tts
}

func DateExcelToGo(DateString string) (time.Time, error) {
	var err error
	_, err = strconv.Atoi(DateString)
	if err != nil {
		return time.Now(), err
	}
	intUnix, err := strconv.ParseInt(DateString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	var timego int64
	timego = ((intUnix - 25569) * 86400)
	myTime := time.Unix(timego, 0)
	return myTime, nil
}

// ASLI
func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("view.html"))
	var err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/", routeIndexGet)

	http.HandleFunc("/importexcelprocess", routeSubmitImportExcel)

	fmt.Println("server started at localhost:1001")
	http.ListenAndServe(":1001", nil)
}
