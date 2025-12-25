//  RECURSION

/*
Recursion is when a function calls itself to solve a smaller version of the problem.

This continues until the problem becomes small enough that it can be solved directly. That smallest case is called the base case.
*/

// Base Case : Stops the recursion [preventing the infinite loop]
// Recursive Case : Calls itself

package main

import "fmt"


// factorial
func factorial(n int) int {
	if n == 1 {
		return 1 // base case
	} else {
		return n * factorial(n-1)
	}
}

// fibonacci
func fibonacci(n int) int {
	if n <= 2 {
		return 1 // base case
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}


func main(){
	fmt.Println("Recursion")
	fmt.Println(factorial(5))
	fmt.Println(fibonacci(10))
}