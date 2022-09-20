package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

/*
const (
	YYYYMMDD = "2006-01-02"
)
*/

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrd")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var startworks = []map[string]string{

		map[string]string{"employee_id": "2", "job_id": "19", "startwork_date": "2016-04-01"},
		map[string]string{"employee_id": "5", "job_id": "23", "startwork_date": "2016-04-01"},
		map[string]string{"employee_id": "6", "job_id": "19", "startwork_date": "2016-04-01"},
	}
	for _, startwork := range startworks {
		fmt.Println(startwork["employee_id"], startwork["job_id"], startwork["startwork_date"])

		intemployee_id, _ := strconv.Atoi(startwork["employee_id"])
		intjob_id, _ := strconv.Atoi(startwork["job_id"])

		_, err = db.Exec("update startworks set job_id = ?, startwork_date = ? where employee_id = ?", intjob_id, startwork["startwork_date"], intemployee_id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("update success!")

	}

}
