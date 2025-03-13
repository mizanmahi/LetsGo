package person

import (
	"fmt"
	"time"
)

// struct needs to be initialized with a type
// important!! the struct name should be capitalized to be exported
// important!!  also the fields should be capitalized to be exported
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
func (p person) OutputPersonDataFromPointer() {

	// (p person) this basically a special type of parameter act as a receiver

	// go automatically dereference the p pointer, its an exception for struct
	// normally we need to dereference

	// fmt.Println(p) // &{mizan mahi 19/07/25 {13973790209902737800 10979873601 0x75ed40}}
	// fmt.Println(&p) // 0xc00008e068

	
    // single line
	fmt.Printf("First Name: %s, Last Name: %s, Birthdate: %s, Created At: %s\n", p.firstName, p.lastName, p.birthdate, p.createdAt)
}

// clear user name with pointer receiver
 func (p *person) ClearName() {
	 p.firstName = ""
	 p.lastName = ""
 }

 // 1. creation or constructor function
// func newPerson(firstName, lastName, birthdate string) person {
// 	return person{firstName, lastName, birthdate, time.Now()}
// }

 // 2. creation or constructor function with pointer receiver
func NewPerson(firstName, lastName, birthdate string) *person {
	// returning a pointer also works
	// could be used to validate the data before creating the struct
	return &person{firstName, lastName, birthdate, time.Now()}
}
