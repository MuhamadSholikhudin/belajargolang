package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id                  int
	name                string
	number_of_employees string
}

func main() {
	// query banyak data
	// sqlQuery()

	// query satu
	// sqlQueryRow()

	//query execute insert, update, delete
	// sqlExec()
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hrdit")
	if err != nil {
		return nil, err
	}

	return db, nil
}

/*
// Query banyak data
func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 0
	rows, err := db.Query("select id, name, number_of_employees from employees where id > ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.number_of_employees)

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

	for _, each := range result {
		fmt.Println(each.number_of_employees, each.name)
	}
}

*/
/*
//  query menampilkan satu data
func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var id = 2
	err = db.
		QueryRow("select id, name, number_of_employees  from employees where id = ?", id).
		Scan(&result.id, &result.name, &result.number_of_employees)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf(" id: %d\n name: %s\n number_of_employees: %s\n", result.id, result.name, result.number_of_employees)
}
*/

// Insert, Update, & Delete Data Menggunakan
/*
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()


		_, err = db.Exec("insert into employees (id, number_of_employees, name) values (?, ?, ?)", 1, "1", "1")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("insert success!")

		_, err = db.Exec("update employees set number_of_employees = ? ,name = ? where id = ?", "2207057000", "APRIZAL", 1)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("update success!")

	_, err = db.Exec("delete from employees where id = ?", 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")

}

*/
