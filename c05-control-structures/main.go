package main

import "fmt"


func main() {
	loops()
	fizzBuzz()
}


func loops() {
	initialValue := 1
	finalValue := 5

	i := initialValue
	for i <= finalValue {
		fmt.Println(i)
		i += 1
	}

	for j := initialValue; j <= finalValue; j++ {
		fmt.Println(j)
	}
}


func fizzBuzz() {
	n := 10
	
	for i := 1; i <= n; i++ {
		word := ""
		if i % 2 == 0 {
			word += "Fizz"
		}
		if i % 5 == 0 {
			word += "Buzz"
		}
		
		if word == "" {
			word = string(i)
		} else {
			fmt.Println(word)
		}
	}
}