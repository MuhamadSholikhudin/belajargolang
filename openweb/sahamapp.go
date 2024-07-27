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

	resp, err := http.Get("https://sholikhudin11.000webhostapp.com/?api=settingapp")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
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
	fmt.Println(data[0]["app_start"])

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
			request_date, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprint(data[0]["app_start"]))
			fmt.Println(timesecond.Unix())
			fmt.Println(request_date.Unix())
			fmt.Println(datetimenow2.Unix())

			var d2 string
			if datetimenow2.Unix() > request_date.Unix() {
				d2 = "false"
				perulangan, _ := strconv.Atoi(string(fmt.Sprint(data[0]["amount"])))

				for i := 0; i < perulangan; i++ {
					fmt.Println("Open aplikasi", (i + 1))
					//OpenSaham(fmt.Sprint(data[0]["name_exe"]))
				}
				//panic(err.Error())
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
