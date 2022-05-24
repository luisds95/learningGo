package main

import "fmt"

func main() {
	value := 5
	makeOne(&value)
	fmt.Println("Value:", value, ". Memory:", &value)

	alsoAPointer := new(int)
	makeOne(alsoAPointer)
	fmt.Println("alsoAPointer:", *alsoAPointer)

	left := 1
	right := 2
	swapInts(&left, &right)
	fmt.Println(left, right)
}

func makeOne(valuePtr *int) {
	*valuePtr = 1
}

func swapInts(x *int, y *int) {
	mid := *x
	*x = *y
	*y = mid
}