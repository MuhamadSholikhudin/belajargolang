package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hwi")
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

	file, err := os.Open("kuesioners.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, record := range records {

		var Count_Resignation_submission, Resignation_submission_id int
		var Number_of_employees string
		err = db.QueryRow("SELECT COUNT(id), COALESCE(id, 0), COALESCE(number_of_employees, 'KOSONG') FROM  resignation_submissions WHERE number_of_employees = ? ", record[2]).
			Scan(&Count_Resignation_submission, &Resignation_submission_id, &Number_of_employees)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if Count_Resignation_submission != 0 {
			_, err = db.Exec("UPDATE kuesioners SET resignation_submission_id = ? WHERE number_of_employees = ?", Resignation_submission_id, record[2])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(Resignation_submission_id)
		}

	}
	fmt.Println("succes !!")
}
