package main

import (
	"database/sql"
	"fmt"

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

func main() {
	sqlQuery()
	// sqlQueryRow()
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrdit")
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

	xlsx.SetCellValue(sheet1Name, "A1", "id")
	xlsx.SetCellValue(sheet1Name, "B1", "name")
	xlsx.SetCellValue(sheet1Name, "C1", "number_of_employees")
	// xlsx.SetCellValue(sheet1Name, "D1", "address_jalan")
	// xlsx.SetCellValue(sheet1Name, "E1", "hire_date")
	// xlsx.SetCellValue(sheet1Name, "D1", " hire_date")
	xlsx.SetCellValue(sheet1Name, "D1", "job_id")
	xlsx.SetCellValue(sheet1Name, "E1", "department_id")
	xlsx.SetCellValue(sheet1Name, "F1", "finger_id")
	xlsx.SetCellValue(sheet1Name, "G1", "gender")
	xlsx.SetCellValue(sheet1Name, "H1", "place_of_birth")
	xlsx.SetCellValue(sheet1Name, "I1", "date_of_birth")
	xlsx.SetCellValue(sheet1Name, "J1", "marital_status")
	xlsx.SetCellValue(sheet1Name, "K1", "religion")
	xlsx.SetCellValue(sheet1Name, "L1", "biological_mothers_name")
	xlsx.SetCellValue(sheet1Name, "M1", "national_id")
	xlsx.SetCellValue(sheet1Name, "N1", "address_jalan")
	xlsx.SetCellValue(sheet1Name, "O1", "address_rt ")
	xlsx.SetCellValue(sheet1Name, "P1", "address_rw")
	xlsx.SetCellValue(sheet1Name, "Q1", "address_village")
	xlsx.SetCellValue(sheet1Name, "R1", "address_district")
	xlsx.SetCellValue(sheet1Name, "S1", "address_city ")
	xlsx.SetCellValue(sheet1Name, "T1", "address_province")
	xlsx.SetCellValue(sheet1Name, "U1", "phone ")
	xlsx.SetCellValue(sheet1Name, "V1", "email ")
	xlsx.SetCellValue(sheet1Name, "W1", "educate")
	xlsx.SetCellValue(sheet1Name, "X1", "major ")
	xlsx.SetCellValue(sheet1Name, "Y1", "hire_date")
	xlsx.SetCellValue(sheet1Name, "Z1", " employee_type ")
	xlsx.SetCellValue(sheet1Name, "AA1", "end_of_contract")
	xlsx.SetCellValue(sheet1Name, "AB1", "date_out")
	xlsx.SetCellValue(sheet1Name, "AC1", "exit_statement ")
	xlsx.SetCellValue(sheet1Name, "AD1", "bank_name ")
	xlsx.SetCellValue(sheet1Name, "AE1", "bank_branch ")
	xlsx.SetCellValue(sheet1Name, "AF1", "bank_account_name ")
	xlsx.SetCellValue(sheet1Name, "AG1", "bank_account_number ")
	xlsx.SetCellValue(sheet1Name, "AH1", "bpjs_ketenagakerjaan")
	xlsx.SetCellValue(sheet1Name, "AI1", "date_bpjs_ketenagakerjaan")
	xlsx.SetCellValue(sheet1Name, "AJ1", "bpjs_kesehatan")
	xlsx.SetCellValue(sheet1Name, "AK1", "date_bpjs_kesehatan")
	xlsx.SetCellValue(sheet1Name, "AL1", "npwp")
	xlsx.SetCellValue(sheet1Name, "AM1", "kode_ptkp")
	xlsx.SetCellValue(sheet1Name, "AN1", "year_ptkp")
	xlsx.SetCellValue(sheet1Name, "AO1", "bagian")
	xlsx.SetCellValue(sheet1Name, "A1P", "cell")
	xlsx.SetCellValue(sheet1Name, "A1Q", "status_employee")
	xlsx.SetCellValue(sheet1Name, "AJ1", "bpjs_kesehatan")
	xlsx.SetCellValue(sheet1Name, "AK1", "date_bpjs_kesehatan")
	xlsx.SetCellValue(sheet1Name, "AL1", "npwp")
	xlsx.SetCellValue(sheet1Name, "AM1", "kode_ptkp")
	xlsx.SetCellValue(sheet1Name, "AN1", "year_ptkp")
	xlsx.SetCellValue(sheet1Name, "AO1", "bagian")
	xlsx.SetCellValue(sheet1Name, "A1P", "cell")
	xlsx.SetCellValue(sheet1Name, "A1Q", "status_employee")

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 0
	rows, err := db.Query("select id, name, number_of_employees, job_id, department_id , COALESCE(finger_id, '') as finger_id,  COALESCE(gender, '') as gender, COALESCE(place_of_birth, '') as place_of_birth, COALESCE(date_of_birth, '') as date_of_birth, COALESCE(marital_status, '') as marital_status,  COALESCE(religion, '') as religion, COALESCE(biological_mothers_name, '') as biological_mothers_name, COALESCE(national_id, '') as national_id, COALESCE(address_jalan, '') as address_jalan, COALESCE(address_rt, '') as address_rt, COALESCE(address_rw, '') as address_rw, COALESCE(address_village, '') as address_village, COALESCE(address_district, '') as address_district, COALESCE(address_city, '') as address_city, COALESCE(address_province, '') as address_province, COALESCE(phone, '') as phone, COALESCE(email, '') as email, COALESCE(educate, '') as educate, COALESCE(major, '') as major, COALESCE(hire_date, '') as hire_date, COALESCE(employee_type, '') as employee_type, COALESCE(end_of_contract, '') as end_of_contract, COALESCE(date_out, '') as date_out, COALESCE(exit_statement, '') as exit_statement, COALESCE(bank_name, '') as bank_name, COALESCE(bank_branch, '') as bank_branch, COALESCE(bank_account_name, '') as bank_account_name, COALESCE(bank_account_number, '') as bank_account_number, COALESCE(bpjs_ketenagakerjaan, '') as bpjs_ketenagakerjaan, COALESCE(date_bpjs_ketenagakerjaan, '') as date_bpjs_ketenagakerjaan, COALESCE(bpjs_kesehatan, '') as bpjs_kesehatan, COALESCE(date_bpjs_kesehatan, '') as date_bpjs_kesehatan, COALESCE(npwp, '') as npwp, COALESCE(kode_ptkp, '') as kode_ptkp, COALESCE(year_ptkp, '') as year_ptkp, COALESCE(bagian, '') as bagian, COALESCE(cell, '') as cell, COALESCE(status_employee, '') as status_employee from employees  where id > ?", age)
	// rows, err := db.Query("select id, name, number_of_employees, job_id, department_id, COALESCE(finger_id, '') as finger_id, COALESCE(gender, '') as gender, COALESCE(place_of_birth, '') as place_of_birth, COALESCE(date_of_birth, '') as date_of_birth, COALESCE(marital_status, '') as marital_status, COALESCE(biological_mothers_name, '') as biological_mothers_name, COALESCE(national_id, '') as national_id, COALESCE(address_jalan, '') as address_jalan, COALESCE(address_rt, '') as address_rt, COALESCE(address_rw, '') as address_rw, COALESCE(address_village, '') as address_village, COALESCE(address_district, '') as address_district, COALESCE(address_city, '') as address_city, COALESCE(address_province, '') as address_province, COALESCE(phone, '') as phone, COALESCE(email, '') as email, COALESCE(educate, '') as educate, COALESCE(major, '') as major, COALESCE(hire_date, '') as hire_date, COALESCE(employee_type, '') as employee_type, COALESCE(end_of_contract, '') as end_of_contract, COALESCE(date_out, '') as date_out, COALESCE(exit_statement, '') as exit_statement, COALESCE(bank_name, '') as bank_name, COALESCE(bank_branch, '') as bank_branch, COALESCE(bank_account_name, '') as bank_account_name, COALESCE(bank_account_number, '') as bank_account_number, COALESCE(bpjs_ketenagakerjaan, '') as bpjs_ketenagakerjaan, COALESCE(date_bpjs_ketenagakerjaan, '') as date_bpjs_ketenagakerjaan, COALESCE(bpjs_kesehatan, '') as bpjs_kesehatan, COALESCE(date_bpjs_kesehatan, '') as date_bpjs_kesehatan, COALESCE(npwp, '') as npwp, COALESCE(kode_ptkp, '') as kode_ptkp, COALESCE(year_ptkp, '') as year_ptkp, COALESCE(bagian, '') as bagian, COALESCE(cell, '') as cell, COALESCE(status_employee, '') as status_employee, COALESCE(bpjs_kesehatan, '') as bpjs_kesehatan, COALESCE(date_bpjs_kesehatan, '') as date_bpjs_kesehatan, COALESCE(npwp, '') as npwp , COALESCE(kode_ptkp, '') as kode_ptkp , COALESCE(year_ptkp, '') as year_ptkp , COALESCE(bagian, '') as bagian , COALESCE(cell, '') as cell , COALESCE(status_employee, '') as status_employee from employees  where id > ?", age)
	// rows, err := db.Query("select id, name, number_of_employees, job_id, department_id, finger_id, gender, place_of_birth, date_of_birth, marital_status, religion, biological_mothers_name, national_id, address_jalan, address_rt, address_rw, address_village, address_district, address_city, address_province, phone, email, educate, major, hire_date, employee_type, end_of_contract, date_out, exit_statement , bank_name, bank_branch, bank_account_name, bank_account_number, bpjs_ketenagakerjaan, date_bpjs_ketenagakerjaan, bpjs_kesehatan, date_bpjs_kesehatan, npwp , kode_ptkp, year_ptkp, bagian, cell, status_employee from employees  where id > ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.number_of_employees, &each.job_id, &each.department_id, &each.finger_id, &each.gender, &each.place_of_birth, &each.date_of_birth, &each.marital_status, &each.religion, &each.biological_mothers_name, &each.national_id, &each.address_jalan, &each.address_rt, &each.address_rw, &each.address_district, &each.address_district, &each.address_city, &each.address_province, &each.phone, &each.email, &each.educate, &each.major, &each.hire_date, &each.employee_type, &each.end_of_contract, &each.date_out, &each.exit_statement, &each.bank_name, &each.bank_branch, &each.bank_account_name, &each.bank_account_number, &each.bpjs_ketenagakerjaan, &each.date_bpjs_ketenagakerjaan, &each.bpjs_kesehatan, &each.date_bpjs_kesehatan, &each.npwp, &each.kode_ptkp, &each.year_ptkp, &each.bagian, &each.cell, &each.status_employee)
		// var err = rows.Scan(&each.id, &each.name, &each.number_of_employees, &each.job_id, &each.department_id, &each.finger_id, &each.gender, &each.place_of_birth, &each.date_of_birth, &each.marital_status, &each.biological_mothers_name, &each.national_id, &each.address_jalan, &each.address_rt, &each.address_rw, &each.address_district, &each.address_city, &each.address_province, &each.phone, &each.email, &each.educate, &each.major, &each.hire_date, &each.employee_type, &each.end_of_contract, &each.date_out, &each.exit_statement, &each.bank_name, &each.bank_branch, &each.bank_account_name, &each.bank_account_number, &each.bpjs_ketenagakerjaan, &each.date_bpjs_ketenagakerjaan, &each.bpjs_kesehatan, &each.date_bpjs_kesehatan, &each.npwp, &each.kode_ptkp, &each.year_ptkp, &each.bagian, &each.cell, &each.status_employee)
		// var err = rows.Scan(&each.id, &each.name, &each.number_of_employees, &each.job_id, &each.department_id, &each.finger_id)

		// &each.id, &each.name, &each.number_of_employees)

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

	err = xlsx.SaveAs("./file1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

/*

func sqlQueryRow() {
    var db, err = connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    var result = student{}
    var id = "E001"
    err = db.
        QueryRow("select name, grade from tb_student where id = ?", id).
        Scan(&result.name, &result.grade)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Printf("name: %s\ngrade: %d\n", result.name, result.grade)
}
*/
