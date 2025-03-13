package main

import (
	"fmt"
	"time"
)

// struct needs to be initialized with a type
type person struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// 1. regular printing
// func outputPersonData(p person) {
	
// 	fmt.Printf("First Name: %s, Last Name: %s, Birthdate: %s, Created At: %s\n", p.firstName, p.lastName, p.birthdate, p.createdAt)

// }

// 2. printing by passing the reference pointer
// func outputPersonDataFromPointer(p *person) {

// 	// go automatically dereference the p pointer, its an exception for struct
// 	// normally we need to dereference

// 	fmt.Println(p) // &{mizan mahi 19/07/25 {13973790209902737800 10979873601 0x75ed40}}
// 	fmt.Println(&p) // 0xc00008e068

	
//     // single line
// 	fmt.Printf("First Name: %s, Last Name: %s, Birthdate: %s, Created At: %s\n", p.firstName, p.lastName, p.birthdate, p.createdAt)
// }

// 3. printing by making a struct method
func (p person) outputPersonDataFromPointer() {

	// (p person) this basically a special type of parameter act as a receiver

	// go automatically dereference the p pointer, its an exception for struct
	// normally we need to dereference

	fmt.Println(p) // &{mizan mahi 19/07/25 {13973790209902737800 10979873601 0x75ed40}}
	fmt.Println(&p) // 0xc00008e068

	
    // single line
	fmt.Printf("First Name: %s, Last Name: %s, Birthdate: %s, Created At: %s\n", p.firstName, p.lastName, p.birthdate, p.createdAt)
}

// clear user name with pointer receiver
 func (p *person) clearName() {
	 p.firstName = ""
	 p.lastName = ""
 }

 // creation or constructor function
func newPerson(firstName, lastName, birthdate string) person {
	return person{firstName, lastName, birthdate, time.Now()}
}


func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data! 
	newPerson := person{firstName, lastName, birthdate, time.Now()}
	// outputPersonData(newPerson)
	// outputPersonDataFromPointer(&newPerson)
	newPerson.outputPersonDataFromPointer() // need to pass the struct instance for the receiver parameter
	newPerson.clearName()
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
