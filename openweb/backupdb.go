package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Replace with your database connection details
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/salary")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Open a file to write backup data
	file, err := os.Create("backup.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Perform a SELECT query to fetch data
	rows, err := db.Query("SELECT * FROM salary")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through rows and write data to the file
	for rows.Next() {
		var data string
		err := rows.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}
		_, err = file.WriteString(fmt.Sprintf("%s\n", data)) // Write data to the file
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Backup completed successfully!")
}