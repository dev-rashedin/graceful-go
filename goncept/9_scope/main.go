// Scope in Go.
// local, global and function parameter scope

package main

import "fmt"

// 1. Local Variables
// Local variables are declared within a function or block and are only accessible there.
func doSum()  {
var a, b, c int;
a = 10;
b = 20;
c = a + b;

fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}

// 2. Global Variables
// Global variables are declared outside of any function or block and are accessible anywhere in the program.

var g int

// Variable Shadowing
// If we declare two variable with the same name in the same scope, the second variable will shadow the first variable
// For example if we declare a global variable and a local variable with the same name, in the local scope, the local variable will be used, not the global variable

func main() {
doSum()
a := 10;
b := 20
g = a + b;

fmt.Printf("a = %d, b = %d, global variable g = %d\n", a, b, g)
}