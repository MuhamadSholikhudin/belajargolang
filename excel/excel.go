package main

import (
	"database/sql"
	"fmt"
	"log"

	// "github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

type M map[string]interface{}

var data = []M{
	M{"Name": "Noval", "Gender": "male", "Age": 18},
	M{"Name": "Nabila", "Gender": "female", "Age": 12},
	M{"Name": "Yasa", "Gender": "male", "Age": 11},
}

// type dept struct {
// 	number_of_employees string
// 	name                string
// }

func main() {
	// magic here

	// var res []dept
	// res = sqlQuery()
	// exportExcel(res)

	// excelWrite()
	ReadRows()

}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrdit")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

// func sqlQuery() []dept {

// 	db, err := connect()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// selalu defer db.close() alias ditutup, defer akan mengeksekusi
// 	// db.close() diakhir setelah semua proses selesai
// 	// simpan di awal ok, mau didefer di akhir juga bagus
// 	defer db.Close()

// 	rows, err := db.Query("select * from employees")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// selalu defer close koneksi setelah dipakai
// 	defer rows.Close()

// 	var result []dept

// 	for rows.Next() {
// 		var each = dept{}
// 		var err = rows.Scan(&each.number_of_employees, &each.name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		result = append(result, each)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return result
// }

// func exportExcel(result []dept) {

// 	xlsx := excelize.NewFile()
// 	sheet1Name := "holy sheet"

// 	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

// 	err := xlsx.AutoFilter(sheet1Name, "A1", "B1", "")

// 	for i, each := range result {
// 		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+1), each.number_of_employees)
// 		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+1), each.name)
// 	}

// 	err = xlsx.SaveAs("./file2.xlsx")

// 	if err != nil {

// 		fmt.Println(err)

// 	}

// }

// func exportExcel(result []dept) {

//     xlsx := excelize.NewFile()
//     sheet1Name := "holy sheet"

//     xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

//     err := xlsx.AutoFilter(sheet1Name, "A1", "B1", "")

//     for i, each := range result {

//         xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+1), each.number_of_employees)

//         xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+1), each.name)

//     }
//     err = xlsx.SaveAs("./file2.xlsx")
//     if err != nil {
//         fmt.Println(err)
//     }

// }

func createExcel() {

	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "Name")
	xlsx.SetCellValue(sheet1Name, "B1", "Gender")
	xlsx.SetCellValue(sheet1Name, "C1", "Age")

	err := xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	for i, each := range data {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each["Name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each["Gender"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each["Age"])
	}

	err = xlsx.SaveAs("./file1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

}

// func excelWrite() {
// 	xlsx, err := excelize.OpenFile("./file1.xlsx")
// 	if err != nil {
// 		log.Fatal("ERROR", err.Error())
// 	}

// 	sheet1Name := "Sheet One"

// 	rows := make([]M, 0)
// 	for i := 2; i < 5; i++ {
// 		row := M{
// 			"Name":   xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
// 			"Gender": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
// 			"Age":    xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
// 		}
// 		rows = append(rows, row)
// 	}

// 	fmt.Printf("%v \n", rows)
// }

func ReadRows() {
	f, err := excelize.OpenFile("sample.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.Rows("Sheet One")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
		}
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	if err = rows.Close(); err != nil {
		fmt.Println(err)
	}
}
