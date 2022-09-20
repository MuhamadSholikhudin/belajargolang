// Golang program to illustrate how to
// remove files from the default directory
package main

import (
	"log"
	"os"
)

func main() {

	hapus()
}

func hapus() {
	// Removing file from the directory
	// Using Remove() function
	e := os.Remove("./files/file1.xlsx")
	if e != nil {
		log.Fatal(e)
	}
}
