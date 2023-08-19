package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	valid "github.com/asaskevich/govalidator"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
)

type Employees struct {
	id            int
	job_id        int
	department_id int
}

// type Promotions struct {
// 	old_job_level        int
// 	new_job_level        int
// 	start_date_job_level string
// }

type Jobs struct {
	id        int
	job_level string
	level     int
}

type Departments struct {
	id         int
	department string
}

type Count struct {
	countid int
}

const (
	YYYYMMDD          = "2006-01-02"
	YYYYMMDDHHMMSS24h = "2006-01-02 15:04:05"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrd")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Convertdateyyyymmdd(excel string) string {

	var dateexcel string
	dateexcel = excel

	var outpoutconvert string

	var cekint bool
	cekint = valid.IsInt(dateexcel) //cek startdate dia int atau string

	if cekint == true { // jika dateexcel bertipe int

		//var dateexcel int // buat variabel  dengan tipe data int
		var start_date_job_levelconvint int

		start_date_job_levelconvint, e := strconv.Atoi(dateexcel) // parse string ke int
		if e == nil {
			fmt.Printf("%T \n %v", start_date_job_levelconvint, start_date_job_levelconvint)
		}

		var i_dateexcel int

		i_dateexcel = (start_date_job_levelconvint - 25569) * 86400 //perhitungan excel(int) - 25569 * 86400(selisih hari pada golang)

		stringstart_date_job_level := strconv.FormatInt(int64(i_dateexcel), 10) //parse int ke string

		intstart_date_job_level, err := strconv.ParseInt(stringstart_date_job_level, 10, 64) //parse string ke time

		if err != nil {
			log.Fatal(err)
		}

		mystart_date_job_level := time.Unix(intstart_date_job_level, 0) // buat unix time

		outpoutconvert = mystart_date_job_level.Format(YYYYMMDD)
		//fmt.Println(mystart_date_job_level.Format(YYYYMMDD)) // outtput unix time to YYYY-MM-DD

		/*
			var cekint bool
			cekint = valid.IsInt(start_date_job_levelcell) //cek startdate dia int atau string

			if cekint == true { // jika start_date_job_levelcell bertipe int

				var i_start_date_job_levelcell int // buat variabel  dengan tipe data int

				start_date_job_levelconvint, _ := strconv.Atoi(start_date_job_levelcell) // parse string ke int

				i_start_date_job_levelcell = (start_date_job_levelconvint - 25569) * 86400 //perhitungan excel(int) - 25569 * 86400(selisih hari pada golang)

				stringstart_date_job_level := strconv.FormatInt(int64(i_start_date_job_levelcell), 10) //parse int ke string

				intstart_date_job_level, err := strconv.ParseInt(stringstart_date_job_level, 10, 64) //parse string ke time
				if err != nil {
					log.Fatal(err)
				}
				mystart_date_job_level := time.Unix(intstart_date_job_level, 0) // buat unix time

				fmt.Println(mystart_date_job_level.Format(YYYYMMDD)) // outtput unix time to YYYY-MM-DD
				fmt.Println(reflect.TypeOf(mystart_date_job_level))
			}
		*/
	} else {
		outpoutconvert = "isstring"
	}

	return outpoutconvert

}

func timestampnow() string {
	// Batas antara data import penginputan dengan
	locat, error := time.LoadLocation("Asia/Jakarta")
	if error != nil {
		panic(error)
	}
	tm := time.Now()
	timestamp := tm.In(locat).Format(YYYYMMDDHHMMSS24h)
	return timestamp
}

