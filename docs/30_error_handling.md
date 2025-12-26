# Error Handling in Go — Complete Beginner-Friendly Guide

Error handling is a **core design philosophy** of Go.  
Instead of exceptions (`try/catch`), Go uses **explicit error values** that make failures **visible, predictable, and easy to reason about**.

This guide explains **how error handling works in Go**, **why it is designed this way**, and **how to write clean, professional error-handling code**, using your example.


## 1. What Is an Error in Go?

In Go, an error is simply a **value**.

Go defines a built-in interface called `error`:

```go
type error interface {
    Error() string
}
```

Any type that implements the `Error() string` method is considered an error.


## 2. Go’s Error Philosophy (Very Important)

Go follows these principles:

- Errors are **returned**, not thrown
- Errors are **checked explicitly**
- A `nil` error means **success**
- A non-`nil` error means **failure**

This forces developers to **handle errors consciously**, instead of accidentally ignoring them.


## 3. Returning Errors from Functions

Most Go functions return **two values**:

```go
result, err := someFunction()
```

- `result` → actual return value
- `err` → error information

### Example from Your Code

```go
func Sqrt(value float64) (float64, error) {
    if value < 0 {
        return 0, errors.New("Negative number passed to Sqrt!")
    }

    return math.Sqrt(value), nil
}
```

### Key Points

- Function returns `(float64, error)`
- On failure → return meaningful error
- On success → return result + `nil`


## 4. Handling Errors (The Go Way)

Error handling always follows this pattern:

```go
result, err := Sqrt(-5)

if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

### Why This Is Good

✔ Clear control flow  
✔ No hidden exceptions  
✔ Easy to debug  
✔ Explicit logic  


## 5. Creating Errors Using `errors.New`

The simplest way to create an error:

```go
errors.New("something went wrong")
```

Example:

```go
if value < 0 {
    return 0, errors.New("Negative number passed to Sqrt!")
}
```

Use this when:
- You only need a message
- No extra context is required


## 6. Custom Errors (Advanced & Powerful)

Go allows you to create **custom error types**.

### Example from Official Go Docs (Your Code)

```go
type MyError struct {
    When time.Time
    What string
}
```

Implement the `Error()` method:

```go
func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}
```

Now `MyError` **implements the error interface automatically**.


## 7. Returning Custom Errors

```go
func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}
```

Handling it:

```go
if err := run(); err != nil {
    fmt.Println(err)
}
```

### Output Example

```
at 2025-01-01 12:00:00, it didn't work
```

✔ Rich error messages  
✔ Extra context  
✔ Professional-grade error handling  


## 8. Why Errors Are Interfaces in Go

Because errors are interfaces:

- You can return **different error types**
- You can compare, inspect, or wrap them
- You can build powerful APIs

This gives Go **flexibility without complexity**.


## 9. Common Error Handling Patterns

### Pattern 1: Early Return (Recommended)

```go
if err != nil {
    return err
}
```

Keeps code:
- Flat
- Readable
- Maintainable


### Pattern 2: Ignoring Errors (❌ Bad Practice)

```go
result, _ := Sqrt(10)
```

Only acceptable when:
- Error truly does not matter
- You are 100% sure


## 10. Error Naming Convention

- Error messages start with **lowercase**
- No punctuation at the end

✔ `"file not found"`  
❌ `"File not found!"`

Reason: Errors may be **chained together**.


## 11. Error Handling vs Exceptions

| Go Errors | Exceptions |
|---------|------------|
| Explicit | Hidden |
| Compile-time visible | Runtime surprises |
| Easy to debug | Hard to trace |
| Simple control flow | Complex stack unwinding |

Go chooses **clarity over magic**.


## 12. Mental Model (Very Important)

Think of Go errors as:

> “Errors are normal return values —  
> not special events.”

Once you accept this, Go error handling becomes **simple and elegant**.


## 13. Final Best Practices

✔ Always check errors  
✔ Return errors early  
✔ Create meaningful error messages  
✔ Use custom errors for context  
✔ Keep error handling readable  


## Final Takeaway

Go does not hide failures.

It **forces you to deal with reality**, which makes Go programs:

- Reliable
- Predictable
- Production-ready

Once you master Go error handling, you’ll **never want exceptions again** 🚀
