package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	tickerclose := time.NewTicker(1 * time.Second)

	done := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second) // wait for 10 seconds
		done <- true
	}()

	for {
		select {

		case <-done:
			fmt.Println("Goroutine Execute")
		case TC := <-tickerclose.C:
			Tel()
			fmt.Println("Hello CLOSE!!", TC)
		}
	}

}

func Tel() {
	url := "https://api.telegram.org/bot5741242750%3AAAGbZr_G2U8ctgvSyuMuIZaEAMvkiMoH8VQ/getUpdates"

	payload := strings.NewReader("{\"offset\":null,\"limit\":null,\"timeout\":null}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
