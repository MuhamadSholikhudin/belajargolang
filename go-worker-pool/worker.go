package main

import (
	"database/sql"
	"fmt"
	"runtime"
	"strconv"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Add(x int) {
	c.val++
}

func (c *counter) Value() (x int) {
	return c.val
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrd")
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Employee struct {
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
	grade_employee            string
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

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var mtx sync.Mutex
	var meter counter

	for i := 0; i < 2; i++ {
		wg.Add(1)

		db, err := connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		xlsx := excelize.NewFile()
		sheet1Name := "Sheet1"

		xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

		err = xlsx.AutoFilter(sheet1Name, "A1", "AX1", "")
		if err != nil {
			fmt.Println(err.Error())
		}

		xlsx.SetCellValue(sheet1Name, "A1", "Line_Type")
		xlsx.SetCellValue(sheet1Name, "B1", "finger_id")
		xlsx.SetCellValue(sheet1Name, "C1", "number_of_employees")
		xlsx.SetCellValue(sheet1Name, "D1", "name")
		xlsx.SetCellValue(sheet1Name, "E1", "gender")
		xlsx.SetCellValue(sheet1Name, "F1", "place_of_birth")
		xlsx.SetCellValue(sheet1Name, "G1", "date_of_birth")
		xlsx.SetCellValue(sheet1Name, "H1", "marital_status")
		xlsx.SetCellValue(sheet1Name, "I1", "religion")
		xlsx.SetCellValue(sheet1Name, "J1", "biological_mothers_name")
		xlsx.SetCellValue(sheet1Name, "K1", "national_id")
		xlsx.SetCellValue(sheet1Name, "L1", "address_jalan")
		xlsx.SetCellValue(sheet1Name, "M1", "address_rt")
		xlsx.SetCellValue(sheet1Name, "N1", "address_rw")
		xlsx.SetCellValue(sheet1Name, "O1", "address_village")
		xlsx.SetCellValue(sheet1Name, "P1", "address_district")
		xlsx.SetCellValue(sheet1Name, "Q1", "address_city")
		xlsx.SetCellValue(sheet1Name, "R1", "address_province")
		xlsx.SetCellValue(sheet1Name, "S1", "phone")
		xlsx.SetCellValue(sheet1Name, "T1", "email")
		xlsx.SetCellValue(sheet1Name, "U1", "hire_date")
		xlsx.SetCellValue(sheet1Name, "V1", "employee_type")
		xlsx.SetCellValue(sheet1Name, "W1", "end_of_contract")
		xlsx.SetCellValue(sheet1Name, "X1", "date_out")
		xlsx.SetCellValue(sheet1Name, "Y1", "exit_statement")
		xlsx.SetCellValue(sheet1Name, "Z1", "job_level")
		xlsx.SetCellValue(sheet1Name, "AA1", "grade_employee")
		xlsx.SetCellValue(sheet1Name, "AB1", "department")
		xlsx.SetCellValue(sheet1Name, "AC1", "bagian")
		xlsx.SetCellValue(sheet1Name, "AD1", "cell")
		xlsx.SetCellValue(sheet1Name, "AE1", "kode_ptkp")
		xlsx.SetCellValue(sheet1Name, "AF1", "year_ptkp")
		xlsx.SetCellValue(sheet1Name, "AG1", "educate")
		xlsx.SetCellValue(sheet1Name, "AH1", "major")
		xlsx.SetCellValue(sheet1Name, "AI1", "status_employee")
		xlsx.SetCellValue(sheet1Name, "AJ1", "bank_name")
		xlsx.SetCellValue(sheet1Name, "AK1", "bank_branch")
		xlsx.SetCellValue(sheet1Name, "AL1", "bank_account_name")
		xlsx.SetCellValue(sheet1Name, "AM1", "bank_account_number")
		xlsx.SetCellValue(sheet1Name, "AN1", "bpjs_ketenagakerjaan")
		xlsx.SetCellValue(sheet1Name, "AO1", "date_bpjs_ketenagakerjaan")
		xlsx.SetCellValue(sheet1Name, "AP1", "bpjs_kesehatan")
		xlsx.SetCellValue(sheet1Name, "AQ1", "date_bpjs_kesehatan")
		xlsx.SetCellValue(sheet1Name, "AR1", "npwp")

		rows, err := db.Query("select name, number_of_employees, job_id, department_id , COALESCE(grade_employee, '') as grade_employee, COALESCE(finger_id, '') as finger_id,  COALESCE(gender, '') as gender, COALESCE(place_of_birth, '') as place_of_birth, COALESCE(date_of_birth, '') as date_of_birth, COALESCE(marital_status, '') as marital_status,  COALESCE(religion, '') as religion, COALESCE(biological_mothers_name, '') as biological_mothers_name, COALESCE(national_id, '') as national_id, COALESCE(address_jalan, '') as address_jalan, COALESCE(address_rt, '') as address_rt, COALESCE(address_rw, '') as address_rw, COALESCE(address_village, '') as address_village, COALESCE(address_district, '') as address_district, COALESCE(address_city, '') as address_city, COALESCE(address_province, '') as address_province, COALESCE(phone, '') as phone, COALESCE(email, '') as email, COALESCE(educate, '') as educate, COALESCE(major, '') as major, COALESCE(hire_date, '') as hire_date, COALESCE(employee_type, '') as employee_type, COALESCE(end_of_contract, '') as end_of_contract, COALESCE(date_out, '') as date_out, COALESCE(exit_statement, '') as exit_statement, COALESCE(bank_name, '') as bank_name, COALESCE(bank_branch, '') as bank_branch, COALESCE(bank_account_name, '') as bank_account_name, COALESCE(bank_account_number, '') as bank_account_number, COALESCE(bpjs_ketenagakerjaan, '') as bpjs_ketenagakerjaan, COALESCE(date_bpjs_ketenagakerjaan, '') as date_bpjs_ketenagakerjaan, COALESCE(bpjs_kesehatan, '') as bpjs_kesehatan, COALESCE(date_bpjs_kesehatan, '') as date_bpjs_kesehatan, COALESCE(npwp, '') as npwp, COALESCE(kode_ptkp, '') as kode_ptkp, COALESCE(year_ptkp, '') as year_ptkp, COALESCE(bagian, '') as bagian, COALESCE(cell, '') as cell, COALESCE(status_employee, '') as status_employee from employees  ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		go func() {

			for j := 0; j < 1; j++ {
				mtx.Lock()

				no := 1
				for rows.Next() {
					var each Employee
					var err = rows.Scan(&each.name, &each.number_of_employees, &each.job_id, &each.department_id, &each.grade_employee, &each.finger_id, &each.gender, &each.place_of_birth, &each.date_of_birth, &each.marital_status, &each.religion, &each.biological_mothers_name, &each.national_id, &each.address_jalan, &each.address_rt, &each.address_rw, &each.address_village, &each.address_district, &each.address_city, &each.address_province, &each.phone, &each.email, &each.educate, &each.major, &each.hire_date, &each.employee_type, &each.end_of_contract, &each.date_out, &each.exit_statement, &each.bank_name, &each.bank_branch, &each.bank_account_name, &each.bank_account_number, &each.bpjs_ketenagakerjaan, &each.date_bpjs_ketenagakerjaan, &each.bpjs_kesehatan, &each.date_bpjs_kesehatan, &each.npwp, &each.kode_ptkp, &each.year_ptkp, &each.bagian, &each.cell, &each.status_employee)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					no += 1
					if (each.status_employee) == "active" {
						xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), "ACTIVE")
					} else {
						xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", no), "NON ACTIVE")
					}

					finger_id, _ := strconv.ParseInt(each.number_of_employees[5:10], 10, 64)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", no), finger_id)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", no), each.number_of_employees)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", no), each.name)

					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", no), each.gender)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", no), each.place_of_birth)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", no), each.date_of_birth)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("H%d", no), each.marital_status)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("I%d", no), each.religion)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", no), each.biological_mothers_name)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", no), each.national_id)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("L%d", no), each.address_jalan)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("M%d", no), each.address_rt)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("N%d", no), each.address_rw)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("O%d", no), each.address_village)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("P%d", no), each.address_district)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("Q%d", no), each.address_city)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("R%d", no), each.address_province)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("S%d", no), each.phone)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("T%d", no), each.email)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("U%d", no), each.hire_date)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("V%d", no), each.employee_type)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("W%d", no), each.end_of_contract)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("X%d", no), each.date_out)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("Y%d", no), each.exit_statement)

					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("Z%d", no), JobDepartment(each.number_of_employees)[0])

					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AA%d", no), each.grade_employee)

					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AB%d", no), JobDepartment(each.number_of_employees)[1])

					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AC%d", no), each.bagian)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AD%d", no), each.cell)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AE%d", no), each.kode_ptkp)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AF%d", no), each.year_ptkp)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AG%d", no), each.educate)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AH%d", no), each.major)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AI%d", no), each.status_employee)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AJ%d", no), each.bank_name)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AK%d", no), each.bank_branch)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AL%d", no), each.bank_account_name)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AM%d", no), each.bank_account_number)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AN%d", no), each.bpjs_ketenagakerjaan)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AO%d", no), each.date_bpjs_ketenagakerjaan)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AP%d", no), each.bpjs_kesehatan)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AQ%d", no), each.date_bpjs_kesehatan)
					xlsx.SetCellValue(sheet1Name, fmt.Sprintf("AR%d", no), each.npwp)

				}
				err = xlsx.SaveAs("./file2.xlsx")
				if err != nil {
					fmt.Println(err)
				}

				meter.Add(1)

				mtx.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
}

func JobDepartment(number_of_employees string) []string {

	var db, err = connect()
	defer db.Close()

	Employee := Employee{}
	err = db.QueryRow("SELECT job_id, department_id FROM employees WHERE number_of_employees = ?", number_of_employees).
		Scan(&Employee.job_id, &Employee.department_id)
	if err != nil {
		fmt.Println(err.Error())
	}
	var Job_level, Department string
	err = db.QueryRow("SELECT job_level FROM jobs WHERE id = ?", Employee.job_id).Scan(&Job_level)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.QueryRow("SELECT department FROM departments WHERE id = ?", Employee.department_id).Scan(&Department)
	if err != nil {
		fmt.Println(err.Error())
	}
	var output = []string{Job_level, Department}
	return output
}
