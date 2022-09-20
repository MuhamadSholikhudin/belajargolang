package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

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
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrdit")
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

	for index, _ := range xlsx.GetRows(sheet1Name) {
		fmt.Println(index)

		tambah := index + 1

		a := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", tambah))
		b := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", tambah))
		c := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", tambah))
		d := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", tambah))
		e := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", tambah))
		f := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", tambah))
		g := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", tambah))
		h := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("H%d", tambah))
		i := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", tambah))
		j := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("J%d", tambah))
		k := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", tambah))
		l := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("L%d", tambah))
		m := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("M%d", tambah))
		// n := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("N%d", tambah))

		fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m)

	}

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
