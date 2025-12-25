# Go Recursion — Beginner-Friendly Guide

Recursion is a powerful programming concept where a **function calls itself** to solve a problem by breaking it into **smaller sub-problems**.

In Go, recursion works the same way as in other languages, but understanding the **base case** is critical to avoid infinite loops.

## 1. What is Recursion?

- **Recursion** happens when a function calls **itself**
- Each call solves a **smaller version of the same problem**
- Recursion stops when it reaches a **base case**

> Think of recursion like climbing down a ladder step by step — once you reach the ground, you stop.

## 2. Two Mandatory Parts of Recursion

Every recursive function **must have**:

### 1️⃣ Base Case

- The condition that **stops recursion**
- Prevents infinite function calls

### 2️⃣ Recursive Case

- The part where the function **calls itself**
- Moves the problem closer to the base case

## 3. Recursive Function Structure

```go
func recursiveFunction(parameters) returnType {
    if baseCondition {
        return result // base case
    }
    return recursiveFunction(smallerProblem) // recursive case
}
```

## 4. Example 1: Factorial Using Recursion

### What is Factorial?

- `factorial(n) = n × (n-1) × (n-2) × ... × 1`
- Example: `5! = 5 × 4 × 3 × 2 × 1 = 120`

### Recursive Implementation

```go
func factorial(n int) int {
    if n == 1 {
        return 1              // Base case
    } else {
        return n * factorial(n-1) // Recursive case
    }
}
```

### How `factorial(5)` Works

```
factorial(5)
→ 5 * factorial(4)
→ 5 * 4 * factorial(3)
→ 5 * 4 * 3 * factorial(2)
→ 5 * 4 * 3 * 2 * factorial(1)
→ 5 * 4 * 3 * 2 * 1
→ 120
```

### Calling the Function

```go
fmt.Println(factorial(5)) // Output: 120
```

## 5. Example 2: Fibonacci Using Recursion

### What is Fibonacci?

- Fibonacci sequence:
  ```
  1, 1, 2, 3, 5, 8, 13, ...
  ```
- Rule:
  ```
  fib(n) = fib(n-1) + fib(n-2)
  ```

### Recursive Implementation

```go
func fibonacci(n int) int {
    if n <= 2 {
        return 1                    // Base case
    } else {
        return fibonacci(n-1) + fibonacci(n-2) // Recursive case
    }
}
```

### Calling the Function

```go
fmt.Println(fibonacci(10)) // Output: 55
```

## 6. Complete Working Example

```go
package main

import "fmt"

// Factorial using recursion
func factorial(n int) int {
    if n == 1 {
        return 1
    }
    return n * factorial(n-1)
}

// Fibonacci using recursion
func fibonacci(n int) int {
    if n <= 2 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println("Recursion Example")
    fmt.Println("Factorial of 5:", factorial(5))   // 120
    fmt.Println("Fibonacci of 10:", fibonacci(10)) // 55
}
```

## 7. Important Notes for Beginners

- ❌ Missing a base case → **infinite recursion → crash**
- ✅ Always make sure recursive calls **move toward the base case**
- ⚠️ Recursive Fibonacci is **slow** for large numbers (repeated calculations)
- 🔁 Some recursive solutions can be optimized using **loops or memoization**

## 8. When to Use Recursion?

Use recursion when:

- The problem is naturally **self-similar**
- Tree traversal, factorial, Fibonacci, DFS
- Divide-and-conquer problems

Avoid recursion when:

- Performance is critical
- A loop can solve it more efficiently

### Summary

- Recursion = function calling itself
- Requires **base case + recursive case**
- Powerful but must be used carefully
- Common in math, algorithms, and tree structures

> Master recursion once — it unlocks advanced topics like **DFS, backtracking, and dynamic programming** 🚀
