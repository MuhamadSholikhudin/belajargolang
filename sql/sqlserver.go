package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	var (
		server   = "10.10.100.246"
		port     = 1033
		user     = "sa"
		password = "Pastibisa123"
		database = "attendance_machine_database_3"
	)

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error creating connection pool:", err.Error())
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err.Error())
		return
	}

	fmt.Println("Successfully connected to SQL Server")
}
