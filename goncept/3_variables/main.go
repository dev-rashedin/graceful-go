// Understanding Variables Declaration in Go
package main

import "fmt"

func main() {

	// Go is a statically typed language. That means that we need to declare the type of a variable before we can use it.

	// If we don't declare the type, Go will infer the type of the variable based on the value assigned to it

	// way 1: Declare and assign on 2 different lines
	var name1 string
	name1 = "John Snow"

	// way 2: Declare and assign on the same line
	var name2 string = "Arya Stark"

	fmt.Println(name1, name2)

	// way 3: Short declaration operator
	// Type inference : Go will infer the type of the variable based on the value assigned to it
	name3 := "Bran Stark"
	fmt.Println(name3)

	house := "Winterfell"
	fmt.Println(house)

	// We can declare two variable of the same type on the same line
	var name4, name5 string = "Daenerys Targaryen", "Melisandre"
	fmt.Println(name4, name5)

	city, country := "King's Landing", "Westeros"
	fmt.Println("The king lives in", city, "of", country)

	var num1, num2 int8 = 10, 20
	fmt.Println(num1 + num2) // 30

	var windows, mac, linux string = "Windows is ok", "Mac is meh", "Linux is GOAT!"

	fmt.Println(windows, mac, linux)


	// We can check they type of our assigned variable
var x float64 = 24
fmt.Printf("x is of type: %T\n", x)

// Mixed Variable Declaration
var a, b, c = 5, 8, "Hello"

fmt.Printf("a is of type %T\n", a)
fmt.Printf("b is of type %T\n", b)
fmt.Printf("c is of type %T\n", c)

	// We cannot redeclare a variable. This will throw an error
	// We cannot change the type of a variable but we can change the value
	// We also cannot have any uninitialized or unused variables in Go
}
