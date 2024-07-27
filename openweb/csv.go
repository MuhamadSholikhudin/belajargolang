package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	// open file
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var dataset string
	for scanner.Scan() {
		// do something with a line
		dataset = scanner.Text()
	}
	fmt.Println(len(dataset))
	fmt.Println(dataset[0:19])
	fmt.Println(dataset[20:22]) // Looping
	fmt.Println(dataset[23:25]) // Interval

	// fmt.Println("RUN PROGRAM")
	// ExecuteSpreadsheet()

	// CheckSyncron := time.NewTicker(2 * time.Second)
	UpdateSpreadsheet := time.NewTicker(1 * time.Second)
	// UpdateMasterToResign := time.NewTicker(2 * time.Minute)
	for {
		select {
		case timesecond := <-UpdateSpreadsheet.C:
			fmt.Println()

			t := time.Now()
			location, err := time.LoadLocation("Asia/Bangkok")
			if err != nil {
				fmt.Println(err)
			}
			datetimenow := t.In(location).Format("2006-01-02 15:04:05")
			datetimenow2, _ := time.Parse("2006-01-02 15:04:05", datetimenow)

			// ExecuteSpreadsheet()
			request_date, _ := time.Parse("2006-01-02 15:04:05", dataset[0:19])
			fmt.Println(timesecond.Unix())
			fmt.Println(request_date.Unix())
			fmt.Println(datetimenow2.Unix())

			var d2 string
			if datetimenow2.Unix() > request_date.Unix() {
				d2 = "false"
				cmd := exec.Command("auto_try.exe")
				stdoutStderr, err := cmd.CombinedOutput()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s\n", stdoutStderr)
				panic(err.Error())
			} else {
				d2 = "true"
			}
			fmt.Println(d2)
		}
	}

}
