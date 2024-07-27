// Go program to illustrate how to rename
// and move a file in default directory
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	// "github.com/xuri/excelize/v2"
)

func main() {

	fmt.Print("Rename start")
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
	// Calling Sleep method
	xlsx, err := excelize.OpenFile("renamefile.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	var gagal int = 0
	// Benar
	sheet1Name := "rename"
	var mulai int
	for index, _ := range xlsx.GetRows(sheet1Name) {
		mulai = (index + 1)

		//Get data from cell A
		a := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", mulai))

		//Get data from cell B
		b := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", mulai))

		Original_Path := a
		New_Path := b
		e := os.Rename(Original_Path, New_Path)
		if e != nil {
			// fmt.Println("file tidak dapat di rename -> ", a)
			gagal = gagal + 1
			xlsx.SetCellValue("gagal", fmt.Sprintf("A%d", gagal), fmt.Sprintf("%s", a))
		} else {
			// fmt.Println(b, "Sukses")
		}
	}
	// Save changes to the Excel file
	err = xlsx.SaveAs("renamefile.xlsx")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	fmt.Println("Rename Sukses")

	time.Sleep(3 * time.Second)

}
