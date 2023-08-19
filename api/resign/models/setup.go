package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnHrd() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrd")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnResign() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/resign")
	if err != nil {
		return nil, err
	}
	return db, nil
}
