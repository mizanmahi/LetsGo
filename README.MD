# **Golang Packages Cheat-sheet**

## **1. Package Structure**

### **Executable vs. Reusable Packages**

-  **Executable (main package)**

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, Go!")
   }
   ```

   -  Every Go program starts execution from `main` package.
   -  `main()` function is the entry point.

-  **Reusable Package (for imports)**

   ```go
   package mypackage

   func SayHello() string {
       return "Hello from mypackage!"
   }
   ```

---

## **2. Go Module (`go mod`)**

### **Initialize a Module**

```sh
go mod init myproject
```

-  Creates `go.mod` file, tracking dependencies.

### **Download Dependencies**

```sh
go mod tidy
```

-  Cleans unused modules and fetches missing ones.

### **Vendor Dependencies**

```sh
go mod vendor
```

-  Copies dependencies to `vendor/` directory.

### **Update a Dependency**

```sh
go get -u github.com/example/package
```

### **List Dependencies**

```sh
go list -m all
```

---

## **3. Commonly Used Standard Packages**

### **1️⃣ fmt (Formatted I/O)**

```go
import "fmt"

fmt.Println("Hello, World!")
fmt.Printf("Number: %d\n", 42)
```

### **2️⃣ os (OS Interaction)**

```go
import "os"

os.Exit(1)  // Exit with status code 1
```

### **3️⃣ io/ioutil (File Handling)**

```go
import "io/ioutil"

data, _ := ioutil.ReadFile("file.txt")
fmt.Println(string(data))
```

### **4️⃣ strings (String Manipulation)**

```go
import "strings"

str := "hello,world"
parts := strings.Split(str, ",")
fmt.Println(parts[0])  // Output: hello
```

### **5️⃣ strconv (String ↔ Number Conversion)**

```go
import "strconv"

num, _ := strconv.Atoi("42")
fmt.Println(num)  // 42
```

### **6️⃣ time (Date & Time)**

```go
import "time"

now := time.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

### **7️⃣ math (Math Operations)**

```go
import "math"

fmt.Println(math.Sqrt(16))  // 4
```

### **8️⃣ net/http (HTTP Client/Server)**

```go
import "net/http"

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
})
http.ListenAndServe(":8080", nil)
```

### **9️⃣ encoding/json (JSON Handling)**

```go
import "encoding/json"

data := map[string]string{"name": "John"}
jsonData, _ := json.Marshal(data)
fmt.Println(string(jsonData))  // {"name":"John"}
```

### **🔟 log (Logging)**

```go
import "log"

log.Println("This is a log message")
```

---

## **4. Building & Running**

### **Run Without Compilation**

```sh
go run main.go
```

-  Executes the Go file.

### **Compile to Executable**

```sh
go build -o myapp main.go
```

-  Generates `myapp` executable.

---

## **5. Custom Packages & Imports**

### **Create a Package**

📁 **Project Structure**

```
myproject/
│── main.go
│── greetings/
│   ├── greetings.go
│   ├── greetings_test.go
```

### **greetings/greetings.go**

```go
package greetings

import "fmt"

func Hello(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}
```

### **main.go (Using Custom Package)**

```go
package main

import (
    "fmt"
    "myproject/greetings"
)

func main() {
    msg := greetings.Hello("Alice")
    fmt.Println(msg)
}
```

### **Run the Project**

```sh
go run main.go
```

---

## **6. Dependency Management with `go get`**

-  **Install a Package**
   ```sh
   go get github.com/gin-gonic/gin
   ```
-  **Use Installed Package**
   ```go
   import "github.com/gin-gonic/gin"
   ```
