package main

import (
	"fmt"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	//runtime.GOMAXPROCS(2)

	go print(50, "HALO")
	print(50, "Apa kabar")
	time.Sleep(5 * time.Second)
	fmt.Println("Hello Word")

	var s1, s2, s3 string
	fmt.Scanln(&s1, &s2, &s3)

	// user inputs: "trafalgar d law"

	fmt.Println(s1) // trafalgar
	fmt.Println(s2) // d
	fmt.Println(s3) // law

}
