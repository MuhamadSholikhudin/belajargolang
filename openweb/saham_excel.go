package main

import (
	"golang.design/x/clipboard"
	"fmt"
	"log"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os/exec"
	"time"
)

func main() {
	// Calling Sleep method
	time.Sleep(3 * time.Second)
	xlsx, err := excelize.OpenFile("saham_excel.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	// Benar
	sheet1Name := "saham_excel"
	var mulai int
	for index, _ := range xlsx.GetRows(sheet1Name) {
		mulai = (index + 1)

		// Copy text
		a := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", mulai))
		ClipboardText(fmt.Sprint(a))

		//tiny 
		b := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", mulai))
		OpenSahamClipboard(fmt.Sprint(b))
	}
	fmt.Println("ACTION SUCCESS !")
}

func ClipboardText(text string) {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(fmt.Sprint(text)))
}

func OpenSahamClipboard(app string) {
	cmd := exec.Command(app)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}