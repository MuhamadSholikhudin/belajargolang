package repository

import (
	"belajargolang/api/resign/models"
	"fmt"
	"reflect"
)

func InsertResign(table string, data map[string]interface{}) {
	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()

	keymap := fmt.Sprintf("")
	valmap := fmt.Sprintf("")
	for key, val := range data {
		keymap = fmt.Sprintf("%s%v,", keymap, key)
		if reflect.TypeOf(val).Kind() == reflect.Int {
			valmap = fmt.Sprintf("%s%d,", valmap, val)
		} else {
			valmap = fmt.Sprintf("%s'%s',", valmap, val)
		}
	}
	keymap = keymap[:len(keymap)-1]
	valmap = valmap[:len(valmap)-1]
	query := fmt.Sprintf("INSERT %s (%s) VALUES (%s)", table, keymap, valmap)
	_, err = dbresign.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func UpdateResign(table string, data map[string]interface{}, where string) {
	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()
	valmap := fmt.Sprintf("")
	for key, val := range data {
		if reflect.TypeOf(val).Kind() == reflect.Int {
			valmap = fmt.Sprintf("%s %v = %d,", valmap, key, val)
		} else {
			valmap = fmt.Sprintf("%s %v = '%s',", valmap, key, val)
		}
	}
	valmap = valmap[:len(valmap)-1]
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, valmap, where)
	_, err = dbresign.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func DeleteResign(table string, where string) {
	dbresign, err := models.ConnResign()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbresign.Close()
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, where)
	_, err = dbresign.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func CountResign(table string, where string) int {
	dbresign, _ := models.ConnResign()
	defer dbresign.Close()
	var id int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s", table, where)
	_ = dbresign.QueryRow(query).Scan(&id)
	return id
}

/*
func QueryRow(table string, where string) map[string]interface{} {
	dbresign, _ := models.ConnResign()
	defer dbresign.Close()
	var id int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s", table, where)
	_ = dbresign.QueryRow(query).
		Scan(&id)
	var data := map[string]interface{}{

	}
	return data
}
*/
