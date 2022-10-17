package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

func Bonnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrdit")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	db, err := Bonnect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	xlsx, err := excelize.OpenFile("./files/upadtedateend.xlsx")

	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	// Benar
	sheet1Name := "Sheet1"

	var mulai int
	for index, _ := range xlsx.GetRows(sheet1Name) {

		mulai = index + 1

		a := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", mulai))
		b := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", mulai))

		violation_id, _ := strconv.Atoi(a)

		_, err = db.Exec("update violations set date_end_violation = ? where id = ? ", b, violation_id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("UPDATE SUCCESS !")

}
