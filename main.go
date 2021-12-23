// in go every file is called a package and we need to provide a name for the package, for the main file it is generally main
package main

// in go we have many packages predefined which we can use in our application by importing, one of them is fmt which is a package that contains functions for formatting outputs.
import "fmt"

// go needs this function declaration otherwise go does'nt know where to start the code execution, the main function is the entry point of a go program, 1 main function per go application.
func main() {

	// creating variables
	var conferenceName = "Go Seminar"

	// const is used when we want to create a variable that will never change, it is a constant variable.
	const conferenceTickets = 50
	var availableTickets = 50

	fmt.Printf("Let's go to the %v to learn more about Go \n", conferenceName) // this is called printline function
	fmt.Printf("Only %v tickets are available out of %v \n", conferenceTickets, availableTickets)
}
