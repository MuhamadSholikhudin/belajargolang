// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

const (
	YYYYMMDD = "2006-01-02"
)

func main() {

	var time1 = time.Now()
	fmt.Println("Now :", time1.Format(YYYYMMDD))
	// time1 2015-09-01 17:59:31.73600891 +0700 WIB

	// rangedategolang := 86400
	// date_1900 := -2208988800

	//Konversi int ke string
	// strconv.Itoa(value)

	// numberexcel := 3000

	var time2 = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("T  %v\n", time2)
	fmt.Printf("D %v\n", time2.Unix())

	// i, err := strconv.ParseInt(date_excel, 0, 0)
	// if err != nil {
	// 	panic(err)
	// }
	// tm := time.Unix(i, 0)
	// fmt.Println("parse :", tm.Format(YYYYMMDD))

}
