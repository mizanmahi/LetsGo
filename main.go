package main

import "fmt"

// fmt is a package that provides I/O functions

// there is basically two types of files in go executable(can be built and produces exe file, must need to have "package main" and a main func) and reusable

func main() {

	card := newCard()



	// default value providing - 1
	greet("Arish", 2)      // Calls original function
	greetWithDefault("Ali") // Uses default age

	// variadic example
	greet2("Arish", 2)  // Uses given age
	greet2("Ali")       // Uses default age 18

	


}

// basic function example return type must be specified
func newCard() string {
	return "Five of Diamonds"
}

// function with parameters
func newCard2(s string) string {
	return s
}

// function with multiple parameters
func newCard3(s string, s2 string) string {
	return s + s2
}

// function with multiple return types
func newCard4(s string, s2 string) (string, string) {
	return s, s2
}

// function with multiple return types and named return types
func newCard5(s string, s2 string) (card string, card2 string) {
	// adding name to return types allows you to return values without specifying return types
	// card and card2 are the named return types and it creates the variables with the same name which can be used in the function and returned automatically
	card = s
	card2 = s2
	return
}


// Function with explicit default value
func greet(name string, age int) {
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

// Wrapper function to provide a default value
func greetWithDefault(name string) {
	greet(name, 18) // Default age = 18
}


func greet2(name string, ages ...int) {
	age := 18 // Default value
	if len(ages) > 0 {
		age = ages[0]
	}
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}