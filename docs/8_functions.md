# Go Functions — Beginner-Friendly Guide

Functions are **reusable blocks of code** that perform a specific task.  
They help **organize code**, make it **modular**, and allow **code reuse**.

## 1. General Form of a Function

```go
func functionName(parameters) returnType {
    // body of the function
}
```

- `func` → keyword to define a function
- `functionName` → name of the function
- `parameters` → input values the function accepts
- `returnType` → the type of value returned by the function (optional)
- The function **body** contains the instructions to execute

## 2. Simple Function

```go
func sayCapital(s string) {
    fmt.Printf("%s is the Capital\n", s) // Output depends on argument
}

func sayCountry(s string) {
    fmt.Printf("%s is the Country\n", s) // Output depends on argument
}

// Usage
sayCountry("Bangladesh") // Output: Bangladesh is the Country
sayCapital("Dhaka")      // Output: Dhaka is the Capital
```

- Functions can take **parameters** to accept inputs.
- Functions can be called **multiple times** with different arguments.

## 3. Function with Single Return Value

```go
func finMaxNum(num1, num2 int) int {
    if num1 > num2 {
        return num1
    } else {
        return num2
    }
}

// Usage
a, b := 10, 20
result := finMaxNum(a, b)
fmt.Printf("The Max value is: %d\n", result) // Output: The Max value is: 20
```

- Returns a **single value** using `return` keyword.
- Can accept **multiple parameters** separated by commas.

## 4. Function with Multiple Return Values

```go
func swap(x, y string) (string, string){
    return y, x
}

// Usage
firstName, lastName := swap("Bran", "The Builder")
fmt.Printf("Swapped names: %s %s\n", firstName, lastName)
// Output: Swapped names: The Builder Bran
```

- Functions in Go can **return multiple values**.
- Useful for **swapping values**, **returning result and error**, etc.

## 5. Function Arguments vs Parameters

- **Parameter** → variable declared in function definition
- **Argument** → value passed to the function when calling it

```go
func saySomething(message string) { // message is parameter
    fmt.Println(message)
}

saySomething("Hello Go!") // "Hello Go!" is argument
```

## 6. Call by Value (Default)

- Go functions **pass arguments by value** by default.
- The **original variable** is **not modified** inside the function.

```go
func increment(number int){
    number++
    fmt.Println("Inside increment:", number) // 11
}

x := 10
increment(x)
fmt.Println("In main after increment:", x) // 10
```

**Explanation:**

- `number` inside `increment` is a **copy of x**.
- Changing `number` **does not affect** `x`.

## 7. Call by Reference (Using Slices or Pointers)

- Slices, maps, and pointers are **reference types** in Go.
- Passing them to a function allows **modifying the original value**.

```go
func modify(slice []int){
    slice[0] = 999
    fmt.Println("Inside modify:", slice) // [999 2 3]
}

mySlice := []int{1, 2, 3}
modify(mySlice)
fmt.Println("In main after modify:", mySlice) // [999 2 3]
```

**Explanation:**

- `slice` in `modify` **points to the same underlying array** as `mySlice`.
- Modifying `slice[0]` **changes the original slice** in `main`.

## 8. Summary

- Functions **modularize your code**.
- They can accept **parameters** and **return values**.
- Go supports **single and multiple return values**.
- Arguments are **passed by value** by default.
- Reference types (slices, maps, pointers) allow **pass by reference** behavior.
- Always use functions to **avoid repetitive code** and **increase readability**.
