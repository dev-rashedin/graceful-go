// Understanding Pointers

// A pointer stores the memory address of another variable

// Declaring Pointers
// var ip *int 
// vr fp *float32

// Basic Pointer Operations
// Storing the address: & to get the address of a variable
// Dereferencing: * to access the value at the pointer's address


package main

import "fmt"

func main() {
	var a int = 32
	var ip *int
	ip = &a;

	fmt.Println("The address of a is: ", ip)
	fmt.Println("The value of a is: ", *ip)

	// Nil Pointers
	var ptr *int
	fmt.Println("The value of ptr is: ", ptr)

}