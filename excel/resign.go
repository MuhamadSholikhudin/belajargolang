package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

type M map[string]interface{}

type ResignSubmission struct {
	Number_of_employees          string `json:"number_of_employees"`
	Name                         string `json:"name"`
	Position                     string `json:"position"`
	Department                   string `json:"department"`
	Building                     string `json:"building"`
	Hire_date                    string `json:"hire_date"`
	Date_out                     string `json:"date_out"`
	Date_resignation_submissions string `json:"date_resignation_submissions"`
	Type                         string `json:"type"`
	Reason                       string `json:"reason"`
	Detail_reason                string `json:"detail_reason"`
	Suggestion                   string `json:"suggestion"`
	Periode_of_service           int    `json:"periode_of_service"`
	Status_resignsubmisssion     string `json:"status_resignsubmisssion"`
	Age                          int    `json:"age"`
	Using_media                  string `json:"using_media"`
	Classification               string `json:"classification"`
	Created_at                   string `json:"created_at"`
	Updated_at                   string `json:"updated_at"`
}

var data = []M{
	M{"Name": "Noval", "Gender": "male", "Age": 18},
	M{"Name": "Nabila", "Gender": "female", "Age": 12},
	M{"Name": "Yasa", "Gender": "male", "Age": 11},
}

func Conn() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hwi")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func excel() {

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

}

func doPrint(wg *sync.WaitGroup, message string, no int) {
	defer wg.Done()

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
	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), message)
	err = xlsx.SaveAs("./file1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ExportSubmission()
	// for rows.Next() {
	// 	var Number_of_employees, National_id string
	// 	var Id int
	// 	var err = rows.Scan(&Number_of_employees, &National_id, &Id)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}

	// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", Id), messages)
	// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", Id), National_id)
	// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", Id), Id)
	// }

	/*
		var datas = []map[string]string{
			map[string]string{"name": "chicken blue", "gender": "male"},
			map[string]string{"name": "chicken red", "gender": "male"},
			map[string]string{"name": "chicken yellow", "gender": "female"},
		}
	*/

	/*
		for _, each := range datas {
			go func(who string) {
				messages <- who
			}(each["name"])
		}

		for i, each := range datas {
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each["name"])
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each["gender"])
		}

	*/

}

func ExportSubmission() {
	db, err := Conn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select number_of_employees from resignation_submissions where id > ?", 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	xlsx := excelize.NewFile()
	sheet1Name := "holy sheet"

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	err = xlsx.AutoFilter(sheet1Name, "A1", "Q1", "")
	if err != nil {
		fmt.Println(err.Error())
	}
	xlsx.SetCellValue(sheet1Name, "A1", "NIK")
	xlsx.SetCellValue(sheet1Name, "B1", "NAME")
	xlsx.SetCellValue(sheet1Name, "C1", "POSISI")
	xlsx.SetCellValue(sheet1Name, "D1", "DEPARTMENT")
	xlsx.SetCellValue(sheet1Name, "E1", "GEDUNG")
	xlsx.SetCellValue(sheet1Name, "F1", "HIRE DATE")
	xlsx.SetCellValue(sheet1Name, "G1", "DATE OUT")
	xlsx.SetCellValue(sheet1Name, "H1", "TYPE")
	xlsx.SetCellValue(sheet1Name, "I1", "ALASAM")
	xlsx.SetCellValue(sheet1Name, "J1", "ALASAN TAMBAHAN")
	xlsx.SetCellValue(sheet1Name, "K1", "UMUR")
	xlsx.SetCellValue(sheet1Name, "L1", "SARAN")
	xlsx.SetCellValue(sheet1Name, "M1", "STATUS RESIGN")
	xlsx.SetCellValue(sheet1Name, "N1", "USING MEDIA")
	xlsx.SetCellValue(sheet1Name, "O1", "CLASSIFIKASI")
	xlsx.SetCellValue(sheet1Name, "P1", "CREATEAD AT")
	xlsx.SetCellValue(sheet1Name, "Q1", "UPDATED AT")
	xlsx.SetCellValue(sheet1Name, "R1", "TGL PERMOHONAN")

	var wg sync.WaitGroup

	no := 1

	for rows.Next() {
		var NIK string
		var err = rows.Scan(&NIK)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		no += 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, message string, no int) {
			defer wg.Done()
			fmt.Println(message, " NIK ", no)
			var Submission ResignSubmission

			err = db.QueryRow("SELECT COALESCE(number_of_employees, ''),	COALESCE(name, ''),	COALESCE(position, ''),	COALESCE(department, ''),	COALESCE(building, ''),	COALESCE(hire_date, '0000-00-00'),	COALESCE(date_out, '0000-00-00'),	COALESCE(date_resignation_submissions, '0000-00-00'),	COALESCE(type, ''),	COALESCE(reason, ''),	COALESCE(detail_reason, ''),	COALESCE(periode_of_service, ''),	COALESCE(age, 0),	COALESCE(suggestion, ''),	COALESCE(status_resignsubmisssion, ''),	COALESCE(using_media, ''),	COALESCE(classification, ''),	COALESCE(created_at, '0000-00-00 00:00:00'),	COALESCE(updated_at, '0000-00-00 00:00:00')	from resignation_submissions where number_of_employees = ?", message).
				Scan(&Submission.Number_of_employees, &Submission.Name, &Submission.Position, &Submission.Department, &Submission.Building, &Submission.Hire_date, &Submission.Date_out, &Submission.Date_resignation_submissions, &Submission.Type, &Submission.Reason, &Submission.Detail_reason, &Submission.Periode_of_service, &Submission.Age, &Submission.Suggestion, &Submission.Status_resignsubmisssion, &Submission.Using_media, &Submission.Classification, &Submission.Created_at, &Submission.Updated_at)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), Submission.Number_of_employees)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", no), Submission.Name)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", no), Submission.Position)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", no), Submission.Department)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", no), Submission.Building)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", no), Submission.Hire_date)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", no), Submission.Date_out)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", no), Submission.Type)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("I%d", no), Submission.Reason)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", no), Submission.Detail_reason)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", no), Submission.Age)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("L%d", no), Submission.Suggestion)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("M%d", no), Submission.Status_resignsubmisssion)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("N%d", no), Submission.Using_media)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("O%d", no), Submission.Classification)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("P%d", no), Submission.Created_at)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("Q%d", no), Submission.Updated_at)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("R%d", no), Submission.Date_resignation_submissions)

		}(&wg, NIK, no)

	}

	wg.Wait()

	err = xlsx.SaveAs("./Pengajuan_resign.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
