package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Employee struct {
	Id                  int    `json:"id"`
	Number_of_employees string `json:"number_of_employees"`
	Name                string `json:"name"`
	National_id         string `json:"national_id"`
}

func Conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrdit")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Resigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var db, err = Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	u, err := url.Parse(r.RequestURI)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	var sqlPaging string = "SELECT id, number_of_employees, name, COALESCE(national_id, '') FROM employees"
	var sqlCount string = "SELECT COUNT(*) FROM employees"

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

	var totalminbyperpage int64
	totalminbyperpage = total - ((total / 10) * 10)

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

	var employee []Employee

	for rows.Next() {
		var each = Employee{}
		var err = rows.Scan(&each.Id, &each.Number_of_employees, &each.Name, &each.National_id)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		employee = append(employee, each)
	}

	links := map[string]interface{}{
		"first": fmt.Sprintf("http://localhost:8989/resigns/?page=%s%s", first, params),
		"last":  fmt.Sprintf("http://localhost:8989/resigns/?page=%s%s", last, params),
		"next":  fmt.Sprintf("http://localhost:8989/resigns/?page=%s%s", next, params),
		"prev":  fmt.Sprintf("http://localhost:8989/resigns/?page=%s%s", prev, params),
	}

	informationpages := map[string]interface{}{
		"currentPage": page,
		"from":        ((page - 1) * 10) + 1,
		"lastPage":    lastPage,
		"perPage":     10,
		"to":          ((page - 1) * 10) + len(employee),
		"total":       total,
	}

	pages := map[string]interface{}{
		"page": informationpages,
	}

	result := map[string]interface{}{
		"meta":  pages,
		"data":  employee,
		"links": links,
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte(resp))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/resigns/", Resigns).Methods("GET")

	fmt.Println("Listen on port localhost:8989")
	http.ListenAndServe(":8989", r)
}
