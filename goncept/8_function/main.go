// Functions - What are they
/* Function in go as in any other language is a reusable block of code that performs a specific task. Functions help organize our code and make it more modular and also allow for code reuse across a program which means that we can write  a function once and call it many times across our program. Also functions can accept parameters and return values.
 */

package main

import "fmt"

// General form
/* func functionName(parameters) returnType {
	 body of function
}
*/

func sayCapital(s string) {
	fmt.Printf("%s, is the Capital \n", s)
}
func sayCountry(s string) {
	fmt.Printf("%s, is the Country \n", s)
}

// function can take multiple parameter and return value
func finMaxNum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

// function can return multiple value
func swap(x, y string) (string, string){
	return y, x
}

// function arguments
// arguments are used when use a function, parameters are used when define or design a function

// call by value (default)
func increment(number int){
	number++
	fmt.Println("Inside increment: ", number)
}

// call by reference
func modify(slice []int){
	slice[0] = 999;
	fmt.Println("Inside modify", slice)
}

func main() {
	sayCountry("Bangladesh")
	sayCapital("Dhaka")

	a, b := 10, 20
  result := finMaxNum(a, b)

	fmt.Printf("The Max value is:%d\n", result)

	firstName, lastName := swap("Bran", "The Builder")

	fmt.Printf("Swaped names: %s %s\n", firstName, lastName)

	// calling increment funcion by value
	x := 10;
	increment(x);
	fmt.Println("In main after increment: ", x)

	// calling increment function by reference
  mySlice := []int{1, 2, 3}
	modify(mySlice)
	fmt.Println("In main after modify: ", mySlice)

}
