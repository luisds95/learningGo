package main

import "fmt"


func main() {
	arrayDemo()
	slicesDemo()
	mapsDemo()
}


func arrayDemo() {
	var arr [5]int
	arr[2] = 25
	arr[4] = 100
	fmt.Println(arr, "len:", len(arr))

	for i, value := range arr {
		fmt.Println(i, value)
	}

	anotherArr := [5]float64{ 98, 93, 77, 82, 83}
	total := 0.0
	for _, value := range anotherArr {
		total += value
	}
	fmt.Println("Avg of anotherArr:", total / float64(len(anotherArr)))
}


func slicesDemo() {
	slice := make([]float64, 5, 10)  // len 5, capacity 10
	fmt.Println(slice)

	anArray := [5]float64{1,2,3,4,5}
	anSlice := anArray[1:4]
	fmt.Println("Array:", anArray, ". Slice:", anSlice)

	appendedSlice := append(anArray[:], 6, 7)
	fmt.Println("Appended", appendedSlice)

	copiedSlice := make([]float64, 7)
	copy(copiedSlice, appendedSlice)
	fmt.Println("Copied", copiedSlice)
}


func mapsDemo() {
	var myMap = make(map[string]int)
	myMap["one"] = 1
	fmt.Println(myMap)

	// Delete key
	myMap["wrong"] = 99
	fmt.Println(myMap)
	delete(myMap, "wrong")
	fmt.Println(myMap)

	// Unexistent key
	fmt.Println(myMap["wrong"]) // Prints default
	value, ok := myMap["wrong"] // ok is false
	fmt.Println(value, ok)

	elements := map[string]map[string]string {
		"H": map[string]string {
			"name": "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string {
			"name":"Helium",
			"state":"gas",
		},
		"Li": map[string]string {
			"name":"Lithium",
			"state":"solid",
		},
		"Be": map[string]string {
			"name":"Beryllium",
			"state":"solid",
		},
		"B": map[string]string{
			"name":"Boron",
			"state":"solid",
		},
		"C": map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
		"N": map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O": map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
		"F": map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne": map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}