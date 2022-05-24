package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	channel1 := make(chan string)
	channel2 := make(chan string)

	go Pinger("Ping", 1100, channel1)
	go Pinger("Pong", 1500, channel2)

	for {
		select {
		case msg1 := <- channel1:
			fmt.Println("Msg 1", msg1)
		case msg2 := <- channel2:
			fmt.Println("Msg 2", msg2)
		case <- time.After(time.Millisecond * 1100):
			fmt.Println("timeout")
			return
		// default:
		// 	fmt.Println("Waiting")
		}
	}
}

func Pinger(msg string, ms int, channel chan<- string) {
	for {
		channel <- msg
		waitTime := time.Millisecond * time.Duration(rand.Intn(1000) - 500 + ms)
		time.Sleep(waitTime)
	}
}