# Range in Go

The `range` keyword in Go is used to iterate over elements in **arrays, slices, maps, strings, and channels**. It is similar to `enumerate` in Python.


## Example: Using Range with Slices and Arrays

```go
package main

import "fmt"

func main() {
    fmt.Println("Range")

    // Slice example
    numbers := []int{0, 1, 2, 3, 4, 5, 6, 7} 

    // Using range returns both index and value
    for i, number := range numbers {
        fmt.Printf("Slice item %d is %d\n", i, number)
    }
    // Output:
    // Slice item 0 is 0
    // Slice item 1 is 1
    // Slice item 2 is 2
    // ...
    
    // Array example
    n := []int{1, 2, 3}

    // Using range with index only
    for i := range n {
        fmt.Printf("Array item %d is %d\n", i, n[i])
    }
    // Output:
    // Array item 0 is 1
    // Array item 1 is 2
    // Array item 2 is 3
}
```


## How `range` works

1. **Slice or Array:**

   * Returns **index** and **value** at that index.
   * You can ignore the value by using `_`.
   * You can ignore the index if not needed.

```go
for _, value := range numbers {
    fmt.Println(value) // Only prints the values
}

for i := range numbers {
    fmt.Println(i) // Only prints the indices
}
```

2. **Map:**

   * Returns **key** and **value**.

```go
m := map[string]int{"a": 1, "b": 2}
for k, v := range m {
    fmt.Println(k, v)
}
```

3. **String:**

   * Returns **index** and **rune** at that index.

```go
s := "hello"
for i, r := range s {
    fmt.Println(i, r) // r is unicode code point
}
```

4. **Channel:**

   * Iterates over received values until channel is closed.

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)

for v := range ch {
    fmt.Println(v)
}
```


### Key Points

* `range` **always copies the value** for each iteration.
* For slices, modifying the value variable does **not change the original slice element**.
* Use `_` to ignore either the index or value.
* Very similar to Python’s `enumerate`.


### Summary

`range` is a **versatile and safe way** to iterate over different Go data structures with concise syntax. It works for **arrays, sli
