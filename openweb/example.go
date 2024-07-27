// Golang program to get the current
// date and time in various format
package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	age "github.com/theTardigrade/golang-age"
)

type Employee struct {
	id                  int
	name                string
	number_of_employees string
}

func ConHrd() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrd")
	if err != nil {
		return nil, err
	}
	return db, nil
}
func ConSalry() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/salary")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	UpdatePeriodicAllowence()
}

func UpdatePeriodicAllowence() {
	// Initilization databae  hrd
	dhrd, err := ConHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dhrd.Close()

	// Initilization databae salary
	dbsalary, err := ConSalry()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbsalary.Close()

	// using time.Now() function
	// to get the current time
	currentTime := time.Now()

	//Get Current Yaer
	CurrentYear, _ := strconv.Atoi(currentTime.Format("2006"))

	//Looping Current Year at start 2016
	for tahun := 2016; tahun <= CurrentYear; tahun++ {

		// Create a date to combine with the current year on the system
		tanggal := fmt.Sprintf("%d-%s", tahun, currentTime.Format("01-02"))
		dateStart_date, error := time.Parse("2006-01-02", tanggal)
		if error != nil {
			fmt.Println(error)
			return
		}

		// Calculate date based on hire date
		agehiredate := age.CalculateToNow(dateStart_date)

		// Query Data from  employees where hiredate
		sql_by_date := fmt.Sprintf("SELECT number_of_employees, job_id FROM `employees` WHERE hire_date = '%s' ", tanggal)
		//fmt.Println(sql_by_date)

		//Get data from employees where hiredate
		rows, err := dhrd.Query(sql_by_date)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		// Loping data employees
		for rows.Next() {
			var number_of_employees string
			var job_id int

			// convert data coloumn from database to type go
			var err = rows.Scan(&number_of_employees, &job_id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var periodic_allowance int
			periodic_allowance = 5000 * agehiredate
			// Selection job this is operator get 3000 * agehiredate or non-operator
			if job_id == 24 || job_id == 23 || job_id == 25 {
				periodic_allowance = 3000 * agehiredate
			}

			// Execute salary data update base on query
			_, err = dbsalary.Exec("update salary set periodic_allowance = ? where number_of_employees = ?", periodic_allowance, number_of_employees)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}
