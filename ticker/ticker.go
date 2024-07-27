package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("RUN PROGRAM")
	GetGoogleSheet := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-GetGoogleSheet.C:
			fmt.Println("Execute on going process")
		}
	}

}
