# Operators in Go

An **operator** is a symbol that tells Go to perform a **mathematical, logical, or technical operation**
on one or more values (called **operands**).

Go operators are grouped based on what they do.

## 1. Arithmetic Operators

Arithmetic operators are used for **basic math calculations**.

- `+` → Addition
- `-` → Subtraction
- `*` → Multiplication
- `/` → Division
- `%` → Modulus (remainder)
- `++` → Increment
- `--` → Decrement

Example:

```go

    A := 20
    B := 10

    A + B // 30 → adds A and B
    A - B // 10 → subtracts B from A
    A * B // 200 → multiplies A and B
    A / B // 2 → divides A by B
    A % B // 0 → remainder after division

    A++ // increases A by 1
    B-- // decreases B by 1
```

## 2. Relational (Comparison) Operators

Relational operators **compare two values** and always return a **boolean** (`true` or `false`).

- (==) Equal to
- (!=) Not equal to
- (>) Greater than
- (<) Less than
- (>) = Greater than or equal to
- (<=) Less than or equal to

Example:

```go

A := 30;
B := 20;

 A == B // false → checks equality
 A != B // true → checks inequality
 A > B  // true → checks if A is greater
 A < B  // false → checks if A is smaller
 A >= B // true → checks greater than or equals
 A <= B // false → less than or equals
```

These are heavily used in:

- if conditions
- loops
- decision making

## 3. Logical Operators

Logical operators work with **boolean values**.

- `&&` → AND → both conditions must be true
- `||` → OR → at least one condition must be true
- `!` → NOT → reverses a boolean value

Example:

```go
 age := 30
 isMarried := true

 age > 18 && isMarried // true
 age > 18 || isMarried // false
 !isMarried           // false
```

## 4. Assignment Operators

Assignment operators are used to **assign or update values** in variables.

- `=` → Assign
- `+=` → Add and assign
- `-=` → Subtract and assign
- `*=` → Multiply and assign
- `/=` → Divide and assign

Example:

```go
 A := 10
 B := 20

 A += 5  // A = A + 5 → 15
 B -= 10 // B = B - 10 → 10
 A *= 2  // A = A * 2 → 20
 B /= 2  // B = B / 2 → 2
```

## 5. Bitwise Operators

Bitwise operators work at the **binary (bit) level**.
They are mainly used in **systems programming and performance-critical code**.

- `&` → AND
- `|` → OR
- `^` → XOR
- `<<` → Left shift
- `>>` → Right shift

### Bitwise Operators Example with Explanations

```go
A := 30        // 30 in binary: 00011110

// Bitwise AND: only bits that are 1 in BOTH operands become 1
A &= 2         // 2 in binary: 00000010
fmt.Println("A &= 2 =", A) // Result: 2 → 00000010

// Bitwise XOR: bits are 1 if they are different, 0 if they are the same
A ^= 4         // 4 in binary: 00000100
fmt.Println("A ^= 4 =", A) // Result: 6 → 00000110

// Bitwise OR: bits are 1 if at least one operand has 1
A |= 8         // 8 in binary: 00001000
fmt.Println("A |= 8 =", A) // Result: 14 → 00001110

// Left Shift: shifts bits to the left, filling with 0s on the right (multiplies by 2^n)
A <<= 2
fmt.Println("A <<= 2 =", A) // Result: 56 → 00111000

// Right Shift: shifts bits to the right, discards bits on the right (divides by 2^n)
A >>= 1
fmt.Println("A >>= 1 =", A) // Result: 28 → 00011100
```

### Notes

- `&` → use when you need **common set bits**.
- `|` → use when you want **union of bits**.
- `^` → use for **toggling or masking bits**.
- `<<` and `>>` → use for **fast multiplication/division by powers of 2**.
- Bitwise operators **only work on integers**. Floating-point numbers cannot use them.

## 6. Miscellaneous Operators

Go provides **special operators** for working with **pointers** and memory.  
These are less commonly used in everyday programming but are essential for **systems programming** and **memory-efficient code**.

### Address-of Operator (`&`)

- Returns the **memory address** of a variable.
- Useful when you want to work with **pointers**.

```go
F := 10           // Declare an integer variable
ptr := &F         // Get the memory address of F and store it in ptr

fmt.Println("Memory address of F:", ptr) // Example output: 0xc0000180a8
```

_Notes:_

- `&` can only be used with **variables**, not constants or literals.
- The result is a **pointer** that stores the address of the variable.

### Dereference Operator (`*`)

- Returns the **value stored at a pointer's memory address**.
- Allows you to **access or modify the value** of the original variable through its pointer.

```go
fmt.Println("Value stored at ptr:", *ptr) // Output: 10

*ptr = 20                              // Modify the value of F through the pointer
fmt.Println("Updated value of F:", F)  // Output: 20
```

_Notes:_

- `*` is used to **dereference a pointer**.
- Combined with `&`, you can **read or change a variable via its memory address**.
- This is the foundation for pointers in Go.

### Summary

- `&` → get the **memory address** of a variable
- `*` → get or **modify the value** at a memory address
- Together, they allow **pointer manipulation** and **direct memory access**
- Mostly used in **low-level programming** or when **optimizing memory usage**
