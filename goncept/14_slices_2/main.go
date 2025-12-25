// SLICES
// Slices are more dynamic and more flexible  than arrays data type. Simply a slice can  allow you to create sequences that can change in size if you want
// Similar to arrays in javaScript and lists in python

package main

import "fmt"


func main(){
	fmt.Println("SLICES in Go")
	// declaring a slice
	var numbers []int 

	// initialization of slice of integers with length 5 filled with [0, 0, 0, 0, 0]
	numbers = make([]int, 5)

	fmt.Println("The slice", numbers)
	fmt.Println("The length of the slice is: ", len(numbers))

	// declaring and initializing a slice in one line
	// numbers2 := make([]int, 5)

	// Subslicing
	numbers3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("The slice", numbers3[3:8])
	fmt.Println("The slice2", numbers3[:3])
	fmt.Println("The slice3", numbers3[4:])

	// Append() and Copy()
	new_numbers := []int{}
	new_numbers = append(new_numbers, 0, 1, 2, 3, 4, 5, 6)
		
	new_numbers1:= make([]int, len(numbers), cap(numbers) * 2)

	fmt.Println("new numbers slice with append", new_numbers)
	fmt.Println("new_numbers1 slice with append", new_numbers1)

	// copying numbers
	copy(new_numbers1, new_numbers)
	fmt.Println("new_numbers1 slice after copy", new_numbers1)

	// Nil Slice
	var num []int // nil slice
	fmt.Println("Nil Slice", num)

	if num == nil {
		fmt.Println("Slice is nill")
	}

}