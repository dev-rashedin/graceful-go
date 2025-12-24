// Operators in Go Language
// An operator is a symbol that performs specific mathematical or logical operations on operands.

// Arithmetic Operators
// Relational Operators
// Logical Operators
// Bitwise Operators
// Assignment Operators
//  & Miscellaneous Operators

package main

import "fmt"

func main() {
	// Arithmetic Operators
	// +, -, *, /, %
	A := 20
	B := 10

	fmt.Println("A + B =", A+B) // Addition --> 30
	fmt.Println("A - B =", A-B) // Subtraction --> 10
	fmt.Println("A * B =", A*B) // Multiplication --> 200
	fmt.Println("A / B =", A/B) // Division --> 2
	fmt.Println("A % B =", A%B) // Modulus or Remainder --> 0

	A++
	fmt.Println("A++ =", A)

	B--
	fmt.Println("B-- =", B)

	// Relational Operators --> Returns true or false
	// ==, !=, >, <, >=, <=
	fmt.Println("A == B =", A == B) // Equals --> false
	fmt.Println("A != B =", A != B) // Not Equals --> true
	fmt.Println("A > B =", A > B)   // Greater Than --> true
	fmt.Println("A < B =", A < B)   // Less Than --> false
	fmt.Println("A >= B =", A >= B) // Greater Than or Equals --> true
	fmt.Println("A <= B =", A <= B) // Less Than or Equals --> false

	// Logical Operators
	// &&, ||, !

	age := 30
	isMarried := true

	fmt.Println("age > 18 && isMarried =", age > 18 && isMarried) // && means AND, both conditions has to be true
	fmt.Println("age > 18 || isMarried =", age > 18 || isMarried) // || means OR, at least one condition has to be true
	fmt.Println("!isMarried =", !isMarried)                       // ! means NOT, it reverses the boolean value

	// Assignment Operators
	// =, +=, -=, *=, /=, %=
	C := 10 // = assigns the value on the right to the variable on the left
	D := 20

	A += C // += adds the value on the right to the variable on the left
	B -= D // -= subtracts the value on the right from the variable on the left
	C *= D // *= multiplies the value on the right with the variable on the left
	D /= C // /= divides the value on the right by the variable on the left

	fmt.Println("A =", A) // 30
	fmt.Println("B =", B) // 10
	fmt.Println("C =", C) // 200
	fmt.Println("D =", D) // 5

	// Bitwise Operators
	// &, |, ^, <<, >>
	E := 30

	// Bitwise AND assignment
	E &= 2

	fmt.Println("E &= 2 =", E) // 2

	// Bitwise XOR Assignment
	E ^= 4

	fmt.Println("E ^= 4 =", E) // 6

	// Bitwise OR Assignment
	E |= 8

	fmt.Println("E |= 8 =", E) // 14

	// Bitwise Left Shift Assignment
	E <<= 2

	fmt.Println("E <<= 2 =", E) // 56

	// Bitwise Right Shift Assignment
	E >>= 1

	fmt.Println("E >>= 1 =", E) // 28

	// Miscellaneous Operators
	F := 10

	//  Address-of operator
	ptr := &F

	fmt.Println("ptr =", ptr) // 0xc00000a0c0

	//   Pointer dereference operator
	fmt.Println("*ptr =", *ptr) // 10

}
