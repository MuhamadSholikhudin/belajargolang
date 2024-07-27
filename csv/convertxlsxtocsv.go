package main

import (
	"fmt"
	"os"
)

func main() {
	oldFileName := "file1.xlsx"
	newFileName := "new_file.csv"

	err := os.Rename(oldFileName, newFileName)
	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}

	fmt.Println("File renamed successfully.")
}
