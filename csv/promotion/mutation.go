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

	csvFile, err := os.Open("mutation.csv")
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
			old_department, _ := strconv.Atoi(line4)
			new_department, _ := strconv.Atoi(line5)
			_, err := db.Exec("INSERT INTO promotionmutations (number_of_employees, name, start_date, action, old_department, new_department, remark, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", line0, line1, line2, line3, old_department, new_department, line6, datetimenow, datetimenow)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			mu.Unlock()
			wg.Done()
		}(&wg, line[0], line[1], line[2], line[3], line[4], line[5], line[6])
		wg.Wait()
	}
	fmt.Println("Successfully MIGRATION CSV file")
}
