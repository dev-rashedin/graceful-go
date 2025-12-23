// Control Flow:
// Conditionals (If Else Statement)


package main

import "fmt"

func main() {

// If Statement
// Executes a block of code if the condition is true

age := 21

if age >= 18 {
	fmt.Println("You are eligible to vote")
} else {
	fmt.Println("You are not eligible to vote")
}

score :=82;

if score >= 90 {
	fmt.Println(("Grade: A"))
} else if score >= 80 {
	fmt.Println(("Grade: B+"))
} else if score >= 70 {
	fmt.Println(("Grade: B"))
} else if score >= 60 {
	fmt.Println(("Grade: C+")) 
} else if score >= 50 {
	fmt.Println(("Grade: C"))
} else {
	fmt.Println(("Grade: F"))
}

} 