package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var mu sync.Mutex

var wg sync.WaitGroup

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

	rows, err := db.Query("SELECT employee_id FROM violations ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var employee_id int
		var err = rows.Scan(&employee_id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, employee_id int) {
			mu.Lock()
			var number_of_employees, name string
			err = db.QueryRow("SELECT number_of_employees, name FROM employees WHERE id = ? ", employee_id).
				Scan(&number_of_employees, &name)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			_, err := db.Exec("UPDATE violations SET number_of_employees = ?, name = ? WHERE employee_id = ?", number_of_employees, name, employee_id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			mu.Unlock()
			wg.Done()
		}(&wg, employee_id)
		wg.Wait()
	}
}
