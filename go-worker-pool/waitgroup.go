package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConnString2   = "root:@tcp(127.0.0.1:3306)/hrd"
	dbMaxIdleConns2 = 4
	dbMaxConns2     = 100
)

func worker(GlobalRank string, ind int) {
	fmt.Printf("Worker %d starting  Worker %s done\n", ind, GlobalRank)
}

func openDbConnection2() (*sql.DB, error) {
	log.Println("=> open db connection")

	db, err := sql.Open("mysql", dbConnString2)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns2)
	db.SetMaxIdleConns(dbMaxIdleConns2)

	return db, nil
}

func main() {

	start := time.Now()
	// var wg sync.WaitGroup

	db2, err := openDbConnection2()
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db2.Query("SELECT number_of_employees  FROM employees")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var ind int = 0
	for rows.Next() {
		var number_of_employees string
		var err = rows.Scan(&number_of_employees)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// wg.Add(1)
		ind = ind + 1
		// go func() {
		// defer wg.Done()
		worker(number_of_employees, ind)
		// }()

		// fmt.Printf("Worker %d starting  Worker %d done\n", ind, number_of_employees)
	}

	// wg.Wait()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")

}
