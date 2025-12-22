// Data Types
//  The Data Types in Go are

// 1. int --> integers
// 2. float --> floating points
// 5. complex --> complex numbers
// 3. string --> string
// 4. bool --> booleans

package main

import "fmt"

func main() {
	// Section 1 : Integer

	// Section 1-A: Signed integers (can be positive or negative)
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	i = 24
	i8 = 127
	i16 = -32768
	i32 = -2147483648
	i64 = 9223372036854775807

	fmt.Println("Signed integers : ", i, i8, i16, i32, i64)

	//  Section 1-B: Unsigned integers (can only be positive)
	var ui uint
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	ui = 24
	ui8 = 255
	ui16 = 65535
	ui32 = 4294967295
	ui64 = 18446744073709551615

	fmt.Println("Unsigned integers : ", ui, ui8, ui16, ui32, ui64)

	// Section 2: Floating Point

	// Section 2A: Float32
	// This is used for less precise calculations
	var f32 float32 = 10.6

	// Section 2B: Float64
	// This is used for more precise calculations
	var f64 float64 = 10.6

	fmt.Println("Float32 : ", f32)
	fmt.Println("Float64 : ", f64)

	// Demonstrating the diff in precision between the f32 and f64
	var HP float64 = 1012345678901345
	var LP float32 = 1012345678901345

	fmt.Println("HP : ", HP)
	fmt.Println("LP : ", LP)


	// Section 3: Boolean Data Type
	var isActive bool = true;
	var isOn bool = false;

	fmt.Println("IsActive", isActive)
	fmt.Println("Is On", isOn)


	//  Section 4: Complex Data Type
	var CN1 complex128 = complex(5, 10)
	var CN2 complex64 = complex(2, 7)

	fmt.Println("CN1 : ", CN1)
	fmt.Println("CN2 : ", CN2)

	// Section 5: String Data Type
	var str string = "Hello World"
	fmt.Println("String : ", str)

}
