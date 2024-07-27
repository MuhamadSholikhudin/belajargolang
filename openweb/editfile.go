package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// Buat Tanggal Hari ini

//

func Modify() {

	tpm := time.Now()
	locationpm, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println(err.Error())
	}
	datetimenow := tpm.In(locationpm).Format("2006-01-02")

	// OpenFile with Read Write Mode
	file, err := os.OpenFile("backuptransit.bat", os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("open file failed: %s", err)
	}
	defer file.Close()

	// txt := fmt.Sprintf("copy %s.txt Z:%sIT%sUBACKUPDAILY", datetimenow, "\\", "\\")
	lenUpdate, err := file.WriteAt([]byte(fmt.Sprintf("copy %s.txt Z:%sIT%stransitbackup", datetimenow, "\\", "\\")), 0) // Write at 0 beginning
	if err != nil {
		log.Fatalf("write to file failed: %s", err)
	}
	fmt.Printf("Length update: %d bytes\n", lenUpdate)

	lenUpdate2, err := file.WriteAt([]byte("cd..\n"), 2) // Write at 0 beginning
	if err != nil {
		log.Fatalf("write to file failed: %s", err)
	}
	fmt.Printf("Length update: %d bytes\n", lenUpdate2)
	fmt.Printf("File Name: %s\n", file.Name())
}

func ReadFile() {
	data, err := os.ReadFile("backuptransit.bat")
	if err != nil {
		log.Panicf("read data from file failed: %s", err)
	}

	fmt.Printf("File Content: %s\n", string(data[:]))
}

func StartBackup(app string) {
	cmd := exec.Command(app)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

//=================================++++++++++++++++++=========================

func main() {
	fmt.Println("Read file before modify")
	ReadFile()

	fmt.Println("Using WriteAt function")
	Modify()

	fmt.Println("Read file after modify")
	ReadFile()

	// fmt.Println("RUN SCRIPT BACK UP")
	// StartBackup("backuptransit.bat")
}

// func main() {
// 	sampledata := []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
// 		"Nunc a mi dapibus, faucibus mauris eu, fermentum ligula.",
// 		"Donec in mauris ut justo eleifend dapibus.",
// 		"Donec eu erat sit amet velit auctor tempus id eget mauris.",
// 	}

// 	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

// 	if err != nil {
// 		log.Fatalf("failed creating file: %s", err)
// 	}

// 	datawriter := bufio.NewWriter(file)

// 	for _, data := range sampledata {
// 		_, _ = datawriter.WriteString(data + "\n")
// 	}

// 	datawriter.Flush()
// 	file.Close()
// }
