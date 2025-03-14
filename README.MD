# **Go Functions Cheat Sheet** 🚀

## **1. Declaring a Function**

A function in Go starts with the `func` keyword, followed by the function name, parameters, return type, and the function body.

```go
func greet() {
    fmt.Println("Hello, World!")
}

func main() {
    greet() // Output: Hello, World!
}
```

---

## **2. Function with Parameters**

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("Mizan") // Output: Hello, Mizan
}
```

---

## **3. Function with Multiple Parameters**

```go
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 7)
    fmt.Println(result) // Output: 12
}
```

---

## **4. Function with Multiple Return Values**

```go
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("Hello", "World")
    fmt.Println(a, b) // Output: World Hello
}
```

---

## **5. Named Return Values**

You can name return values, which makes the function more readable.

```go
func divide(a, b float64) (quotient float64, remainder float64) {
    quotient = a / b
    remainder = float64(int(a) % int(b))
    return // Implicit return of named values
}

func main() {
    q, r := divide(10, 3)
    fmt.Println(q, r) // Output: 3.3333333333333335 1
}
```

---

## **6. Variadic Functions (Multiple Arguments of the Same Type)**

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15
}
```

---

## **7. Function as a Value (First-Class Functions)**

Functions can be assigned to variables and passed around.

```go
func multiply(a, b int) int {
    return a * b
}

func main() {
    op := multiply
    fmt.Println(op(3, 4)) // Output: 12
}
```

---

## **8. Anonymous (Lambda) Functions**

You can define functions without a name.

```go
func main() {
    sum := func(a, b int) int {
        return a + b
    }
    fmt.Println(sum(10, 5)) // Output: 15
}
```

---

## **9. Function Closures**

A function can return another function that retains state.

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    increment := counter()
    fmt.Println(increment()) // Output: 1
    fmt.Println(increment()) // Output: 2
}
```

---

## **10. Passing Functions as Arguments**

```go
func applyOperation(a, b int, op func(int, int) int) int {
    return op(a, b)
}

func main() {
    add := func(x, y int) int { return x + y }
    fmt.Println(applyOperation(5, 3, add)) // Output: 8
}
```

---

## **11. Defer (Delayed Execution)**

`defer` postpones function execution until surrounding function exits.

```go
func main() {
    defer fmt.Println("This is printed last")
    fmt.Println("This is printed first")
}

// Output:
// This is printed first
// This is printed last
```

---

## **12. Panic & Recover (Handling Errors)**

Use `panic()` to terminate a function and `recover()` to catch panics.

```go
func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    if b == 0 {
        panic("division by zero")
    }

    fmt.Println("Result:", a/b)
}

func main() {
    safeDivide(10, 0) // Output: Recovered from panic: division by zero
}
```

---

## **13. Recursive Function**

A function calling itself (useful for factorial, Fibonacci, etc.).

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5)) // Output: 120
}
```

---

## **14. Higher-Order Functions (Function Returning a Function)**

```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := multiplier(2)
    fmt.Println(double(5)) // Output: 10
}
```

---

## **15. Methods (Functions with Structs)**

Methods in Go are functions that operate on structs.

```go
type Person struct {
    Name string
    Age  int
}

// Method on struct
func (p Person) greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    p := Person{"Alice", 30}
    p.greet() // Output: Hello, my name is Alice and I am 30 years old.
}
```

---

## **16. Pointer Receiver Methods (Modifying Structs)**

Use pointer receivers (`*`) to modify struct fields.

```go
type Counter struct {
    Value int
}

func (c *Counter) Increment() {
    c.Value++
}

func main() {
    c := Counter{}
    c.Increment()
    fmt.Println(c.Value) // Output: 1
}
```

---

## **17. Function with Interface Parameters**

```go
type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof!")
}

func makeSpeak(s Speaker) {
    s.Speak()
}

func main() {
    d := Dog{}
    makeSpeak(d) // Output: Woof!
}
```

---

## **18. Function Pointer Example**

You can use function pointers for dynamic function calls.

```go
func sayHello() {
    fmt.Println("Hello, World!")
}

func executeFunction(f func()) {
    f()
}

func main() {
    executeFunction(sayHello) // Output: Hello, World!
}
```

---

## **19. Benchmarking Functions**

Use the `testing` package for benchmarking.

```go
import "testing"

func BenchmarkSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = sum(1, 2, 3, 4, 5)
    }
}
```

---

## **20. Function Composition (Combining Functions)**

```go
func addOne(x int) int {
    return x + 1
}

func double(x int) int {
    return x * 2
}

func compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

func main() {
    addThenDouble := compose(double, addOne)
    fmt.Println(addThenDouble(3)) // Output: 8
}
```

---

## **Summary**

✅ Functions in Go are **first-class citizens** (pass them as arguments, return them).  
✅ Supports **multiple return values**, **named returns**, and **variadic arguments**.  
✅ Methods can be **attached to structs** (value or pointer receivers).  
✅ **Defer**, **panic**, and **recover** help with execution flow & error handling.