func routeSubmitImportExcel(w http.ResponseWriter, r *http.Request) {
	//Jika tidak ada status POST status
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	//Jika status
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadedFile, _, err := r.FormFile("file")

	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	xlsx, err := excelize.OpenReader(uploadedFile)

	sheet1Name := "PromotionMutation"

	var db, _ = Connect()
	defer db.Close()

	for index, _ := range xlsx.GetRows(sheet1Name) {

		tambah := index + 1

		// employee_id
		//nocell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", tambah))

		// number_of_employees
		number_of_employeescell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", tambah))

		// name
		//namecell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", tambah))

		// old_job_level
		old_job_levelcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", tambah))

		// new_Job_level
		new_job_levelcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", tambah))

		// start_date_job_level
		start_date_job_levelcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", tambah))

		//activity
		//activitycell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", tambah))

		// old_department
		old_departmentcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("H%d", tambah))

		// new_department
		new_departmentcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", tambah))

		//start_date_department
		start_date_departmentcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("J%d", tambah))

		//bagian
		bagiancell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", tambah))

		// cell
		cellcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("L%d", tambah))

		//remark
		remarkcell := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("M%d", tambah))

		if start_date_job_levelcell != "" && Convertdateyyyymmdd(start_date_job_levelcell) != "isstring" { //jika startdatejo ada maka cek format tanggal

			//membuar variabel penampung data dengan struct
			var count = Count{}

			// execution query
			//menampilkan data employee
			err = db.
				QueryRow("select count(id) as countid from employees where number_of_employees = ?", number_of_employeescell).
				Scan(&count.countid)

			if count.countid > 0 { // jika jumlah nik ada karena lebih dari 0
				var jobemp = Jobs{}
				var jobold = Jobs{}
				var joboldsel = Jobs{}
				var jobnewsel = Jobs{}
				var jobnew = Jobs{}

				var employee = Employees{}

				err = db.
					QueryRow("select id from employees where number_of_employees = ?", number_of_employeescell).
					Scan(&employee.id)

				if old_job_levelcell == "" { // jika old_job_levelcell kosong maka tampilkan job level sekarang berdasarkan tabel jobs

					//fmt.Printf("old_job_levelcell Kosong")

					//menampilkan data job_id berdasarkan karyawan sekarang
					err = db.
						QueryRow("select job_id as id from employees where number_of_employees = ?", number_of_employeescell).
						Scan(&jobemp.id)

					//menampilkan data old_job_level
					err = db.
						QueryRow("select id, level from jobs where id = ?", jobemp.id).
						Scan(&jobold.id, &jobold.level)

					//menampilkan data new_job_level
					err = db.
						QueryRow("select count(id) as id from jobs  where job_level = ?", new_job_levelcell).
						Scan(&jobnewsel.id)

					if jobnewsel.id > 0 { // job level baru ada maka exexute promosi demosi
						err = db.
							QueryRow("select id, level from jobs where job_level = ?", new_job_levelcell).
							Scan(&jobnew.id, &jobnew.level)

						if jobnew.level < jobold.level { // Jika new_job_level lebih besar dari old_job_level maka promosi

							// fmt.Printf("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, action, remark) values (%d,'%s',%d,%d,'%s', '%s')", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "promotion", remarkcell)
							_, err = db.Exec("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, activity, remark, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "promotion", remarkcell, timestampnow(), timestampnow())
							if err != nil {
								fmt.Println(err.Error())
								return
							}

							//fmt.Printf("update employees set job_id = %d where id = %d ", jobnew.id, employee.id)
							_, err = db.Exec("update employees set job_id = ? where id = ? ", jobnew.id, employee.id)
							if err != nil {
								fmt.Println(err.Error())
								return
							}
							fmt.Println("promotion success!")

							//fmt.Println("insert into promosimutations (employee_id, start_date_job_level, old_job_level, new_job_level, action, remark) values (%d,'%s',%d,%d,'%s', '%s')", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "promotion", remarkcell)

						} else if jobnew.level > jobold.level { // Jika new_job_level lebih kecil dari old_job_level maka demosi

							// fmt.Printf("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, action, remark) values (%d,'%s',%d,%d,'%s', '%s')", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "demotion", remarkcell)
							_, err = db.Exec("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, activity, remark, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "demotion", remarkcell, timestampnow(), timestampnow())
							if err != nil {
								fmt.Println(err.Error())
								return
							}

							// fmt.Printf("update employees set job_id = %d where id = %d ", jobnew.id, employee.id)
							_, err = db.Exec("update employees set job_id = ? where id = ? ", jobnew.id, employee.id)
							if err != nil {
								fmt.Println(err.Error())
								return
							}
							fmt.Println("demotion success!")

						} else if jobnew.level == jobold.level { // Jika new_job_level sama dengan old_job_level tidak ada action
							// fmt.Println("Old Job Level EQUAL New Job Level")
							http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data Promosi atau demosi "+number_of_employeescell+" tidak dapat di execute karena job levelnya sama !", http.StatusFound)

						} else { // Jika tidak terjadi maka kembali ke rule default
							// fmt.Println("Not execute")
							http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data Promosi atau demosi "+number_of_employeescell+" tidak dapat di execute karena data error !", http.StatusFound)
						}

					} else {
						http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data Promosi atau demosi "+number_of_employeescell+" tidak dapat di execute karena new_job_level tidak ada pada database !", http.StatusFound)
					}

				} else { // jika old_job_levelcell tidak kosong maka tampilkan berdasarkan tabel jobs

					//mencari job_level lama
					err = db.
						QueryRow("select count(id) as id from jobs where job_level = ?", old_job_levelcell).
						Scan(&joboldsel.id)

					//menampilkan data new_job_level
					err = db.
						QueryRow("select count(id) as id from jobs where job_level = ?", new_job_levelcell).
						Scan(&jobnewsel.id)

					if jobnewsel.id > 0 && joboldsel.id > 0 { // job level baru ada maka

						//fmt.Printf("new_job_levelcell Ada")

						//menampikan data old job level
						err = db.
							QueryRow("select id, level from jobs where job_level = ?", old_job_levelcell).
							Scan(&jobold.id, &jobold.level)

						err = db.
							QueryRow("select id, level from jobs where job_level = ?", new_job_levelcell).
							Scan(&jobnew.id, &jobnew.level)

						if jobnew.level < jobold.level { // Jika new_job_level lebih besar dari old_job_level maka promosi

							// fmt.Printf("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, action, remark) values (%d,'%s',%d,%d,'%s', '%s')", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "promotion", remarkcell)
							_, err = db.Exec("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, activity, remark, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "promotion", remarkcell, timestampnow(), timestampnow())
							if err != nil {
								fmt.Println(err.Error())
								return
							}

							// fmt.Printf("update employees set job_id = %d where id = %d ", jobnew.id, employee.id)
							_, err = db.Exec("update employees set job_id = ? where id = ? ", jobnew.id, employee.id)
							if err != nil {
								fmt.Println(err.Error())
								return
							}
							fmt.Println("promotion success!")

						} else if jobnew.level > jobold.level { // Jika new_job_level lebih kecil dari old_job_level maka demosi

							// fmt.Printf("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, action, remark) values (%d,'%s',%d,%d,'%s', '%s')", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "demotion", remarkcell)
							_, err = db.Exec("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, activity, remark, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "demotion", remarkcell, timestampnow(), timestampnow())
							if err != nil {
								fmt.Println(err.Error())
								return
							}

							// fmt.Printf("update employees set job_id = %d where id = %d ", jobnew.id, employee.id)
							_, err = db.Exec("update employees set job_id = ? where id = ? ", jobnew.id, employee.id)
							if err != nil {
								fmt.Println(err.Error())
								return
							}
							fmt.Println("demotion success!")

						} else if jobnew.level == jobold.level { // Jika new_job_level sama dengan old_job_level tidak ada action
							// fmt.Printf("Sama")
							http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data Promosi atau demosi "+number_of_employeescell+" tidak dapat di execute karena job levelnya sama !", http.StatusFound)

						} else { // Jika tidak terjadi maka kembali ke rule default
							// fmt.Printf("Else ")
							http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data Promosi atau demosi "+number_of_employeescell+" tidak dapat di execute karena data error !", http.StatusFound)

						}

					} else { // job level baru tidak ada maka
						http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data Promosi atau demosi "+number_of_employeescell+" tidak dapat di execute karena data error !", http.StatusFound)
					}

				}

			} else {
				http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data Promosi atau demosi "+number_of_employeescell+" tidak dapat di execute karena nik tidak di temukan !", http.StatusFound)
			}

		}

		// start_date_departmentcell
		if start_date_departmentcell == "" {

		} else {

			if Convertdateyyyymmdd(start_date_departmentcell) == "isstring" { // jika start_date_job_levelcell bertipe string
				http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data mutasi "+number_of_employeescell+" tidak dapat di execute karena format tanggal salah !", http.StatusFound)
			} else {

				//membuar variabel penampung data dengan struct
				var countempdept = Count{}

				var deptold = Departments{}
				var deptnewsel = Departments{}
				var deptoldsel = Departments{}
				var deptnew = Departments{}
				var employeedept = Employees{}

				// execution query
				//menampilkan data employee
				err = db.
					QueryRow("select count(id) as countid from employees where number_of_employees = ?", number_of_employeescell).
					Scan(&countempdept.countid)

				if countempdept.countid > 0 { // jika jumlah nik ada karena lebih dari 0

					err = db.
						QueryRow("select id from employees where number_of_employees = ?", number_of_employeescell).
						Scan(&employeedept.id)

					//menampilkan data new_department
					err = db.
						QueryRow("select count(id) as id from departments  where department = ?", new_departmentcell).
						Scan(&deptnewsel.id)

					if old_departmentcell == "" { //Jika old_department kosong

						//menampilkan data old_job_level
						err = db.
							QueryRow("select department_id as id from employees where id = ?", employeedept.id).
							Scan(&deptold.id)

						if deptnewsel.id > 0 { // jika department baru ada maka execute

							err = db.
								QueryRow("select id, department from departments where department = ?", new_departmentcell).
								Scan(&deptnew.id, &deptnew.department)

							if deptold.id != deptnew.id { // Jika department tidak sama maka execute mutasi

								// fmt.Printf("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, action, remark) values (%d,'%s',%d,%d,'%s', '%s')", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "promotion", remarkcell)
								_, err = db.Exec("insert into promotionmutations (employee_id, start_date_department, old_department, new_department, bagian, cell, activity, remark, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", employeedept.id, Convertdateyyyymmdd(start_date_departmentcell), deptold.id, deptnew.id, bagiancell, cellcell, "mutation", remarkcell, timestampnow(), timestampnow())
								if err != nil {
									fmt.Println(err.Error())
									return
								}

								_, err = db.Exec("update employees set department_id = ? where id = ? ", deptnew.id, employeedept.id)
								if err != nil {
									fmt.Println(err.Error())
									return
								}

								fmt.Println("Mutation Department NULL Success Insert in Database")
							}

						} else { // jika department baru tidak ada maka abaikan
							http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data mutasi "+number_of_employeescell+" tidak dapat di execute karena new_department tidak di temukan pada database !", http.StatusFound)

						}

					} else { //Jika old_department tidak kosong
						//menampilkan data new_department
						err = db.
							QueryRow("select count(id) as id from departments  where department = ?", old_departmentcell).
							Scan(&deptoldsel.id)

						if deptoldsel.id > 0 && deptnewsel.id > 0 { //Jika department old dan department new ada
							err = db.
								QueryRow("select id, department from departments  where department = ?", old_departmentcell).
								Scan(&deptold.id, &deptold.department)
							err = db.
								QueryRow("select id, department from departments where department = ?", new_departmentcell).
								Scan(&deptnew.id, &deptnew.department)

							if deptold.id != deptnew.id { // Jika department tidak sama maka execute mutasi

								// fmt.Printf("insert into promotionmutations (employee_id, start_date_job_level, old_job_level, new_job_level, action, remark) values (%d,'%s',%d,%d,'%s', '%s')", employee.id, Convertdateyyyymmdd(start_date_job_levelcell), jobold.id, jobnew.id, "promotion", remarkcell)
								_, err = db.Exec("insert into promotionmutations (employee_id, start_date_department, old_department, new_department, bagian, cell, activity, remark, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", employeedept.id, Convertdateyyyymmdd(start_date_departmentcell), deptold.id, deptnew.id, bagiancell, cellcell, "mutation", remarkcell, timestampnow(), timestampnow())
								if err != nil {
									fmt.Println(err.Error())
									return
								}
								_, err = db.Exec("update employees set department_id = ? where id = ? ", deptnew.id, employeedept.id)
								if err != nil {
									fmt.Println(err.Error())
									return
								}
								fmt.Println("Mutation Department Not NULL Success Insert in Database")
							}

						} else {
							w.Header().Set("Content-Type", "application/json")
							http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data mutasi "+number_of_employeescell+" tidak dapat di execute karena new_department tidak di temukan pada database !", http.StatusFound)
						}

					}

				} else {
					http.Redirect(w, r, "http://127.0.0.1:8000/notifications/promotionmutation/Data mutasi "+number_of_employeescell+" tidak dapat di execute karena NIK tidak di temukan pada database !", http.StatusFound)
				}

			}

		}

	}
}

// ASLI
func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("excelview.html"))
	var err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", routeIndexGet)

	// http.HandleFunc("/form", routeIndexGet)
	http.HandleFunc("/importexcelprocess", routeSubmitImportExcel)

	fmt.Println("server started at localhost:1001")
	http.ListenAndServe(":1001", nil)
}
