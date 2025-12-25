// Interfaces

/*
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

*/

// Interface is similar to struct, in interface we pass methods and in struct we pass variable

package main

import (
	"fmt"
	"math"
)

// Define an interface
type thisIsAnInterface interface {
	// methods
	// MethodName1 ReturnType
	// MethodName2 ReturnType
	// MethodName3 ReturnType
}

type Shape interface {
	area() float64
}

// Define a struct
type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

// Implement area method for Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implement area method for Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Interfaces")

	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 20}

	fmt.Printf("Circle are: %f\n", getArea(circle))
	fmt.Printf("Rectangle are: %f\n", getArea(rectangle))

}
