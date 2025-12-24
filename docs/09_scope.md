# Scope

In Go, **scope** determines where a variable is **visible and accessible**.  
There are mainly **local** and **global** scopes, plus **function parameters**.


## 1. Local Variables

- **Declared inside a function or block**  
- Only accessible **within that function or block**  
- They are **created when the function starts** and **destroyed when it ends**

```go
func doSum() {
    var a, b, c int
    a = 10
    b = 20
    c = a + b
    fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c) 
    // Output: a = 10, b = 20, c = 30
}

doSum()
```

**Explanation:**  

- `a`, `b`, `c` exist **only inside `doSum()`**.  
- They **cannot be accessed outside** this function.


## 2. Global Variables

- **Declared outside any function or block**  
- Accessible **from any function** in the same package

```go
var g int  // global variable

func main() {
    a := 10
    b := 20
    g = a + b
    fmt.Printf("a = %d, b = %d, global variable g = %d\n", a, b, g)
    // Output: a = 10, b = 20, global variable g = 30
}
```

**Explanation:**  

- `g` is accessible in `main()` and can also be accessed from other functions in the same package.  
- Global variables **persist for the lifetime of the program**.


## 3. Function Parameters Scope

- Parameters are **local variables** defined in the function signature  
- Accessible **only inside the function**

```go
func add(x, y int) int {
    return x + y
}

sum := add(5, 7)
fmt.Println("Sum is", sum) // Output: Sum is 12
```

**Explanation:**  

- `x` and `y` exist **only inside `add()`**  
- Outside the function, they **do not exist**.


## 4. Variable Shadowing

- Happens when a **local variable has the same name** as a **global variable**  
- The **local variable "shadows" the global one** within its scope

```go
var g int = 100 // global variable

func main() {
    g := 50 // local variable shadows global
    fmt.Println("Local g:", g)   // Output: Local g: 50
    fmt.Println("Global g:", g)  // Output: Local g: 50 (global g is shadowed here)
}
```

**Explanation:**  

- Inside `main()`, the **local `g` is used**  
- The global `g` is **temporarily hidden** in this scope  


## Summary

* **Local variables** → accessible only inside their function/block  
* **Global variables** → accessible anywhere in the package  
* **Function parameters** → behave like local variables  
* **Shadowing** → local variable with same name hides outer/global variable  

> ⚠️ Tip: Use **local variables whenever possible** to avoid accidental modification of global variables.  
> Global variables are helpful for **shared data across functions**, but overuse can make code **hard to maintain**.
