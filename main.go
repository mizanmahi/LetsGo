package main

import "fmt"

func main() {

	// anonymous function
	func() {
		println("Hello World")
	}() 
	// function with parameter
	func(a int, b int) {
		println(a + b)
	}(1, 2) 

	fmt.Println(super(5))
	
	
	
}
func super(n int) int {
	if n == 1 {
		return 1
	}
	return n * super(n-1)
}

// variadic function
func sum(num ...int) int {
	total := 0
	for _, v := range num {
		total += v
	}
	return total
}

