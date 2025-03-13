package main

import (
	"fmt"
	"structPrac/person"
)


func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data! 
	newPerson := person.NewPerson(firstName, lastName, birthdate)
	// outputPersonData(newPerson)
	// outputPersonDataFromPointer(&newPerson)
	newPerson.OutputPersonDataFromPointer() // need to pass the struct instance for the receiver parameter
	newPerson.ClearName()
}

// func outputPersonData(p person) {
	
// 	fmt.Printf("First Name: %s, Last Name: %s, Birthdate: %s, Created At: %s\n", p.firstName, p.lastName, p.birthdate, p.createdAt)

// }


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
