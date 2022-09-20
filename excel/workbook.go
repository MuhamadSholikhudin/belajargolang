package main

import (
	"fmt"
	"net/http"

	"github.com/xuri/excelize/v2"
)

func process(w http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	f, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	f.NewSheet("NewSheet")
	w.Header().Set("Content-Disposition", "attachment; filename=Book1.xlsx")
	w.Header().Set("Content-Type", req.Header.Get("Content-Type"))
	if _, err := f.WriteTo(w); err != nil {
		fmt.Fprintf(w, err.Error())
	}
	return
}

func main() {
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8090", nil)
}
