package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hwi1234"
	dbname   = "hi"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")
	// age := 21
	rows, err := db.Query("SELECT reported FROM list_cases")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var reported string
		var err = rows.Scan(&reported)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("reported", reported)
	}

	var id int

	_ = db.QueryRow("SELECT id FROM list_cases WHERE LIMIT 1 ").
		Scan(&id)

	fmt.Println("id", id)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
