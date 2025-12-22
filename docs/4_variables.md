# Variables

## What is a Variable in Go?

A **variable** is a named container used to store a value. Go is a **statically typed language**, which means:

* Every variable has a **type** (string, int, float, etc.)
* The type must be **declared** or **inferred** before use
* Variables cannot be **redeclared** in the same scope
* All variables must be **used**, otherwise Go throws an error

### Declaring Variables

**1. Declare and assign separately**

```go
var name1 string
name1 = "John Snow"

fmt.Println(name1)
```

**2. Declare and assign on the same line**

```go
var name2 string = "Arya Stark"

fmt.Println(name2)
```

**3. Short declaration operator (`:=`) with type inference**

```go
name3 := "Bran Stark"  // type inferred as string
house := "Winterfell"

fmt.Println(name3, house)
```

### Declaring Multiple Variables

**Same type on one line**

```go
var name4, name5 string = "Daenerys Targaryen", "Melisandre"

fmt.Println(name4, name5)
```

**Short declaration multiple variables**

```go
city, country := "King's Landing", "Westeros"

fmt.Println("The king lives in", city, "of", country)
```

**Numeric variables**

```go
var num1, num2 int8 = 10, 20

fmt.Println(num1 + num2)  // 30
```

### Mixed Type Variables with Type Inference

```go
var a, b, c = 5, 8, "Hello"
```

### Notes & Rules

* Go does **not allow redeclaration** of variables in the same scope.
* You **cannot change the type** of a variable after declaration.
* Variables mu
