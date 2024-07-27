package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	i := 0
	for {
		fmt.Printf("-> Send Job: %d\n", i)
		ch <- i
		i++
	}
}

func echoworker(in, out chan int) {
	for {
		n := <-in
		//process
		time.Sleep(1000 * time.Microsecond)
		//output
		out <- n
	}

}

func main() {

	in := make(chan int)
	out := make(chan int)

	//initialization worker pool
	for i := 0; i < 4; i++ {
		go echoworker(in, out)
	}

	//input
	go producer(in)

	//output
	for n := range out {
		fmt.Printf("<--- Receive job: %d\n", n)
	}
}
