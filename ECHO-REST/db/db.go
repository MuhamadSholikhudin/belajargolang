package db

import (
	"belajargolang/ECHO-REST/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString1 := conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	// connectionString := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/echo_rest")
	fmt.Println(connectionString1)
	db, err = sql.Open("mysql", connectionString1)
	if err != nil {
		panic("Connestion String error")
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

}

func CreateConn() *sql.DB {
	return db
}
