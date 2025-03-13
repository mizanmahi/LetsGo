package main

// example of pointers use case
import "fmt"


func outputAge(age *int) {
	// & is used to get the memory address of a variable ex: &age
	// * is used to get the value from a memory address ex: *age
    *age = 40  // Directly modifies original variable
    fmt.Println("Inside function:", *age)
}

func testFunc() {
    age := 32
    outputAge(&age)  // Passing memory address
    fmt.Println("Outside function:", age) // Now modified to 40
	
}
