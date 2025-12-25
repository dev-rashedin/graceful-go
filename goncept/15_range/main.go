


package main

import "fmt"

func main() {
  fmt.Println("Range")
	// Using Range with Slice and Array

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7} 

	for i, number := range numbers {
		fmt.Printf("Slice item %d is %d\n", i, number)
	} // it is like enumerate in python

	n := []int{1, 2, 3}

	for i := range n {
		fmt.Printf("Array item  %d is %d\n", i, n[i])
	}
}