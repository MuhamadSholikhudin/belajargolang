package main

import (
	"fmt"
	"io/ioutil"
	"log"

	//"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/xuri/excelize/v2"
)

func main() {

	Setexl, err := excelize.OpenFile("renamefile.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for i, f := range files {
		baris := i + 1
		// Set value of cell D4 to 88
		err = Setexl.SetCellValue("listfile", fmt.Sprintf("A%d", baris), f.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Save the changes to the file.
	err = Setexl.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("list data sudah di copy")
}
