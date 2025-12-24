// Strings

package main


import "fmt"
import "strings"


func main() {

	// Creating Strings
	var greeting string = "Hello, World!"

	fmt.Printf("Normal String: ")
	fmt.Printf("%s:", greeting)
	fmt.Printf("\n")


	// Display the hexadecimal byte values of the string
	fmt.Printf("HEX bytes: ")

	for i :=0; i< len(greeting); i++ {
		fmt.Printf("%x ", greeting[i]) // x is for hexadecimal
	}
	fmt.Printf("\n")

	// Measuring String length
	fmt.Printf("The length of this string is : ")
	fmt.Println(len(greeting))

	// String concatenation
	// create a slice of strings for concat
	fruits := []string{"apples", "oranges", "and bananas"}
	// concat. using strings.Join
	fmt.Println(strings.Join(fruits, " "))
}