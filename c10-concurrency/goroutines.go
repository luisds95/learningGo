package main

import (
	"fmt" 
	"time" 
	"math/rand"
)

func main() {
	// Send 5 goroutines
	for i := 0; i < 5; i++ {
		go printTen(i)
	}	
	var input string
	fmt.Scanln(&input)
}

func printTen(process int) {
	for i := 0; i < 10; i++ {
		waitTime := time.Millisecond * time.Duration(rand.Intn(250))
		fmt.Println("Process:", process, "Iteration:", i, "Wait time:", waitTime)
		time.Sleep(waitTime)
	}
}