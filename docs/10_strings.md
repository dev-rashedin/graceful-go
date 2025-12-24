# Go Strings — Beginner-Friendly Guide

Strings in Go are **sequences of characters** and are widely used for text manipulation. Go also provides the **`strings` package** with many helpful methods.

## 1. Creating Strings

```go
var greeting string = "Hello, World!"

fmt.Printf("Normal String: %s\n", greeting)
// Output: Normal String: Hello, World!
```

- Strings can be declared using **`var`** with type `string`
- Go also supports **short declaration**: `greeting := "Hello, World!"`

## 2. Inspecting Strings

### 2A. Display hexadecimal byte values

```go
for i := 0; i < len(greeting); i++ {
    fmt.Printf("%x ", greeting[i])
}
// Output: 48 65 6c 6c 6f 2c 20 57 6f 72 6c 64 21
```

**Explanation:**

- `greeting[i]` accesses the **byte value** of each character
- `%x` formats the byte in **hexadecimal**

### 2B. String Length

```go
fmt.Println(len(greeting))
// Output: 13
```

- `len()` returns the number of **bytes** in the string
- For ASCII, this is the **number of characters**

## 3. String Concatenation

### 3A. Using `+` operator

```go
first := "Hello"
second := "World"
result := first + ", " + second + "!"
fmt.Println(result)
// Output: Hello, World!
```

### 3B. Using `strings.Join` with a slice

```go
fruits := []string{"apples", "oranges", "and bananas"}
fmt.Println(strings.Join(fruits, " "))
// Output: apples oranges and bananas
```

**Explanation:**

- `strings.Join(slice, sep)` concatenates slice elements using the separator `sep`
- More efficient than repeated `+` for multiple strings

## 4. Common String Methods (from `strings` package)

```go
str := "  GoLang is Awesome!  "

// Trim spaces
fmt.Println(strings.TrimSpace(str))
// Output: GoLang is Awesome!

// Convert to lower case
fmt.Println(strings.ToLower(str))
// Output:   golang is awesome!

// Convert to upper case
fmt.Println(strings.ToUpper(str))
// Output:   GOLANG IS AWESOME!

// Check prefix
fmt.Println(strings.HasPrefix(str, "  GoLang"))
// Output: true

// Check suffix
fmt.Println(strings.HasSuffix(str, "Awesome!  "))
// Output: true

// Contains substring
fmt.Println(strings.Contains(str, "Lang"))
// Output: true

// Count occurrences
fmt.Println(strings.Count(str, " "))
// Output: 5

// Replace substring (first n occurrences)
fmt.Println(strings.Replace(str, " ", "_", 2))
// Output: __GoLang_is Awesome!
```

## 5. Notes

- Strings in Go are **immutable**; methods like `strings.Replace` return a **new string**
- For **efficient concatenation** of many strings, use **`strings.Join`** instead of repeated `+`
- `len()` counts **bytes**, not runes; for **Unicode characters**, use `utf8.RuneCountInString()`

```go
import "unicode/utf8"
fmt.Println(utf8.RuneCountInString("Go💖Lang"))
// Output: 7
```
