// in go every file is called a package and we need to provide a name for the package, for the main file it is generally main.
package main

// in go we have many packages predefined which we can use in our application by importing, one of them is fmt which is a package that contains functions for formatting outputs.
import (
	"fmt"
)

// go needs this function declaration otherwise go does'nt know where to start the code execution, the main function is the entry point of a go program, 1 main function per go application.
func main() {

	// creating variables
	var conferenceName = "Go Seminar" // Go infer the type of the variable based on the value assigned to it, and will throw error if later want to assign a different type to the variable. it could also be written like this with type declaration 
	// var conferenceName string = "Go Seminar"

	// const is used when we want to create a variable that will never change, it is a constant variable.
	const conferenceTickets int = 50

	var availableTickets = 50

	//@ we can also use the short hand notation for creating variables, the short hand notation is the := operator.

	// we can print the type of a variable by using the %T placeholder
	fmt.Printf("The type of the conferenceName is %T\n", conferenceName)

	// fmt.Println =>  this is called printline function.
	// fmt.Printf => is used when we want to print a formatted string, it is a formatted print function.

	fmt.Printf("Let's go to the %v to learn more about Go \n", conferenceName)
	fmt.Printf("Only %v tickets are available out of %v \n", conferenceTickets, availableTickets)

	// when we create a variable without assigning any value to it, we need to provide the type of the variable.
	var userName string
	var userTickets int // some other int types are int8, int16, int32, int64, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128

	fmt.Printf("Please enter your name: ")
	fmt.Scan(&userName) // this is a function that reads the input from the user and stores it in the variable. & is used to pass the address of the variable. which is called a pointer.

	fmt.Printf("Please the amount of ticket: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Hello %v, you have just purchased %v ticket\n", userName, userTickets)

	availableTickets = availableTickets - userTickets

	fmt.Printf("Only %v tickets are available out of %v \n", availableTickets, conferenceTickets)

}
