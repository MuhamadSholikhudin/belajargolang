// Golang program to illustrate the usage of
// strings.Replace() function

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

/*
func main() {

	var chicken = map[string]string{
		"januari":  ",,,,,,,,,,,,13,14,,,,,,,,,,,,,,,,",
		"februari": ",,,,,,,,9,,,,,,,,,,,,,,,,,,,,,",
		"maret":    ",,,,,,,,,,,,,,,,,,,,,,,,,,27,28,,",
		"april":    ",,,,,6,,,9,,,,,,,,,,,20,21,,,,,,27,,,",
	}

	var no int = 1
	for key, val := range chicken {

		fmt.Println(key, "  \t:", val)

		//Initialization data string
		val_string := ",,,,,,,,,,,,13,14,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,27,28,,"

		ganti_1 := strings.Replace(val_string, ",,", ",", -1)
		ganti_2 := strings.Replace(ganti_1, ",,", ",", -1)
		ganti_3 := strings.Replace(ganti_2, ",,", ",", -1)
		ganti_4 := strings.Replace(ganti_3, ",,", ",", -1)
		ganti_5 := strings.Replace(ganti_4, ",,", ",", -1)
		ganti_6 := strings.Replace(ganti_5, ",,", ",", -1)
		ganti_7 := strings.Replace(ganti_6, ",,", ",", -1)

		// Jadikan data menjadi array
		data_array := strings.Split(ganti_7, ",")

		// Looping data array cuti
		for _, fruit := range data_array {
			if fruit != "" {
				no += 1
				// SET DATA EXCEL
				fmt.Printf("bulan %s %d : %s\n", key, no, fruit)
			}
		}

	}

	// ganti ,, dengan ,
	// fmt.Println(",,,,,,,,,,,,13,14,,,,,,,,,,,,,,,,")

}

*/
func main() {

	readxlsx, err := excelize.OpenFile("./bank/files/convert.xlsx")

	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	writexlsx := excelize.NewFile()

	sheet1Name1 := "SEMESTER 1 2021"
	writexlsx.SetSheetName(writexlsx.GetSheetName(1), sheet1Name1)

	writexlsx.SetCellValue(sheet1Name1, "A1", "NIK")
	writexlsx.SetCellValue(sheet1Name1, "B1", "NAMA")
	writexlsx.SetCellValue(sheet1Name1, "C1", "TANGGAL")

	// Benar
	sheet1Name := "Sheet1"
	var no int = 0
	for index, _ := range readxlsx.GetRows(sheet1Name) {

		CEL_A := readxlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", (index+1)))
		CEL_B := readxlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", (index+1)))
		CEL_C := readxlsx.GetCellValue(sheet1Name, fmt.Sprintf("c%d", (index+1)))

		if CEL_B != "" {
			val_string := CEL_C
			ganti_1 := strings.Replace(val_string, ",,", ",", -1)
			ganti_2 := strings.Replace(ganti_1, ",,", ",", -1)
			ganti_3 := strings.Replace(ganti_2, ",,", ",", -1)
			ganti_4 := strings.Replace(ganti_3, ",,", ",", -1)
			ganti_5 := strings.Replace(ganti_4, ",,", ",", -1)
			ganti_6 := strings.Replace(ganti_5, ",,", ",", -1)
			ganti_7 := strings.Replace(ganti_6, ",,", ",", -1)

			// Jadikan data menjadi array
			data_array := strings.Split(ganti_7, ",")

			// Looping data array cuti
			for _, tgl := range data_array {
				if tgl != "" {
					no += 1
					// SET DATA EXCEL
					writexlsx.SetCellValue(sheet1Name1, fmt.Sprintf("A%d", no), CEL_A)
					writexlsx.SetCellValue(sheet1Name1, fmt.Sprintf("B%d", no), CEL_B)

					tanggal := fmt.Sprintf("2021-08-%s", tgl)
					if len(tgl) == 1 {
						tanggal = fmt.Sprintf("2021-08-0%s", tgl)
					}

					dateStart_date, error := time.Parse("2006-01-02", tanggal)
					if error != nil {
						fmt.Println(error)
						return
					}
					writexlsx.SetCellValue(sheet1Name1, fmt.Sprintf("B%d", no), dateStart_date)
				}
			}
		}

	}

	err = writexlsx.SaveAs("./file123.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	// var no int = 1
	// for key, val := range chicken {

	// 	fmt.Println(key, "  \t:", val)

	// 	//Initialization data string
	// 	val_string := ",,,,,,,,,,,,13,14,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,27,28,,"

	// 	ganti_1 := strings.Replace(val_string, ",,", ",", -1)
	// 	ganti_2 := strings.Replace(ganti_1, ",,", ",", -1)
	// 	ganti_3 := strings.Replace(ganti_2, ",,", ",", -1)
	// 	ganti_4 := strings.Replace(ganti_3, ",,", ",", -1)
	// 	ganti_5 := strings.Replace(ganti_4, ",,", ",", -1)
	// 	ganti_6 := strings.Replace(ganti_5, ",,", ",", -1)
	// 	ganti_7 := strings.Replace(ganti_6, ",,", ",", -1)

	// 	// Jadikan data menjadi array
	// 	data_array := strings.Split(ganti_7, ",")

	// 	// Looping data array cuti
	// 	for _, fruit := range data_array {
	// 		if fruit != "" {
	// 			no += 1
	// 			// SET DATA EXCEL
	// 			fmt.Printf("bulan %s %d : %s\n", key, no, fruit)
	// 		}
	// 	}

	// }
}
