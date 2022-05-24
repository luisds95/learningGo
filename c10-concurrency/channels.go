package main

import (
	"fmt"
	"time"
)

func main() {
	var channel chan int = make(chan int) // Synchronous
	// var channel chan int = make(chan int, 5) // Asynchronous
	go pinger(channel)
	go printer(channel)

	var input string
	fmt.Scanln(&input)
}

func pinger(channel chan<- int) {
	for i := 0; i < 10; i++ {
		channel <- i
		// time.Sleep(time.Second * 1)
	}
}

func printer(channel <-chan int) {
	for {
		value := <- channel
		fmt.Println("Value:", value)
		time.Sleep(time.Millisecond * 500)
	}
}