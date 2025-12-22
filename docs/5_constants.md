# Go Constants — Beginner-Friendly Guide

### What is a Constant?

A **constant** is a value that **cannot be changed** once defined. Constants are **immutable** and their value is known at **compile time**.

### Declaring Constants

* **Single constant declaration**

```go
const NAME string = "John Snow"
const AGE int = 24
fmt.Println("Name is", NAME)
fmt.Println("Age is", AGE)
```

* **Grouped constant declaration**

```go
const (
    DECIMAL     = 255      // decimal
    OCTAL       = 0o377    // octal
    HEXADECIMAL = 0xff     // hexadecimal
)
fmt.Println("Decimal:", DECIMAL, "Octal:", OCTAL, "Hexadecimal:", HEXADECIMAL)
```

### Floating-Point Constants

* Can include **integer part, decimal part, fractional part, and exponent**

```go
const PI float64 = 3.14159
const AVOGADRO float64 = 6.02214076e23
fmt.Println("PI:", PI, "Avogadro:", AVOGADRO)
```

### String Constants and Escape Sequences

* Escape sequences allow special characters in strings:

  * `\n` → newline
  * `\"` → double quote
  * `\\` → backslash
  * `\a` → bell
  * `\b` → backspace
  * `\f` → form feed
  * `\r` → carriage return
  * `\t` → horizontal tab
  * `\v` → vertical tab

```go
const GREETING = "Hello, Earth!\n"
const QUOTE = "\"Life is what happens to you while you're busy making other plans.\" - John Lennon"
const BELL = "Bell is \a"
fmt.Println(GREETING)
fmt.Println(QUOTE)
fmt.Println(BELL)
```

* Multi-line or concatenated string constants

```go
const MULTILINE = "John Snow is now the King in the North." + 
                  " He is actually a Targaryen and not a Bastard." + 
                  " Therefore, he is the true heir of the Iron Throne."
fmt.Println(MULTILINE)
```

### Boolean Constants

```go
const ACTIVE = true
const READY = false
fmt.Println(ACTIVE, READY)
```

### Constants for Calculations

```go
const LENGTH = 50
const WIDTH  = 25
const AREA   = LENGTH * WIDTH
fmt.Println("The area of the rectangle is:", AREA)
```

### Notes

* Constants **cannot be reassigned**.
* Use constants for values that **do not change** during program execution.
* Hel
