// Type Casting

package main

import (
	"fmt"
	"strconv"
)


func main(){
	fmt.Println("---- Type Casting -")

	// covert int to float
	var total int = 20;
	var items int = 7;

	var average float64;

	average = float64(total) / float64(items);

	fmt.Printf("Average is: %.2f\n", average)

	// converting float to int
	var price float64 = 19.99;

	var quantity int = int(price);

	fmt.Printf("Quantity: %d\n", quantity)

	// converting string to int
	str := "123"

	num, err := strconv.Atoi(str) 
	// if anything but a number in string is passed in strconv.Atoi function it will return error

	if err != nil {
		fmt.Println("error converting :", err)
	}

	if err == nil {
		fmt.Printf("Converted number: %v\n", num)
	}

}