package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	OpenUrl()
	tickerclose := time.NewTicker(8 * time.Minute)

	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Minute) // wait for 10 seconds
		done <- true
	}()

	for {
		select {

		case <-done:
			fmt.Println("Goroutine Execute")
		case TC := <-tickerclose.C:
			CloseUrl()
			fmt.Println("Hello CLOSE!!", TC)
			OpenUrl()
			fmt.Println("Hello OPEN", TC)
		}
	}
}

func OpenUrl() {
	cmd := exec.Command("urlopen.bat")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

func CloseUrl() {
	cmd := exec.Command("urlclose.bat")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
