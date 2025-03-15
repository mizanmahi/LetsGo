# **Goroutines, Channels, and Select in Go** 🚀

## **1. Goroutines (Concurrency in Go)**

A **goroutine** is a lightweight thread managed by the Go runtime. You can start a goroutine using the `go` keyword.

### **Example: Running a Function as a Goroutine**

```go
package main

import (
	"fmt"
	"time"
)

func printMessage() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Goroutine", i)
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

func main() {
	go printMessage() // Start goroutine
	fmt.Println("Main function executing...")

	time.Sleep(3 * time.Second) // Allow goroutine to finish
}
```

🔹 **Output:**

```
Main function executing...
Hello from Goroutine 1
Hello from Goroutine 2
...
```

---

## **2. Channels (Communication Between Goroutines)**

A **channel** is used to safely communicate between goroutines.

### **Declaring and Using Channels**

```go
ch := make(chan int)  // Creating an integer channel
```

### **Example: Sending and Receiving from a Channel**

```go
package main

import (
	"fmt"
)

func sendMessage(ch chan string) {
	ch <- "Hello from Goroutine!" // Sending data to channel
}

func main() {
	ch := make(chan string) // Create a channel
	go sendMessage(ch)      // Start goroutine

	msg := <-ch // Receive data from channel
	fmt.Println(msg) // Output: Hello from Goroutine!
}
```

---

## **3. Buffered Channels (Async Communication)**

Buffered channels allow sending multiple values before a receiver is required.

### **Example: Buffered Channel**

```go
package main

import "fmt"

func main() {
	ch := make(chan string, 2) // Buffer size of 2

	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch) // Output: Hello
	fmt.Println(<-ch) // Output: World
}
```

🔹 If you send more data than the buffer size, it blocks until space is available.

---

## **4. Closing a Channel**

A **closed channel** signals that no more data will be sent.

### **Example: Closing and Iterating Over a Channel**

```go
package main

import "fmt"

func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // Close the channel
}

func main() {
	ch := make(chan int)
	go sendNumbers(ch)

	for num := range ch { // Reads until channel is closed
		fmt.Println(num)
	}
}
```

---

## **5. Select Statement (Multiple Channel Operations)**

The `select` statement allows a goroutine to **wait on multiple channels**.

### **Example: Using Select**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from ch1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	// Waiting for whichever channel responds first
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
```

🔹 The **first available** channel is selected.

---

## **6. Timeout with Select**

If a goroutine takes too long, we can **set a timeout**.

### **Example: Timeout Handling**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Done!"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(2 * time.Second): // Timeout if no response in 2 sec
		fmt.Println("Timeout!")
	}
}
```

🔹 If no message is received within **2 seconds**, it prints **"Timeout!"**.

---

## **7. Non-Blocking Channel Operations**

Avoid blocking by using `default` in a `select` statement.

### **Example: Non-Blocking Send & Receive**

```go
package main

import "fmt"

func main() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received") // Non-blocking
	}
}
```

🔹 Since the channel is empty, it **doesn’t block** and prints **"No message received"**.

---

## **8. Worker Pool using Goroutines and Channels**

A **worker pool** efficiently distributes tasks among multiple goroutines.

### **Example: Worker Pool**

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulating work
		results <- job * 2 // Returning result
	}
}

func main() {
	const numWorkers = 3
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs to process

	// Collect results
	for k := 1; k <= 5; k++ {
		fmt.Println("Result:", <-results)
	}
}
```

🔹 **3 workers** process **5 jobs** concurrently.

---

## **9. Deadlocks in Goroutines**

A deadlock happens when all goroutines are waiting for data that **never arrives**.

### **Example: Deadlock (Incorrect Code)**

```go
package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 10 // Deadlock: No receiver!
	fmt.Println(<-ch)
}
```

🔹 **Fix:** Use a goroutine or buffer:

```go
go func() { ch <- 10 }()
fmt.Println(<-ch)
```

---

## **10. Fan-In and Fan-Out Patterns**

### **Fan-In (Multiple Goroutines Sending to One Channel)**

```go
package main

import (
	"fmt"
	"time"
)

func producer(ch chan string, msg string) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan string)

	go producer(ch, "Worker 1")
	go producer(ch, "Worker 2")

	for i := 0; i < 6; i++ {
		fmt.Println(<-ch)
	}
}
```

### **Fan-Out (One Channel Sending to Multiple Goroutines)**

```go
package main

import "fmt"

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d processed job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
}
```

---

## **Summary**

✅ **Goroutines** → Lightweight concurrency (`go myFunction()`).  
✅ **Channels** → Safe communication between goroutines (`ch := make(chan int)`).  
✅ **Buffered Channels** → Allow async message passing (`make(chan int, 2)`).  
✅ **Closing Channels** → Signals no more data (`close(ch)`).  
✅ **Select** → Waits on multiple channels (`select { case msg := <-ch: ... }`).  
✅ **Timeouts** → Prevents waiting indefinitely (`time.After()`).  
✅ **Worker Pool** → Efficient concurrency management.
