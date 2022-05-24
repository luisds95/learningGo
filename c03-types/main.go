package main

import "fmt"

/*
types:
int, uint
float32, float64

*/

func main() {
	fmt.Println("1 + 1 =", 1.0 + 1.0)
	fmt.Println("Len of str:", len("my str"))
	fmt.Println("First letter", "sentence"[0])
	fmt.Println("String" + "concatenation")
	fmt.Println(true && false, true || false, !true)
}