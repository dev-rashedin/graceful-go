// Error Handling

/*
Errors
Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
A nil error denotes success; a non-nil error denotes failure.
*/

package main

import (
	"errors" // imported to create error message
	"fmt"
	"math"
	"time"
)

// The error interface
type error interface {
	Error() string
}


// Sqrt calculate the square root of a function
func Sqrt(value float64) (float64, error) {
if value < 0 {
	return 0, errors.New("Negative number passed to Sqrt!")
}

// Return the square root and a nill error
return math.Sqrt(value), nil
}

// official error docs
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}


func main(){
  // Test the Sqrt function with a negative value
	result, err := Sqrt(-5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Test the Sqrt function with a positive value
  result, error := Sqrt(5);

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}

	// Official error docs
if err := run(); err != nil {
		fmt.Println(err)
	}

}