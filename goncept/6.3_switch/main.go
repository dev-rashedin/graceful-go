package main

import "fmt"

func main() {
	age := 21

	switch {
	case age >= 21:
		fmt.Println("You are eligible for marriage")
	case age >= 18:
		fmt.Println("You are eligible for voting")
	default:
		fmt.Println("You are not eligible for anything")
	}

	score := 82

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B+")
	case score >= 70:
		fmt.Println("Grade: B")
	case score >= 60:
		fmt.Println("Grade: C+")
	case score >= 50:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}
}
