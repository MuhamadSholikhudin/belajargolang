package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

func Vonnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrd")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	db, err := Vonnect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	xlsx, err := excelize.OpenFile("./files/bpjs.xlsx")

	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	// Benar
	sheet1Name := "Sheet1"
	// row := make([]M, 0)

	var mulai int
	for index, _ := range xlsx.GetRows(sheet1Name) {

		mulai = index + 1

		a := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", mulai))
		b := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", mulai))
		c := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", mulai))
		d := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", mulai))
		e := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", mulai))

		_, err = db.Exec("update employees set bpjs_ketenagakerjaan = ?, date_bpjs_ketenagakerjaan = ?, bpjs_kesehatan = ?, date_bpjs_kesehatan = ?  where number_of_employees = ?", b, c, d, e, a)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(a, b, c, d, e)

	}
	fmt.Println("UPDATE SUCCESS !")

}
