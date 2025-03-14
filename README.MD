A **cheat sheet** for **Go interfaces** and **generics**. 🚀

---

# **Go Interfaces & Generics Cheat Sheet**

## **1. Interfaces in Go**

### What is an Interface?

-  A **contract** that defines a **set of methods** that a type must implement.
-  Interfaces **only specify behavior**, not implementation.
-  Any type that implements **all methods** of an interface **automatically satisfies it** (no explicit declaration needed).

### Declaring an Interface:

```go
type Speaker interface {
    Speak() string
}
```

### Implementing an Interface

-  A type satisfies an interface if it implements **all methods** defined by the interface.
-  **No need to explicitly specify** that a type implements an interface.

```go
type Person struct {
    Name string
}

// Implementing the interface method
func (p Person) Speak() string {
    return "Hello, I am " + p.Name
}

// This function can accept any type that implements the Speaker interface
func SpeakSomething(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    p := Person{Name: "Mizan"}
    fmt.Println(p.Speak()) // "Hello, I am Mizan"

    var s Speaker = p // Assign struct to interface
    fmt.Println(s.Speak()) // "Hello, I am Mizan"
}
```

---

### Interface with Multiple Methods

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(s string)
}

type ReaderWriter interface {
    Reader
    Writer
}

// Implementing multiple interfaces
type File struct{}

func (f File) Read() string {
    return "Reading from file"
}

func (f File) Write(s string) {
    fmt.Println("Writing:", s)
}

func main() {
    var rw ReaderWriter = File{}
    fmt.Println(rw.Read()) // Works because File implements Read
    rw.Write("Hello")      // Works because File implements Writer
}
```

---

## **Go Generics (Since Go 1.18)**

Generics allow defining functions or types that work with **any data type**.

### Generic Function Example

```go
package main

import "fmt"

// Generic function
func Print[T any](value T) {
    fmt.Println(value)
}

func main() {
    Print(42)      // 42
    Print("hello") // hello
}
```

### Generic Structs

```go
type Box[T any] struct {
    value T
}

func (b Box[T]) Get() T {
    return b.value
}

func main() {
    intBox := Box[int]{value: 42}
    fmt.Println("Int Box:", intBox.value) // Int Box: 42

    strBox := Box[string]{value: "Hello"}
    fmt.Println("String Box:", strBox.value) // Hello
}
```

### Generic Interface

```go
package main

import "fmt"

type Printer[T any] interface {
    Print(T)
}

type ConsolePrinter[T any] struct{}

func (c ConsolePrinter[T]) Print(data T) {
    fmt.Println(data)
}

func main() {
    var intPrinter ConsolePrinter[int]
    intPrinter.Print(100) // Prints: 100

    var strPrinter ConsolePrinter[string]
    strPrinter.Print("Hello, Generics!") // Hello, Generics!
}
```

---

### Type Constraints

Go 1.18 introduced the `constraints` package for defining **type constraints**.

```go
package main

import "fmt"

// Constraint: T must be int or float64
type Number interface {
    int | float64
}

// Generic function with constraint
func Add[T Number](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Add(10, 20))     // Works with integers
    fmt.Println(Add(3.5, 2.2)) // Works with floats
}
```

---

### Using `any` as a Generic Type

-  **`any`** is an alias for `interface{}` in Go 1.18+.
-  It allows any type, but doesn't enforce method constraints.

```go
func PrintAnything[T any](value T) {
    fmt.Println(value)
}

func main() {
    PrintAnything(100)
    PrintAnything("Go Generics!")
}
```

---

### Generic Methods in Struct

```go
type Container[T any] struct {
    value T
}

func (c *Container[T]) GetValue() T {
    return c.value
}

func main() {
    c := Container[int]{value: 10}
    fmt.Println(c.GetValue()) // 10

    strContainer := Container[string]{value: "Hello"}
    fmt.Println(strContainer.GetValue()) // Hello
}
```

---

### Type Switch with Interface

If you use `any`, you may need **type assertions** or **type switches**:

```go
func PrintValue(value any) {
    switch v := value.(type) {
    case string:
        fmt.Println("String:", v)
    case int:
        fmt.Println("Integer:", v)
    default:
        fmt.Println("Unknown Type")
    }
}

func main() {
    PrintAnything(42)      // Integer: 42
    PrintAnything("Go")    // String: Go
    PrintAnything(3.14)    // Unknown Type
}
```

---

### Benefits of Generics

✅ **Code Reusability** – Write reusable functions and structs that work with different data types.  
✅ **Type Safety** – Ensures that type safety is maintained at **compile time**.  
✅ **Avoids Code Duplication** – No need to write the same logic for multiple types.  
✅ **Improved Performance** – Reduces the need for type assertions, improving runtime efficiency.
