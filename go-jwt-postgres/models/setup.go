package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB, db *gorm.DB
var err error

func ConnectDatabase() {
	// db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_jwt_mux"))
	dsn := "host=localhost user=postgres password=hwi1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}
	db.AutoMigrate(&User{})
	DB = db
}
