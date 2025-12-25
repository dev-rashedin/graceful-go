# Go Interfaces — Beginner-Friendly, In-Depth Guide

Interfaces are one of the **most powerful features of Go**.  
They help you write **flexible**, **reusable**, and **clean** code without inheritance.

This guide explains interfaces **from zero**, using **simple English** and your example code.

## 1. What is an Interface in Go?

An **interface** is a **set of method signatures**.

👉 It defines **what a type can do**, not **what it is**.

```go
type Shape interface {
    area() float64
}
```

This means:

- Any type that has an `area() float64` method
- Automatically satisfies (`implements`) the `Shape` interface

⚠️ Important:

> Go has **implicit interfaces**  
> You do **NOT** write `implements` like Java or TypeScript.

## 2. Why Interfaces Exist

Interfaces allow you to:

- Write **generic code**
- Accept **multiple types** with the same behavior
- Avoid tight coupling
- Follow **composition over inheritance**

Example idea:

> “Anything that can calculate area is a Shape.”

## 3. Interface vs Struct (Very Important)

| Struct | Interface |
||--|
| Holds **data (fields)** | Holds **method signatures** |
| Defines **what something is** | Defines **what something can do** |
| Has memory | No memory |
| Example: `Circle`, `Rectangle` | Example: `Shape` |

## 4. Defining an Interface

```go
type Shape interface {
    area() float64
}
```

This interface says:

- Any type must have:
  - A method named `area`
  - Returning `float64`

## 5. Creating Structs That Implement the Interface

### Circle Struct

```go
type Circle struct {
    radius float64
}
```

### Rectangle Struct

```go
type Rectangle struct {
    width, height float64
}
```

At this point:

- ❌ They do NOT yet implement `Shape`

## 6. Implementing Interface Methods

### Circle Implements `area()`

```go
func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
```

### Rectangle Implements `area()`

```go
func (r Rectangle) area() float64 {
    return r.width * r.height
}
```

✅ Now BOTH:

- `Circle`
- `Rectangle`

Automatically implement `Shape`

No keyword. No declaration.  
Just **method matching**.

## 7. Interface as Function Parameter (Most Important Use)

```go
func getArea(s Shape) float64 {
    return s.area()
}
```

This function:

- Accepts **ANY type**
- As long as it implements `area() float64`

## 8. Using Interface with Multiple Types

```go
circle := Circle{radius: 5}
rectangle := Rectangle{width: 10, height: 20}

fmt.Printf("Circle area: %f\n", getArea(circle))     // Works
fmt.Printf("Rectangle area: %f\n", getArea(rectangle)) // Works
```

### Output

```text
Circle area: 78.539816
Rectangle area: 200.000000
```

✔ Same function  
✔ Different types  
✔ Same behavior

This is **polymorphism in Go**.

## 9. Full Working Example

```go
package main

import (
    "fmt"
    "math"
)

// Interface
type Shape interface {
    area() float64
}

// Structs
type Circle struct {
    radius float64
}

type Rectangle struct {
    width, height float64
}

// Methods
func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

// Function using interface
func getArea(s Shape) float64 {
    return s.area()
}

func main() {
    circle := Circle{radius: 5}
    rectangle := Rectangle{width: 10, height: 20}

    fmt.Printf("Circle area: %f\n", getArea(circle))
    fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
```

## 10. Empty Interface (`interface{}`)

```go
var x interface{}
```

This can hold:

- int
- string
- struct
- anything

But:
⚠️ Avoid overusing it  
⚠️ Loses type safety

Used mainly for:

- Generic data
- JSON decoding
- Reflection

## 11. Key Rules to Remember

- Interfaces define **behavior**, not data
- Implementation is **implicit**
- A type can implement **multiple interfaces**
- Interfaces make Go **clean and scalable**
- Small interfaces are better than large ones

## 12. Why Interfaces Are So Powerful in Go

- No inheritance hell
- No forced hierarchy
- Easy testing (mocking)
- Encourages clean architecture

> “Accept interfaces, return structs” — Go proverb

### Final Takeaway

If you understand interfaces,  
you understand **real Go programming**.

They are the backbone of:

- Standard library
- Clean architecture
- Production-grade Go code 🚀
