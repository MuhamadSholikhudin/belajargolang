package main

import "fmt"

func main() {
	/*
		// 1. Channel
		runtime.GOMAXPROCS(2)

		var messages = make(chan string)

		var sayHelloTo = func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}

		go sayHelloTo("john wick")
		go sayHelloTo("ethan hunt")
		go sayHelloTo("jason bourne")
		go sayHelloTo("Udin")
		go sayHelloTo("1")

		var message1 = <-messages
		fmt.Println(message1)

		var message2 = <-messages
		fmt.Println(message2)

		var message3 = <-messages
		fmt.Println(message3)
		var message4 = <-messages
		fmt.Println(message4)
		var message5 = <-messages
		fmt.Println(message5)
	*/

	// 2. Channel sebagai Parameter

	var messages = make(chan string)

	for i, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string, index int) {
			var data = fmt.Sprintf("hello %s dan %d", who, index)
			messages <- data
		}(each, i)
		// printMessage(messages)
	}
	var tryk []string
	for i := 0; i < 3; i++ {
		// printMessage(messages)
		tryk = append(tryk, printMessage(messages))
	}
	fmt.Println(tryk)
}

func printMessage(what chan string) string {
	// fmt.Println(<-what)
	var oke string
	oke = <-what
	return oke
}

// func printMessage(what chan string) {
// 	fmt.Println(<-what)
// }
