package main

import ("fmt"; "math")

func main(){
	mainCircle := Circle{x: 0, y: 0, radius: 5}
	fmt.Println("Circle:", mainCircle)
	fmt.Println("Radius:", mainCircle.radius)
	fmt.Println("Area:", circleArea(mainCircle))
	fmt.Println("Perimeter:", mainCircle.perimeter())

	luis := Person{name: "Luis"}
	luis.greet()
	android := new(Android)
	android.name = "Android"
	android.greet()
}

type Shape interface {
	perimeter() float64
}

type Circle struct {
	x, y, radius float64
}

func circleArea(circle Circle) float64 {
	return math.Pi * circle.radius * circle.radius
}

// (circle *Circle) is a receiver
func (circle *Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

type Person struct {
	name string
}
func (person *Person) greet() {
	fmt.Println("Hi! My name is", person.name)
}

type Android struct {
	Person
}