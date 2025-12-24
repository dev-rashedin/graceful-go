//  Constants in go programming language
//  // Constants are immutable values known at compile time and cannot be changed.

package main

import "fmt"

func main(){
	
	//  Declare strings and numeric constants
	const NAME string = "John Snow"
	const AGE int = 24

	fmt.Println("Name is", NAME)
	fmt.Println("Age is", AGE)
	// Declare the integer literals
	// An integer literal can be decimal, octal, or hexadecimal
	// 0x --> hex, 0o --> octal, 0b --> binary

	const (
		DECIMAL = 255 // decimals with no prefix
		OCTAL = 0o377 // octal with leading 0
		Hexadecimal = 0xff // hexadecimal with leading 0x
	)

	fmt.Println("Decimal:", DECIMAL, "Octal:", OCTAL, "Hexadecimal:", Hexadecimal)

	// Floating-point Literals
	// A floating-point literal can have an integer part, a decimal part, a fractional part and an exponent part.

	const PI float64 = 3.14159
	const AVOGADRO float64 = 6.02214076e23

	fmt.Println("PI:", PI, "Avogadro:", AVOGADRO)


	// Escape Sequences in string literals
  
	// Newline \n
	// Quote \"
	// Backslash \
	// Alert or Bell \a
	// Backspace \b
	// Form Feed \f
	// Newline \n
	// Carriage Return \r
	// Horizontal Tab \t
	// Vertical Tab \v

	// This is new line
	const GREETING = "Hello, Earth!\n";

	// This is a quote
	const QUOTE = "\"Life is what happens to you while you're busy making other plans.\" - John Lennon";

	fmt.Println(GREETING);
	fmt.Println(QUOTE);

	const BELL = "Bell is \a"

	fmt.Println(BELL);

	// Multi-line and Concatenated String Literals
  const MULTILINE = "John Snow is now the King in the North." + " He is actually a Targaryen and not a Bastard." + " Therefore, he is the true heir of the Iron Throne."

	fmt.Println(MULTILINE)


	// Boolean Constants
	const ACTIVE = true
	const READY = false

	fmt.Println(ACTIVE)

	// Constants for calculations
	const LENGTH = 50
	const WIDTH = 25
	const AREA = LENGTH * WIDTH

	fmt.Println("The area of the rectangle is:", AREA)
}