package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
	"github.com/360EntSecGroup-Skylar/excelize"

)

type M map[string]interface{}

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	http.HandleFunc("/viewexcel", routeIndexGetexcel)

	http.HandleFunc("/excelprocess", routeSubmitPostExcel)
	http.HandleFunc("/importexcelprocess", routeSubmitImportExcel)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func someURL() string {
	url := url.URL{
		Scheme: "https",
		Host:   "example.com",
	}
	return url.String()
}

func routeIndexGetexcel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("excelview.html"))
	var err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// hapus()
}

func routeSubmitImportExcel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// alias := r.FormValue("alias")

	uploadedFile, _, err := r.FormFile("file")

	// xlsx, err := excelize.OpenFile("./files/file1.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	xlsx, err := excelize.OpenReader(uploadedFile)
	// if err != nil {
	// 	return err
	// }
	// for index, name := range f.GetSheetMap() {
	// 	fmt.Println(index, name)
	// }

	// Benar
	sheet1Name := "Sheet One"
	// row := make([]M, 0)
	for index, _ := range xlsx.GetRows(sheet1Name) {
		fmt.Println(index)

		tambah := index + 1

		// row := M{
		a := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", tambah))
		b := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", tambah))
		c := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", tambah))
		d := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", tambah))
		e := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", tambah))
		f := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", tambah))
		g := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", tambah))
		h := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("H%d", tambah))
		i := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", tambah))
		j := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("J%d", tambah))
		k := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", tambah))
		l := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("L%d", tambah))
		m := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("M%d", tambah))
		n := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("N%d", tambah))
		o := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("O%d", tambah))
		p := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("P%d", tambah))
		q := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Q%d", tambah))
		r := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("R%d", tambah))
		s := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("S%d", tambah))
		t := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("T%d", tambah))
		u := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("U%d", tambah))
		v := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("V%d", tambah))
		w := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("W%d", tambah))
		x := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("X%d", tambah))
		y := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Y%d", tambah))
		z := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Z%d", tambah))
		aa := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AA%d", tambah))
		ab := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AB%d", tambah))
		ac := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AC%d", tambah))
		ad := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AD%d", tambah))
		ae := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AE%d", tambah))
		af := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AF%d", tambah))
		ag := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AG%d", tambah))
		ah := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AH%d", tambah))
		ai := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AI%d", tambah))
		aj := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AJ%d", tambah))
		ak := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AK%d", tambah))
		al := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AL%d", tambah))
		am := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AM%d", tambah))
		an := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AN%d", tambah))
		ao := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AO%d", tambah))
		ap := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AP%d", tambah))
		aq := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AQ%d", tambah))
		// }

		// for _, colCell := range row {
		fmt.Print(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z, aa, ab, ac, ad, ae, af, ag, ah, ai, aj, ak, al, am, an, ao, ap, aq)
		// }
	}
	// redirectURL := resp.Header.Get(someURL())
	// fmt.Println(someURL())
}

	// func routeSubmitPostExcel(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != "POST" {
	// 		http.Error(w, "", http.StatusBadRequest)
	// 		return
	// 	}

	// 	if err := r.ParseMultipartForm(1024); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}


	alias := r.FormValue("alias")

	uploadedFile, handler, err := r.FormFile("file")


	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	defer uploadedFile.Close()


	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	excelWrite()

	w.Write([]byte("done"))
}


// ASLI
func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("view.html"))
	var err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	alias := r.FormValue("alias")

	uploadedFile, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))
}

const (
	YYYYMMDD = "2006-01-02"
)

func excelWrite() {
	/*

			var (
				file multipart.File
				size int64
				err  error
			)



		file, _, err = req.FormFile("key")
		// size = // Calculate size
		xlsx.OpenReaderAt(file, size)

	*/

	xlsx, err := excelize.OpenFile("./files/file1.xlsx")

	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	sheet1Name := "Sheet One"

	rows := make([]M, 0)

	now := time.Now()
	after := now.AddDate(0, 0, 0)

	// buat tanggal default excel 01/00/1900
	var timedefaultexcel = time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)

	// Tangkap data dari excel
	date_of_birth_excel := 44771

	// Konfersi data
	//date_of_birth_convert := 1
	after = timedefaultexcel.AddDate(0, 0, date_of_birth_excel)
	
	fmt.Println(after.Format(YYYYMMDD))

	for i := 2; i < 1005; i++ {
		row := M{
			"a":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
			"b":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
			"c":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
			"d":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)),
			"e":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i)),
			"f":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", i)),
			"g":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", i)),
			"h":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("H%d", i)),
			"i":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", i)),
			"j":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("J%d", i)),
			"k":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", i)),
			"l":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("L%d", i)),
			"m":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("M%d", i)),
			"n":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("N%d", i)),
			"o":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("O%d", i)),
			"p":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("P%d", i)),
			"q":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Q%d", i)),
			"r":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("R%d", i)),
			"s":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("S%d", i)),
			"t":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("T%d", i)),
			"u":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("U%d", i)),
			"v":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("V%d", i)),
			"w":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("W%d", i)),
			"x":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("X%d", i)),
			"y":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Y%d", i)),
			"z":  xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Z%d", i)),
			"aa": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AA%d", i)),
			"ab": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AB%d", i)),
			"ac": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AC%d", i)),
			"ad": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AD%d", i)),
			"ae": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AE%d", i)),
			"af": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AF%d", i)),
			"ag": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AG%d", i)),
			"ah": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AH%d", i)),
			"ai": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AI%d", i)),
			"aj": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AJ%d", i)),
			"ak": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AK%d", i)),
			"al": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AL%d", i)),
			"am": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AM%d", i)),
			"an": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AN%d", i)),
			"ao": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AO%d", i)),
			"ap": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AP%d", i)),
			"aq": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("AQ%d", i)),
		}
		rows = append(rows, row)
	}

	for _, val := range rows {
		fmt.Println(val)
	}

	// for _, row := range rows {
	// 	// fmt.Printf("%s\n", row)

	// }

	// for _, row := range rows {
	// 	fmt.Printf("nama buah : %s\n", row)
	// }

	// fmt.Printf("%v \n", rows)
}

func hapus() {
	// Removing file from the directory
	// Using Remove() function
	e := os.Remove("./files/file1.xlsx")
	if e != nil {
		log.Fatal(e)
	}
}
