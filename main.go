package main

// fmt is a package that provides I/O functions

// there is basically two types of files in go executable(can be built and produces exe file, must need to have "package main" and a main func) and reusable

func main() {

	// variables
	var name string = "John"
	var age int = 30
	var isCool = true
	var size float32 = 2.3

	// name = 50 // cannot reassign a variable with a different type in go

	// shorthand
	name2 := "John"
	age2 := 30
	isCool2 := true
	size2 := 2.3

	// multiple variables
	var name3, email string = "John", "john@gmail.com"
	name4, email4 := "John", "john@gmail.com"



	// const
	const isCool3 = true
	// const isCool3 = false // cannot reassign


}
