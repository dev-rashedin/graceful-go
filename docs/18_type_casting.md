# Go Type Casting

**Type casting** (also called **type conversion**) means converting a value from **one data type to another**.

Go is a **statically typed language**, which means:

- Every variable has a **fixed type**
- Go does **not** allow automatic type conversion
- You must convert types **explicitly**

## 1. Why Type Casting is Needed in Go

In Go:

- You **cannot** mix different types in calculations
- Even similar types like `int` and `float64` are **not compatible**

❌ This will cause an error:

```go
var a int = 10
var b float64 = 3.5
fmt.Println(a + b) // ERROR
```

✅ Correct way:

```go
fmt.Println(float64(a) + b)
```

## 2. General Syntax of Type Casting

```go
newValue := TargetType(originalValue)
```

Example:

```go
var x int = 10
var y float64 = float64(x)
```

## 3. Converting int to float (Numeric Type Casting)

### Example: Calculating Average

```go
var total int = 20
var items int = 7

var average float64
average = float64(total) / float64(items)

fmt.Printf("Average is: %.2f\n", average) // Output: 2.86
```

### Explanation:

- `total` and `items` are `int`
- Division of `int / int` gives an `int`
- Converting to `float64` preserves decimal precision

## 4. Converting float to int

```go
var price float64 = 19.99
var quantity int = int(price)

fmt.Printf("Quantity: %d\n", quantity) // Output: 19
```

### Important Note:

- Go **does not round**
- It **truncates** (cuts off the decimal part)

```
19.99 → 19
```

## 5. Converting string to int (Using strconv Package)

Strings **cannot** be directly cast to numbers.
Go provides the **strconv** package for this.

### Import Required Package

```go
import "strconv"
```

### Example: String to Integer Conversion

```go
str := "123"

num, err := strconv.Atoi(str)

if err != nil {
    fmt.Println("Error converting:", err)
}

if err == nil {
    fmt.Printf("Converted number: %d\n", num) // Output: 123
}
```

### Explanation:

- `strconv.Atoi` converts **string → int**
- It returns **two values**:
  - The converted number
  - An error (`nil` if conversion succeeds)

## 6. What Happens on Invalid Conversion?

```go
str := "abc"
num, err := strconv.Atoi(str)
```

Output:

```
error converting: strconv.Atoi: parsing "abc": invalid syntax
```

✔ Always check the error before using the value.

## 7. Complete Working Example

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println("- Type Casting -")

    // int → float
    var total int = 20
    var items int = 7
    average := float64(total) / float64(items)
    fmt.Printf("Average is: %.2f\n", average) // 2.86

    // float → int
    var price float64 = 19.99
    var quantity int = int(price)
    fmt.Printf("Quantity: %d\n", quantity) // 19

    // string → int
    str := "123"
    num, err := strconv.Atoi(str)

    if err != nil {
        fmt.Println("Error converting:", err)
    } else {
        fmt.Printf("Converted number: %d\n", num) // 123
    }
}
```

## 8. Common Type Casting Functions in Go

| Conversion     | Function                |
| -------------- | ----------------------- |
| string → int   | `strconv.Atoi()`        |
| int → string   | `strconv.Itoa()`        |
| string → float | `strconv.ParseFloat()`  |
| float → string | `strconv.FormatFloat()` |

## 9. Key Rules to Remember

- ❌ Go does **not** allow implicit type conversion
- ✅ Always convert types **explicitly**
- ⚠️ Float → int **truncates**, does not round
- ✔ Always handle errors when converting strings

### Summary

- Type casting is mandatory in Go
- Keeps code **safe**, **predictable**, and **clear**
- Essential for calculations, input handling, and data processing

> Once you master type casting, handling real-world data in Go becomes much easier 🚀
