# Go Control Flow — (if else and loop)

Control flow allows you to **control the order of execution** in a program.  
Go supports **conditional statements** (`if-else`) and **loops** (`for`, nested loops, and infinite loops).

## 1. Conditional Statements (If-Else)

Used to **execute a block of code only if a condition is true**.

### If Statement

```go
age := 21

if age >= 18 {
    fmt.Println("You are eligible to vote") // Output: You are eligible to vote
} else {
    fmt.Println("You are not eligible to vote")
}
```

- `if` checks a condition.
- `else` executes if the condition is false.
- Curly braces `{}` define the **code block** for each condition.

### If-Else If Ladder

Used for **multiple conditions**.

```go
score := 82

if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B+") // Output: Grade: B+
} else if score >= 70 {
    fmt.Println("Grade: B")
} else if score >= 60 {
    fmt.Println("Grade: C+")
} else if score >= 50 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: F")
}
```

- Conditions are **evaluated in order**.
- The first **true condition** executes.
- Use `else if` for additional checks.
- `else` handles all **remaining cases**.

## 2. Loops

Loops allow you to **repeat a block of code** multiple times.

### 2A. Basic For Loop

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
    // Output: 0 1 2 3 4 5 6 7 8 9 (each on a new line)
}
```

- `i := 0` → start
- `i < 10` → condition
- `i++` → increment after each iteration

### 2B. Nested Loops

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("%d x %d = %d\t", i, j, i*j)
    }
    fmt.Println()
}

// Output:
// 1 x 1 = 1  1 x 2 = 2  1 x 3 = 3
// 2 x 1 = 2  2 x 2 = 4  2 x 3 = 6
// 3 x 1 = 3  3 x 2 = 6  3 x 3 = 9
```

- Inner loop executes **completely for each iteration** of the outer loop.
- Useful for **tables, grids, or matrix operations**.

### 2C. Loop Control Statements

#### Break

Stops the loop **immediately**.

```go
for i := 0; i < 10; i++ {
    fmt.Println("Hi Mr.", i)
    if i == 5 {
        break
    }
}

// Output: Hi Mr. 0 ... Hi Mr. 5
```

#### Continue

Skips the **current iteration** and continues to the next.

```go
for i := 0; i < 10; i++ {
    if i % 2 == 0 {
        continue
    }
    fmt.Println("Hi Mr.", i)
}

// Output: Hi Mr. 1, 3, 5, 7, 9
```

#### Goto

Jumps to a **labeled statement**.

```go
for i := 0; i < 10; i++ {
    fmt.Println("Hi Mr.", i)
    if i == 5 {
        goto end
    }
}
end:
fmt.Println("Hi Mr. Endrewvaski!!")
// Output: Hi Mr. 0 ... Hi Mr. 5
//         Hi Mr. Endrewvaski!!
```

- `goto` is **rarely used** in modern Go.
- Can create **jump points** in code.

### 2D. Infinite Loop

```go
for {
    fmt.Println("Hello World")
    break // Uncomment this to stop, else it runs forever
}

// Or using goto
start:
fmt.Println("This is an infinite loop!")
goto start
```

- Loops without a **termination condition** run indefinitely.
- Useful for **servers or event loops**.
- Always ensure a way to **exit** when needed.

### Notes

- Use **if-else** for **conditional execution**.
- Use **for loops** for **repetition**.
- Nested loops allow **complex patterns**.
- `break`, `continue`, and `goto` help **control loop execution**.
- Infinite loops are common in **concurrent or server programs**.
