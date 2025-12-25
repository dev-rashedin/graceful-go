# Go `map` Documentation with Examples

## Maps in Go

A **map** in Go is a built-in data type used to store data in **key:value pairs**.

Maps are similar to:

* **Dictionary** in Python
* **Object** in JavaScript
* **HashMap** in Java


## Key Characteristics of Maps

* Each element is a **key:value** pair
* Keys are **unique** (no duplicates)
* Maps are **unordered**
* Maps are **mutable** (can be updated)
* Default value of a map is **nil**
* Maps store **references** to an underlying hash table
* Length of a map can be found using `len()`


## Declaring and Initializing Maps

### Declare a map (nil map)

```go
var myMap map[string]string
```

At this point:

* `myMap` is **nil**
* You **cannot insert values** yet


### Initialize a map using `make`

```go
myMap = make(map[string]string)
fmt.Println("Map value is:", myMap)
```

Now:

* Map is initialized
* You can add key:value pairs


## Declare and Initialize in One Line

```go
countryAndCapital := make(map[string]string)
```


## Adding Elements to a Map

```go
countryAndCapital["Bangladesh"] = "Dhaka"
countryAndCapital["India"] = "New Delhi"
countryAndCapital["USA"] = "Washington, D.C."
countryAndCapital["France"] = "Paris"
```


## Iterating Over a Map

### Iterate Over Keys Only

```go
for country := range countryAndCapital {
    fmt.Println("The capital of", country, "is", countryAndCapital[country])
}
```


### Iterate Over Keys and Values

```go
for country, capital := range countryAndCapital {
    fmt.Println("The capital of", country, "is", capital)
}
```

> ⚠️ Order is **not guaranteed** in maps


## Checking if a Key Exists

```go
capital, ok := countryAndCapital["Japan"]

if ok {
    fmt.Println("The capital of Japan is", capital)
} else {
    fmt.Println("Japan is not in the map")
}
```

Explanation:

* `ok` is a boolean
* `true` → key exists
* `false` → key does not exist


## Deleting an Element

```go
delete(countryAndCapital, "India")
fmt.Println("Entry for India is deleted")
```


## Updating a Map Value

```go
countryAndCapital["USA"] = "Alaska"
```

Maps allow **updating values directly** using the key


## Printing Updated Map

```go
for country, capital := range countryAndCapital {
    fmt.Println("The capital of", country, "is", capital)
}
```


## Length of a Map

```go
fmt.Println(len(countryAndCapital))
```

Returns the number of key:value pairs


## Important Notes (Must Remember)

* Maps are **reference types**
* Assigning a map to another variable copies the reference
* Modifying one modifies the other

```go
m1 := map[string]int{"a": 1}
m2 := m1
m2["a"] = 100
fmt.Println(m1["a"]) // 100
```


## Summary

* Maps store data as **key:value pairs**
* Keys must be unique
* Maps are unordered
* Use `make()` to initialize
* Use `range` to iterate
* Use `delete()` to remove entries
* Maps behave like references


