# Go Constants — Beginner-Friendly Guide

### What is a Constant?

A **constant** is a value that **cannot be changed** once defined.

- Immutable
- Value known at **compile time**
- Helps prevent accidental modification and improves **readability**

### Declaring Constants

#### Single constant declaration

```go
const NAME string = "John Snow"
const AGE int = 24

fmt.Println("Name is", NAME) // Output: Name is John Snow
fmt.Println("Age is", AGE)   // Output: Age is 24
```

- `const` keyword is used
- Must assign a value during declaration
- Type can be explicit (`string`, `int`) or inferred

#### Grouped constant declaration

```go
const (
    DECIMAL     = 255      // decimal
    OCTAL       = 0o377    // octal
    HEXADECIMAL = 0xff     // hexadecimal
)

fmt.Println("Decimal:", DECIMAL)         // Output: Decimal: 255
fmt.Println("Octal:", OCTAL)             // Output: Octal: 255
fmt.Println("Hexadecimal:", HEXADECIMAL) // Output: Hexadecimal: 255
```

- Use parentheses to declare **multiple constants together**
- Improves **organization** and **readability**

### Floating-Point Constants

- Can include **integer part, decimal part, fractional part, and exponent**

```go
const PI float64 = 3.14159
const AVOGADRO float64 = 6.02214076e23

fmt.Println("PI:", PI)             // Output: PI: 3.14159
fmt.Println("Avogadro:", AVOGADRO) // Output: Avogadro: 6.02214076e+23
```

- Useful for **precise calculations**
- `float64` is preferred for most high-precision calculations

### String Constants and Escape Sequences

- Escape sequences allow **special characters** in strings:

  - `\n` → newline
  - `\"` → double quote
  - `\\` → backslash
  - `\a` → bell/alert
  - `\b` → backspace
  - `\f` → form feed
  - `\r` → carriage return
  - `\t` → horizontal tab
  - `\v` → vertical tab

```go
const GREETING = "Hello, Earth!\n"
const QUOTE = "\"Life is what happens to you while you're busy making other plans.\" - John Lennon"
const BELL = "Bell is \a"

fmt.Println(GREETING) // Output: Hello, Earth!  (prints a new line)
fmt.Println(QUOTE)    // Output: "Life is what happens to you while you're busy making other plans." - John Lennon
fmt.Println(BELL)     // Output: Bell is  (with bell sound on supported terminals)
```

- Multi-line or concatenated string constants

```go
const MULTILINE = "John Snow is now the King in the North." +
                  " He is actually a Targaryen and not a Bastard." +
                  " Therefore, he is the true heir of the Iron Throne."

fmt.Println(MULTILINE)
// Output: John Snow is now the King in the North. He is actually a Targaryen and not a Bastard. Therefore, he is the true heir of the Iron Throne.
```

### Boolean Constants

```go
const ACTIVE = true
const READY = false

fmt.Println(ACTIVE) // Output: true
fmt.Println(READY)  // Output: false
```

- Represents **true/false values**
- Often used for **flags** or **conditional checks**

### Constants for Calculations

```go
const LENGTH = 50
const WIDTH  = 25
const AREA   = LENGTH * WIDTH

fmt.Println("The area of the rectangle is:", AREA)
// Output: The area of the rectangle is: 1250
```

- Constants can be used in **calculations**
- Value **cannot change**, making code safer

### Notes

- Constants **cannot be reassigned**.
- Use constants for values that **do not change** during program execution.
- Improves **readability** and **prevents accidental modification**.
- Always prefer constants for **fixed values** like PI, configuration values, or fixed strings.
