package main

import "fmt"

// don't need to export to use this function as it is in the same package which is main
func communicate () {
	    fmt.Println("Choose an action")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")		
}