package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	bacaexcel()
}

func bacaexcel() {
	xlsx, err := excelize.OpenFile("PERMOHONAN.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	sheet1Name := "Input"
	// for i := 2; i < 5; i++ {
	// 	fmt.Println(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)))
	// 	fmt.Println(reflect.TypeOf(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i))))
	// 	fmt.Println(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i)))
	// 	typ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i))
	// 	fmt.Println(reflect.TypeOf(typ))
	// 	fmt.Println(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)))
	// }

	// fmt.Println(len(xlsx.GetRows(sheet1Name)))

	rows, err := xlsx.Rows(sheet1Name)
	var mulai int = 1
	for rows.Next() {
		// for i, _ := range rows.Columns() {
		// 	fmt.Print(i)
		// }
		fmt.Println(mulai + 1)
	}
	// var mulai int
	// for i, _ := range xlsx.GetRows(sheet1Name) {
	// 	mulai = i + 1
	// 	fmt.Println(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", mulai)), mulai)
	// 	fmt.Println(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", mulai)), reflect.TypeOf(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", mulai))))
	// 	typ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", mulai))
	// 	fmt.Println(typ, reflect.TypeOf(typ))
	// 	fmt.Println(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", mulai)))
	// }
}
