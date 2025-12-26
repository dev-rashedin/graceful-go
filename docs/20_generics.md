# Go Generics

Generics were introduced in **Go 1.18** to allow writing **type-safe**, **reusable**, and **clean** code **without duplication**.

This guide explains generics **from scratch**, using **simple English**, clear rules, and practical examples.


## 1. What Are Generics?

**Generics** allow you to write **functions and data structures** that work with **multiple data types** while still being **type-safe**.

Before generics:
- You had to write the same function multiple times
- Or use `interface{}` (unsafe and error-prone)

With generics:
- One function
- Multiple types
- Compile-time safety ✅


## 2. Why Go Needed Generics

### Without Generics ❌

```go
func sumInt(a, b int) int {
    return a + b
}

func sumFloat(a, b float64) float64 {
    return a + b
}
```

Problems:
- Code duplication
- Hard to maintain
- Not scalable


## 3. Basic Generic Function

### Syntax Overview

```go
func FunctionName[T any](param T) T {
    return param
}
```

Explanation:
- `T` → type parameter (placeholder for a type)
- `any` → allows any type (alias for `interface{}`)
- `T` is decided when the function is called


## 4. Your First Generic Function

```go
func printValue[T any](value T) {
    fmt.Println(value)
}
```

Usage:

```go
printValue        // Output: 10
printValue[string]("Go")  // Output: Go
printValue[float64](3.14) // Output: 3.14
```

✔ One function  
✔ Multiple types  
✔ Fully type-safe  


## 5. Type Inference (Go Is Smart)

You can skip the type most of the time:

```go
printValue(100)
printValue("Hello")
```

Go **infers** the type automatically.


## 6. Generics with Constraints

Sometimes you **don’t want any type** — only specific ones.

Example: Sum numbers only.

### Using Constraints

```go
func sumNumbers[T int | float64](a, b T) T {
    return a + b
}
```

Usage:

```go
fmt.Println(sumNumbers(10, 20))       // 30
fmt.Println(sumNumbers(3.5, 2.5))     // 6.0
```

❌ This will fail:
```go
sumNumbers("a", "b") // compile-time error
```

---

## 7. Creating Custom Constraints

```go
type Number interface {
    int | int32 | int64 | float32 | float64
}
```

Use it:

```go
func multiply[T Number](a, b T) T {
    return a * b
}
```

Benefits:
- Clean
- Reusable
- Readable


## 8. Generic Structs

Generics are not only for functions.

```go
type Box[T any] struct {
    value T
}
```

Usage:

```go
intBox := Box[int]{value: 10}
strBox := Box[string]{value: "Go"}

fmt.Println(intBox.value) // 10
fmt.Println(strBox.value) // Go
```


## 9. Generics vs interface{}

| Generics | interface{} |
|--------|-------------|
| Type-safe | Not type-safe |
| Compile-time checks | Runtime errors |
| Better performance | Slower |
| Clean APIs | Messy type assertions |


## 10. When to Use Generics

Use generics when:
- Same logic for multiple types
- Data structures (Stack, Queue, Map)
- Utility functions
- Libraries and frameworks

Avoid generics when:
- Logic is simple
- Only one type is used
- Code becomes harder to read


## 11. Real-World Example: Generic Max Function

```go
type Ordered interface {
    int | float64 | string
}

func max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

Usage:

```go
fmt.Println(max(10, 20))       // 20
fmt.Println(max(3.2, 1.5))    // 3.2
fmt.Println(max("a", "b"))    // b
```

## 12. Key Rules to Remember

- Generics use **type parameters**
- Constraints limit allowed types
- Compile-time safety is guaranteed
- Prefer **simple generics**
- Readability comes first


## 13. Mental Model (Very Important)

Think of generics as:

> “Write the logic once,  
> let Go fill in the type later.”


### Final Takeaway

Generics make Go:
- Cleaner
- Safer
- More powerful

But Go still values:
- Simplicity
- Readability
- Explicitness

Use generics **wisely**, not everywhere 🚀
