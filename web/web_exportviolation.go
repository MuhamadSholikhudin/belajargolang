package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id                        int
	name                      string
	number_of_employees       string
	job_id                    int
	department_id             int
	finger_id                 string
	gender                    string
	place_of_birth            string
	date_of_birth             string
	marital_status            string
	religion                  string
	biological_mothers_name   string
	national_id               string
	address_jalan             string
	address_rt                string
	address_rw                string
	address_village           string
	address_district          string
	address_city              string
	address_province          string
	phone                     string
	email                     string
	educate                   string
	major                     string
	hire_date                 string
	employee_type             string
	end_of_contract           string
	date_out                  string
	exit_statement            string
	bank_name                 string
	bank_branch               string
	bank_account_name         string
	bank_account_number       string
	bpjs_ketenagakerjaan      string
	date_bpjs_ketenagakerjaan string
	bpjs_kesehatan            string
	date_bpjs_kesehatan       string
	npwp                      string
	kode_ptkp                 string
	year_ptkp                 string
	bagian                    string
	cell                      string
	status_employee           string
}

type job struct {
	id        int
	job_level string
}
type department struct {
	id         int
	department string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrd")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {

	// Membuat excel
	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellValue(sheet1Name, "A1", "NIK")
	xlsx.SetCellValue(sheet1Name, "B1", "NAMA KARYAWAN")
	xlsx.SetCellValue(sheet1Name, "C1", "JABATAN")
	xlsx.SetCellValue(sheet1Name, "D1", "DEPARTEMENT")
	xlsx.SetCellValue(sheet1Name, "E1", "TGL MASUK")
	xlsx.SetCellValue(sheet1Name, "F1", "NO SP")
	xlsx.SetCellValue(sheet1Name, "G1", "NO SP")
	xlsx.SetCellValue(sheet1Name, "H1", "FORMAT AT")
	xlsx.SetCellValue(sheet1Name, "I1", "BULAN SP")
	xlsx.SetCellValue(sheet1Name, "J1", "ROM")
	xlsx.SetCellValue(sheet1Name, "K1", "HARI LAP (angka)")
	xlsx.SetCellValue(sheet1Name, "L1", "HARI LAPORAN")
	xlsx.SetCellValue(sheet1Name, "M1", "TGL LAPORAN")
	xlsx.SetCellValue(sheet1Name, "N1", "TAHUN")
	xlsx.SetCellValue(sheet1Name, "O1", "KATA PENGANTAR")
	xlsx.SetCellValue(sheet1Name, "P1", "PASAL YANG DI LANGGAR")
	xlsx.SetCellValue(sheet1Name, "Q1", "BUNYI PASAL PELANGGARAN SEKARANG JIKA SUDAH PERNAH")
	xlsx.SetCellValue(sheet1Name, "R1", "BUNYI PASAL")
	xlsx.SetCellValue(sheet1Name, "S1", "PASAL 25 AYAT 2A, 3A, 4A , B, 5A,B DAN C JIKA PERNAH DAPAT SP (PELANGGARAN SEKARANG)")
	xlsx.SetCellValue(sheet1Name, "T1", "BUNYI PASAL1")
	xlsx.SetCellValue(sheet1Name, "U1", "KETERANGAN LAIN")
	xlsx.SetCellValue(sheet1Name, "V1", "KETENGAN LAIN 2")
	xlsx.SetCellValue(sheet1Name, "W1", "KETERANGAN LAIN 1")
	xlsx.SetCellValue(sheet1Name, "X1", "PELANGGARAN SEBELUMNYA")
	xlsx.SetCellValue(sheet1Name, "Y1", "TGL SP")
	xlsx.SetCellValue(sheet1Name, "Z1", " KETERANGAN")
	xlsx.SetCellValue(sheet1Name, "AA1", "REKAP SESUAI DENGAN LAPORAN PELANGGARAN")
	xlsx.SetCellValue(sheet1Name, "AB1", "Tambahan Keterangan")
	xlsx.SetCellValue(sheet1Name, "AC1", "AN HRD")
	xlsx.SetCellValue(sheet1Name, "AD1", "Resign")
	xlsx.SetCellValue(sheet1Name, "AE1", "TANGGAL PENYAMPAIAN SP")
	xlsx.SetCellValue(sheet1Name, "AF1", "CEKLIST")
	xlsx.SetCellValue(sheet1Name, "AG1", "Status Violation")

	// xlsx.SetCellValue(sheet1Name, "AH1", "bpjs_ketenagakerjaan")
	// xlsx.SetCellValue(sheet1Name, "AI1", "date_bpjs_ketenagakerjaan")
	// xlsx.SetCellValue(sheet1Name, "AJ1", "bpjs_kesehatan")
	// xlsx.SetCellValue(sheet1Name, "AK1", "date_bpjs_kesehatan")
	// xlsx.SetCellValue(sheet1Name, "AL1", "npwp")
	// xlsx.SetCellValue(sheet1Name, "AM1", "kode_ptkp")
	// xlsx.SetCellValue(sheet1Name, "AN1", "year_ptkp")
	// xlsx.SetCellValue(sheet1Name, "AO1", "bagian")
	// xlsx.SetCellValue(sheet1Name, "A1P", "cell")
	// xlsx.SetCellValue(sheet1Name, "A1Q", "status_employee")
	// xlsx.SetCellValue(sheet1Name, "AJ1", "bpjs_kesehatan")
	// xlsx.SetCellValue(sheet1Name, "AK1", "date_bpjs_kesehatan")
	// xlsx.SetCellValue(sheet1Name, "AL1", "npwp")
	// xlsx.SetCellValue(sheet1Name, "AM1", "kode_ptkp")
	// xlsx.SetCellValue(sheet1Name, "AN1", "year_ptkp")
	// xlsx.SetCellValue(sheet1Name, "AO1", "bagian")
	// xlsx.SetCellValue(sheet1Name, "A1P", "cell")
	// xlsx.SetCellValue(sheet1Name, "A1Q", "status_employee")

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 0
	rows, err := db.Query("select id, name, number_of_employees, job_id, department_id , COALESCE(finger_id, '') as finger_id,  COALESCE(gender, '') as gender, COALESCE(place_of_birth, '') as place_of_birth, COALESCE(date_of_birth, '') as date_of_birth, COALESCE(marital_status, '') as marital_status,  COALESCE(religion, '') as religion, COALESCE(biological_mothers_name, '') as biological_mothers_name, COALESCE(national_id, '') as national_id, COALESCE(address_jalan, '') as address_jalan, COALESCE(address_rt, '') as address_rt, COALESCE(address_rw, '') as address_rw, COALESCE(address_village, '') as address_village, COALESCE(address_district, '') as address_district, COALESCE(address_city, '') as address_city, COALESCE(address_province, '') as address_province, COALESCE(phone, '') as phone, COALESCE(email, '') as email, COALESCE(educate, '') as educate, COALESCE(major, '') as major, COALESCE(hire_date, '') as hire_date, COALESCE(employee_type, '') as employee_type, COALESCE(end_of_contract, '') as end_of_contract, COALESCE(date_out, '') as date_out, COALESCE(exit_statement, '') as exit_statement, COALESCE(bank_name, '') as bank_name, COALESCE(bank_branch, '') as bank_branch, COALESCE(bank_account_name, '') as bank_account_name, COALESCE(bank_account_number, '') as bank_account_number, COALESCE(bpjs_ketenagakerjaan, '') as bpjs_ketenagakerjaan, COALESCE(date_bpjs_ketenagakerjaan, '') as date_bpjs_ketenagakerjaan, COALESCE(bpjs_kesehatan, '') as bpjs_kesehatan, COALESCE(date_bpjs_kesehatan, '') as date_bpjs_kesehatan, COALESCE(npwp, '') as npwp, COALESCE(kode_ptkp, '') as kode_ptkp, COALESCE(year_ptkp, '') as year_ptkp, COALESCE(bagian, '') as bagian, COALESCE(cell, '') as cell, COALESCE(status_employee, '') as status_employee from employees  where id > ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.number_of_employees, &each.job_id, &each.department_id, &each.finger_id, &each.gender, &each.place_of_birth, &each.date_of_birth, &each.marital_status, &each.religion, &each.biological_mothers_name, &each.national_id, &each.address_jalan, &each.address_rt, &each.address_rw, &each.address_district, &each.address_district, &each.address_city, &each.address_province, &each.phone, &each.email, &each.educate, &each.major, &each.hire_date, &each.employee_type, &each.end_of_contract, &each.date_out, &each.exit_statement, &each.bank_name, &each.bank_branch, &each.bank_account_name, &each.bank_account_number, &each.bpjs_ketenagakerjaan, &each.date_bpjs_ketenagakerjaan, &each.bpjs_kesehatan, &each.date_bpjs_kesehatan, &each.npwp, &each.kode_ptkp, &each.year_ptkp, &each.bagian, &each.cell, &each.status_employee)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	var job = job{}
	var department = department{}

	for _, each := range result {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", each.id), each.id)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", each.id), each.name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", each.id), each.number_of_employees)

		err = db.
			QueryRow("select id, job_level  from jobs where id = ?", each.job_id).
			Scan(&job.id, &job.job_level)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", each.id), job.job_level)

		err = db.
			QueryRow("select id, department  from departments where id = ?", each.department_id).
			Scan(&department.id, &department.department)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", each.id), department.department)

		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", each.id), each.finger_id)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", each.id), each.gender)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", each.id), each.place_of_birth)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("I%d", each.id), each.date_of_birth)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", each.id), each.marital_status)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", each.id), each.religion)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("L%d", each.id), each.biological_mothers_name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("M%d", each.id), each.national_id)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("N%d", each.id), each.address_jalan)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("O%d", each.id), each.address_rt)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("P%d", each.id), each.address_rw)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("Q%d", each.id), each.address_village)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("R%d", each.id), each.address_district)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("S%d", each.id), each.address_city)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("T%d", each.id), each.address_province)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("U%d", each.id), each.phone)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("V%d", each.id), each.email)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("W%d", each.id), each.educate)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("X%d", each.id), each.major)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("Y%d", each.id), each.hire_date)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("Z%d", each.id), each.employee_type)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AA%d", each.id), each.end_of_contract)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AB%d", each.id), each.date_out)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AC%d", each.id), each.exit_statement)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AD%d", each.id), each.bank_name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AE%d", each.id), each.bank_branch)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AF%d", each.id), each.bank_account_name)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AG%d", each.id), each.bank_account_number)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AH%d", each.id), each.bpjs_ketenagakerjaan)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AI%d", each.id), each.date_bpjs_ketenagakerjaan)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AJ%d", each.id), each.bpjs_kesehatan)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AK%d", each.id), each.date_bpjs_kesehatan)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AL%d", each.id), each.npwp)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AM%d", each.id), each.kode_ptkp)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AN%d", each.id), each.year_ptkp)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AO%d", each.id), each.bagian)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AP%d", each.id), each.cell)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AQ%d", each.id), each.status_employee)
	}

	err = xlsx.SaveAs("./employees.xlsx")
	if err != nil {
		fmt.Println(err)
	}

}

func exportviolations(w http.ResponseWriter, r *http.Request) {
	sqlQuery()
	fmt.Fprintln(w, "Download Sukses File Master Pelanggaran ")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo! Apa kabar saya")

	})

	http.HandleFunc("/exportviolations", exportviolations)

	fmt.Println("starting web server at http://localhost:2000/")
	http.ListenAndServe(":2000", nil)
}
