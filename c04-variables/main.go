package main

import "fmt"


func main() {
	howToDeclareVars()
	printDouble()
}


func howToDeclareVars() {
	var hello string = "Hello :)"
	itsMe := "again"
	fmt.Println(hello, itsMe)

	const myConstant string = "I'm constant"
	fmt.Println(myConstant)

	var (
		yet = "Yet"
		more = "more"
		vars = "vars"
	)
	fmt.Println(yet, more, vars)
}


func printDouble() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println("Its double is:", output)
}