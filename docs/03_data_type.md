# Data Types

A **data type** tells Go **what kind of value** a variable can store.
Go is a **statically typed language**, which means:
- Every variable has a fixed type
- The type is known before the program runs
- Go will throw an error if types don’t match


## MAIN DATA TYPES IN GO

1. int        → whole numbers (integers)
2. float     → decimal numbers
3. complex   → complex numbers (real + imaginary)
4. string    → text
5. bool      → true / false


## SECTION 1: INTEGER DATA TYPES


Integers are **whole numbers** (no decimals).

Go has two kinds of integers:
- Signed integers (positive + negative)
- Unsigned integers (only positive)


### 1A. SIGNED INTEGERS

Signed integers can store:
- positive numbers
- negative numbers
- zero

Types:
int, int8, int16, int32, int64

Example:

```go
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

fmt.Println("Signed integers:", i, i8, i16, i32, i64)
```


### WHAT DOES “SIGNED” MEAN?


“Signed” means the number can be **positive or negative**.

In memory:
- one bit is used to store the sign
- remaining bits store the value


### SIGNED INTEGER RANGES (IMPORTANT)

int8   →  -128 to 127  
int16  →  -32768 to 32767  
int32  →  -2147483648 to 2147483647  
int64  →  -9223372036854775808 to 9223372036854775807  

Why this matters:
- Go will NOT allow numbers outside the range
- You must choose the correct size


### 1B. UNSIGNED INTEGERS


Unsigned integers can store:
- only positive numbers
- no negative values

Because no bit is used for sign,
they can store **larger positive numbers**.

Types:
uint, uint8, uint16, uint32, uint64

Example:

```go
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

fmt.Println("Unsigned integers:", ui, ui8, ui16, ui32, ui64)
```

Why unsigned exists:
- memory efficiency
- useful for counts, IDs, sizes
- values that can never be negative


## SECTION 2: FLOATING POINT NUMBERS


Floating-point numbers store **decimal values**.

Go has two float types:
- float32 → less precision
- float64 → more precision (recommended)


### FLOAT32

Used when:
- memory is important
- precision is less critical

Example:

```go
var f32 float32 = 10.6

fmt.Println("Float32:", f32)
```


### FLOAT64

Used when:
- accuracy is important
- default choice in Go

Example:

```go
var f64 float64 = 10.6

fmt.Println("Float64:", f64)
```

### PRECISION DIFFERENCE (VERY IMPORTANT)

```go
var HP float64 = 1012345678901345
var LP float32 = 1012345678901345

fmt.Println("HP:", HP) // This prints --> HP :  1.012345678901345e+15
fmt.Println("LP:", LP) // This prints --> LP :  1.0123457e+15
```

Explanation:
- float64 keeps more digits
- float32 loses precision
- Always use float64 unless you have a reason not to


### SECTION 3: BOOLEAN DATA TYPE


Boolean stores only:
- true
- false

Used in:
- conditions
- logic
- decision making

Example:

```go
var isActive bool = true
var isOn bool = false

fmt.Println("IsActive:", isActive)
fmt.Println("IsOn:", isOn)
```

### SECTION 4: COMPLEX DATA TYPE

Complex numbers have:
- real part
- imaginary part

Format:
complex(real, imaginary)

Types:
- complex64
- complex128

Example:

```go
var CN1 complex128 = complex(5, 10)
var CN2 complex64 = complex(2, 7)

fmt.Println("CN1:", CN1)
fmt.Println("CN2:", CN2)
```

Used in:
- scientific computing
- signal processing
- advanced math


## SECTION 5: STRING DATA TYPE

String stores **text data**.

Properties:
- enclosed in double quotes
- immutable (cannot be changed)

Example:

```go
var str string = "Hello World"

fmt.Println("String:", str)
```

### ONE-LINE SUMMARY

Go data types are strict, predictable, and memory-efficient — perfect for reliable systems programming
