// Understanding Arrays in Golang
/*
Fixed Size: The size of an array is fixed at declaration.
Type-Specific: All elements must be of the same type.
Contiguous Memory: Provides fast access to elements
Zero Indexing: The first element is accessed with index 0
*/

package main

import "fmt"

// Declaring Arrays
var balance [10]float32

// Initializing Array at declaration
var newBalance = [5]float32{100.00, 2.00, 3.00, 4.00, 5.00}

// Initializing while omitting the size of the array elements
var newBalance2 = [...]float32{100.00, 2.00, 3.00, 4.00, 5.00}

var grades = [4]float64{97, 85, 93, 78}

var names = [3]string{"John", "Jane", "Joe"}

var programmingLang = [3]string{"JavaScript", "Python", "Go"}


func main(){
	var n [10]int
	// Initializing array elements
	for i :=0; i < 10; i++{
		n[i] = i + 100
	}


	fmt.Println("Array Elements: ", n)

	// Displaying array elements
	for i := 0; i < 10; i++{
		fmt.Printf("Element[%d] = %d\n", i, n[i])
	}
}
