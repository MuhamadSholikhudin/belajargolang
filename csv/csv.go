package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
)

func main() {

	file, err := os.Open("Export.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	fmt.Println(reflect.TypeOf(records))
	for _, record := range records {
		fmt.Println(record[0], record[1], record[2], record[3], record[4], record[5])
		// 	for ind, rec := range record {
		// 		fmt.Print("Index Array => ", i, " Ini ind ke-2 => ", ind, " Ini rec => ", rec)
		// 	}
		// 	fmt.Println()
	}
}
