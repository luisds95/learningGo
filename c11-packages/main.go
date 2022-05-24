package main

import "fmt"
import "go_intro/c11-packages/math"

func main() {
	numbers := []float64{1,2,3,4}
	avg := math.Average(numbers)
	fmt.Println(avg)
}