/*
* Http (curl) request in golang
* @author Shashank Tiwari
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	url := "https://imedia-smm.com/api/profile"

	payload := strings.NewReader("api_id=55171&api_key=f851ed-b7d25f-769e65-ebd055-f53844")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(string(body))

	var data map[string]interface{}
	var err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data)

	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(resp))

}

func main() {
	http.HandleFunc("/profile", Profile)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
