package repository

import (
	"belajargolang/api/resign/models"
	"fmt"
	"reflect"
)

func UpdateHrd(table string, data map[string]interface{}, where string) {
	dbhrd, err := models.ConnHrd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dbhrd.Close()
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
	_, err = dbhrd.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
