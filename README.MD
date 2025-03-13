**Golang Function Parameter Passing Cheatsheet**, covering **value passing** and **pointer-based mutation** with pros and cons.

---

# **Golang Function Parameter Passing Cheatsheet**

## **1. Passing Values to Functions (Call by Value)**

-  In Go, function arguments are passed **by value** by default.
-  A **copy** of the value is created and passed to the function.

### **Example: Passing by Value**

```go
package main

import "fmt"

func outputAge(age int) {
    age = 40  // This change is local to the function
    fmt.Println("Inside function:", age)
}

func main() {
    age := 32
    outputAge(age)
    fmt.Println("Outside function:", age) // Remains 32
}
```

### **Memory Allocation**

-  **Original variable (`age`)** is stored at a memory address.
-  When passed to `outputAge`, a **new copy** is created at a different memory address.
-  Changes inside the function do **not** affect the original variable.

### **Pros of Pass by Value**

✅ Safer since original data is not modified.  
✅ Avoids unintended side effects.  
✅ Suitable for small values (e.g., int, float, bool).

### **Cons of Pass by Value**

❌ Creates additional memory allocation.  
❌ Inefficient for large data structures (e.g., structs, slices, maps).

---

## **2. Modifying Values Using Pointers (Call by Reference)**

-  Instead of passing a copy, we pass a **memory address (pointer)** so the function modifies the original value.

### **Example: Passing by Pointer**

```go
package main

import "fmt"

func outputAge(age *int) {
    *age = 40  // Directly modifies original variable
    fmt.Println("Inside function:", *age)
}

func main() {
    age := 32
    outputAge(&age)  // Passing memory address
    fmt.Println("Outside function:", age) // Now modified to 40
}
```

### **Memory Allocation**

-  `&age` passes the **memory address** of `age` instead of a copy.
-  `*age = 40` modifies the **actual** value at that memory address.

### **Pros of Using Pointers**

✅ Efficient for large data structures (avoids copying).  
✅ Allows modifying the original variable directly.  
✅ Useful when working with mutable data like slices, maps, and structs.

### **Cons of Using Pointers**

❌ Can lead to **unintended side effects** if not managed properly.  
❌ More complex, requires understanding of memory management.  
❌ Risk of **nil pointer dereference** if not handled properly.

---

## **3. Choosing Between Value & Pointer Passing**

| Scenario                             | Use Value Passing       | Use Pointer Passing            |
| ------------------------------------ | ----------------------- | ------------------------------ |
| Primitive types (int, float)         | ✅ Yes                  | ❌ No (unless mutation needed) |
| Large structs or arrays              | ❌ No (memory overhead) | ✅ Yes (efficient)             |
| Need to modify the original variable | ❌ No                   | ✅ Yes                         |
| Safer, avoids side effects           | ✅ Yes                  | ❌ No                          |
