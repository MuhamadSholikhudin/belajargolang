package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	dsn := "root:@tcp(localhost:3306)/hrd"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a backup file
	backupFile, err := ioutil.TempFile("", "mysql_backup_")
	if err != nil {
		log.Fatal(err)
	}
	defer backupFile.Close()

	// Perform the backup
	err = db.Exec("mysqldump -u root -p hrd > " + backupFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	// Print the backup file name
	fmt.Println("Backup file:", backupFile.Name())
}
