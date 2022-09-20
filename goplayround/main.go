package main

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exportdata(rw http.ResponseWriter, request *http.Request) {

	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "a")

	var b bytes.Buffer
	writr := bufio.NewWriter(&b)
	xlsx.Write(writr)
	writr.Flush()

	fileContents := b.Bytes()
	fileSize := strconv.Itoa(len(fileContents))

	rw.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	rw.Header().Set("Content-disposition", "attachment;filename=export.xlsx")
	rw.Header().Set("Content-Length", fileSize)

	t := bytes.NewReader(b.Bytes())
	io.Copy(rw, t)
}

func main() {

	http.HandleFunc("/", exportdata)
	http.ListenAndServe(":8080", nil)

}
