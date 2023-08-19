package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mu sync.Mutex

var wg sync.WaitGroup

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/promotionmutation")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	csvFile, err := os.Open("promotion.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	t := time.Now()
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println(err)
	}
	datetimenow := t.In(location).Format("2006-01-02 15:04:05")
	for _, line := range csvLines {
		wg.Add(1)
		go func(wg *sync.WaitGroup, line0 string, line1 string, line2 string, line3 string, line4 string, line5 string, line6 string) {
			mu.Lock()
			old_job, _ := strconv.Atoi(line2)
			new_job, _ := strconv.Atoi(line3)
			_, err := db.Exec("INSERT INTO promotionmutations (number_of_employees, name, old_job_level, new_job_level, start_date, action, remark, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", line0, line1, old_job, new_job, line4, line6, line5, datetimenow, datetimenow)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			mu.Unlock()
			wg.Done()
		}(&wg, line[0], line[1], line[2], line[3], line[4], line[5], line[6])
		wg.Wait()
	}
	fmt.Println("Successfully Opened CSV file")
}
