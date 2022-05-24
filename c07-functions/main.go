package main

import "fmt"

func main() {
	numbers := []float64{5, 10, 15, 20, 25}
	result := mean(numbers)
	fmt.Println(result)

	five, six := multipleReturnValues()
	fmt.Println(five, six)

	totalAddition := add(numbers...)
	fmt.Println(totalAddition)

	evenGenerator := makeEvenGenerator()
	for i := 0; i < 5; i++ {
		fmt.Println(evenGenerator())
	}

	/*  Defer. Close will happen when the function finishes
	file, _ := os.Open(filename)
	defer file.Close()
	*/
	recoverFromPanic()
}

func mean(numbers []float64) float64 {
	var total float64 = 0
	for _, n := range numbers {
		total += n
	}
	return total / float64(len(numbers))
}

func multipleReturnValues() (int, int) {
	return 5, 6
}

func add(args ...float64) (total float64) {
	total = 0
	for _, value := range args {
		total += value
	}
	return
}

func makeEvenGenerator() func() uint {
	var next uint = 0
	return func() (this uint) {
		this = next
		next += 2
		return
	}
}

func recoverFromPanic() {
	getValueBack := func() {
		value := recover()
		fmt.Println(value)
	}
	defer getValueBack()

	fmt.Println("Panicking")
	panic("some value")
}