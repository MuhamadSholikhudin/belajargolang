package models

import (
	"belajargolang/ECHO-REST/db"
	"belajargolang/ECHO-REST/helpers"
	"database/sql"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateConn()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).
		Scan(&obj.Id, &obj.Username, &pwd)
	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found")
		return false, err
	}
	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}
	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash password don't match")
		return false, err
	}
	return true, nil
}
