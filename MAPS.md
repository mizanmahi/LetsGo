# **Go Maps Cheat Sheet** 🚀

## **1. What is a Map in Go?**

-  A **map** is a built-in data structure that **stores key-value pairs**.
-  **Keys must be unique** and of a comparable type.
-  Values can be of any type.

```go
var myMap map[string]int // Declare a map (uninitialized)
```

---

## **2. Declaring & Initializing a Map**

```go
// Using make (Recommended)
myMap := make(map[string]int)

// Using a map literal
userAges := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
fmt.Println(userAges) // Output: map[Alice:25 Bob:30]
```

---

## **3. Adding & Updating Elements**

```go
userAges := make(map[string]int)

userAges["Alice"] = 25 // Adding a key-value pair
userAges["Bob"] = 30
userAges["Alice"] = 26 // Updating an existing key

fmt.Println(userAges) // Output: map[Alice:26 Bob:30]
```

---

## **4. Accessing Elements**

```go
age := userAges["Alice"]
fmt.Println(age) // Output: 26
```

---

## **5. Checking If a Key Exists**

```go
age, exists := userAges["Charlie"]

if exists {
    fmt.Println("Charlie's age:", age)
} else {
    fmt.Println("Charlie not found!") // Output: Charlie not found!
}
```

---

## **6. Deleting a Key**

```go
delete(userAges, "Bob")
fmt.Println(userAges) // Output: map[Alice:26]
```

---

## **7. Iterating Over a Map**

```go
for key, value := range userAges {
    fmt.Printf("%s is %d years old\n", key, value)
}
```

---

## **8. Getting the Length of a Map**

```go
fmt.Println(len(userAges)) // Output: 1
```

---

## **9. Nested Maps**

```go
users := map[string]map[string]string{
    "Alice": {
        "age":  "25",
        "city": "New York",
    },
    "Bob": {
        "age":  "30",
        "city": "London",
    },
}

fmt.Println(users["Alice"]["city"]) // Output: New York
```

---

## **10. Using Maps with Structs**

```go
type Person struct {
    Age  int
    City string
}

people := map[string]Person{
    "Alice": {Age: 25, City: "Paris"},
    "Bob":   {Age: 30, City: "Berlin"},
}

fmt.Println(people["Alice"].City) // Output: Paris
```

---

## **11. Difference Between `nil` Map & Empty Map**

```go
var nilMap map[string]int  // nil map (uninitialized)
emptyMap := make(map[string]int) // Empty but initialized

fmt.Println(nilMap == nil)  // Output: true
fmt.Println(emptyMap == nil) // Output: false

// nilMap["key"] = 10 // This will cause a runtime error
emptyMap["key"] = 10 // This works fine
```

---

### **12. Convert Map Keys or Values to a Slice**

```go
keys := make([]string, 0, len(userAges))
values := make([]int, 0, len(userAges))

for k, v := range userAges {
    keys = append(keys, k)
    values = append(values, v)
}

fmt.Println(keys)   // Output: [Alice]
fmt.Println(values) // Output: [26]
```

---

### **13. Sorting a Map by Keys**

```go
import (
	"fmt"
	"sort"
)

func main() {
	myMap := map[string]int{"Charlie": 35, "Alice": 25, "Bob": 30}

	// Extract keys
	keys := make([]string, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}

	// Sort keys
	sort.Strings(keys)

	// Print sorted map
	for _, k := range keys {
		fmt.Println(k, ":", myMap[k])
	}
}

// Output:
// Alice : 25
// Bob : 30
// Charlie : 35
```

---

### **14. Comparing Two Maps**

Go does **not** support direct map comparisons (`map1 == map2`), but you can compare them manually:

```go
func areMapsEqual(m1, m2 map[string]int) bool {
    if len(m1) != len(m2) {
        return false
    }
    for k, v := range m1 {
        if v2, ok := m2[k]; !ok || v2 != v {
            return false
        }
    }
    return true
}

m1 := map[string]int{"A": 1, "B": 2}
m2 := map[string]int{"A": 1, "B": 2}
fmt.Println(areMapsEqual(m1, m2)) // Output: true
```

---

## **15. Performance Considerations**

✅ **Use Maps for Fast Lookup** (O(1) time complexity).  
❌ **Avoid Large Maps with Frequent Deletions** (Can cause memory leaks).  
✅ **Use `sync.Map` for Concurrent Access in Goroutines**.

---

## **Conclusion**

Maps are powerful for **storing key-value pairs** and allow **fast lookups**.  
However, they are **unordered**, so sorting requires additional steps.
