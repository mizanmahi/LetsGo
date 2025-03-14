# **Go Arrays Cheat Sheet** 🚀

## **1. What is an Array in Go?**

-  **Fixed-size** collection of elements of the **same type**.
-  Cannot grow or shrink in size.
-  Indexed from `0` (first element) to `n-1` (last element).

```go
var arr [5]int // Declaring an array of size 5 with default values (0)
fmt.Println(arr) // Output: [0 0 0 0 0]
```

---

## **2. Initializing Arrays**

```go
// Declare and Initialize
var nums = [5]int{1, 2, 3, 4, 5}
fmt.Println(nums) // Output: [1 2 3 4 5]

// Using shorthand
nums2 := [3]int{10, 20, 30}
fmt.Println(nums2) // Output: [10 20 30]

// Let Go determine the size
nums3 := [...]int{100, 200, 300}
fmt.Println(nums3) // Output: [100 200 300]

// Default values (all 0)
var defaultArr [4]int
fmt.Println(defaultArr) // Output: [0 0 0 0]
```

---

## **3. Accessing & Modifying Elements**

```go
arr := [3]string{"Go", "Rust", "Python"}

// Access element by index
fmt.Println(arr[0]) // Output: Go

// Modify element
arr[2] = "JavaScript"
fmt.Println(arr) // Output: [Go Rust JavaScript]
```

---

## **4. Looping Through an Array**

```go
nums := [5]int{10, 20, 30, 40, 50}

// Using for loop (index-based)
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}

// Using range (index, value)
for index, value := range nums {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Ignoring index (_)
for _, value := range nums {
    fmt.Println(value)
}
```

---

## **5. Array Length**

```go
arr := [4]string{"Go", "Python", "Java", "C++"}
fmt.Println(len(arr)) // Output: 4
```

---

## **6. Copying Arrays (Value Copy)**

-  Arrays are **copied by value**, meaning changes in one **do not** affect the other.

```go
arr1 := [3]int{1, 2, 3}
arr2 := arr1  // Copying array (new independent array)
arr2[0] = 100 // Changing arr2 does not affect arr1

fmt.Println(arr1) // Output: [1 2 3]
fmt.Println(arr2) // Output: [100 2 3]
```

---

## **7. Multidimensional Arrays**

```go
var matrix [2][3]int // 2 rows, 3 columns
matrix = [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

// Accessing elements
fmt.Println(matrix[0][1]) // Output: 2

// Looping through 2D array
for _, row := range matrix {
    for _, col := range row {
        fmt.Print(col, " ")
    }
    fmt.Println()
}
```

---

## **8. Array Manipulations**

### **Reverse an Array**

```go
arr := [5]int{1, 2, 3, 4, 5}
n := len(arr)

for i := 0; i < n/2; i++ {
    arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
}

fmt.Println(arr) // Output: [5 4 3 2 1]
```

---

### **Find the Maximum Value**

```go
arr := [5]int{10, 50, 30, 90, 20}
max := arr[0]

for _, v := range arr {
    if v > max {
        max = v
    }
}

fmt.Println("Max value:", max) // Output: Max value: 90
```

---

### **Sum of Array Elements**

```go
arr := [4]int{5, 10, 15, 20}
sum := 0

for _, v := range arr {
    sum += v
}

fmt.Println("Sum:", sum) // Output: Sum: 50
```

---

### **Checking if an Element Exists**

```go
arr := [5]int{10, 20, 30, 40, 50}
target := 30
found := false

for _, v := range arr {
    if v == target {
        found = true
        break
    }
}

fmt.Println("Element exists:", found) // Output: Element exists: true
```

---

## **9. Difference Between Arrays & Slices**

| Feature     | Array                                    | Slice                                              |
| ----------- | ---------------------------------------- | -------------------------------------------------- |
| Size        | Fixed                                    | Dynamic                                            |
| Memory      | Stores all elements in contiguous memory | References underlying array                        |
| Copying     | Creates a new copy                       | Shares memory with the original                    |
| Flexibility | Limited                                  | More flexible (supports appending, resizing, etc.) |

---

### **10. Convert Array to Slice**

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[:] // Convert entire array to slice
fmt.Println(slice) // Output: [1 2 3 4 5]
```
