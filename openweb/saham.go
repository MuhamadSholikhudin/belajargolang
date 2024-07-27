package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	/*
		f, err := os.Open("shares.txt")
		if err != nil {
			log.Fatal(err)
		}
		// remember to close the sfile at the end of the program
		defer f.Close()

		// read the file line by line using scanner
		scanner := bufio.NewScanner(f)

		var dataset string
		for scanner.Scan() {
			// do something with a line
			dataset = scanner.Text()
		}

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
					perulangan, _ := strconv.Atoi(string(dataset[20:22]))

					for i := 0; i < perulangan; i++ {
						fmt.Println("Open aplikasi", (i + 1))
						OpenSaham(fmt.Sprint(dataset[23:38]))
					}
					panic(err.Error())
				} else {
					d2 = "true"
				}
				fmt.Println(d2)
			}
		}
	*/
	resp, err := http.Get("https://sholikhudin11.000webhostapp.com/?api=settingapp")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	UpdateSpreadsheet := time.NewTicker(1 * time.Second)
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
			request_date, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprint(data[2]["app_start"]))
			fmt.Println(timesecond.Unix())
			fmt.Println(request_date.Unix())
			fmt.Println(datetimenow2.Unix())
			var d2 string
			if datetimenow2.Unix() > request_date.Unix() {
				d2 = "false"
				perulangan, _ := strconv.Atoi(string(fmt.Sprint(data[2]["amount"])))
				for i := 0; i < perulangan; i++ {
					fmt.Println("Open aplikasi", (i + 1))
					OpenSaham(fmt.Sprint(data[2]["name_exe"]))
				}
				panic(err.Error())
			} else {
				d2 = "true"
			}
			fmt.Println(d2)
		}
	}
}

func OpenSaham(app string) {
	cmd := exec.Command(app)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
