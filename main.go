package main

import "fmt"

// fmt is a package that provides I/O functions

// there is basically two types of files in go executable(can be built and produces exe file, must need to have "package main" and a main func) and reusable

func main() {
	accountBalance := 1000.0




	fmt.Println("Welcome to Go bank")
	fmt.Println("Choose an action")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var userChoice int
	fmt.Scan(&userChoice)
	fmt.Printf("You have selected %d\n", userChoice)

	// if userChoice == 1 {
	// 	fmt.Println("Your account balance is", accountBalance)
	// } else if userChoice == 2 {
	// 	fmt.Println("Enter the amount to deposit")
	// 	var depositAmount float64
	// 	fmt.Scan(&depositAmount)
	// 	accountBalance += depositAmount
	// 	fmt.Println("Your account balance is", accountBalance)
	// } else if userChoice == 3 {
	// 	fmt.Println("Enter the amount to withdraw")
	// 	var withdrawAmount float64
	// 	fmt.Scan(&withdrawAmount)
	// 	if withdrawAmount > accountBalance {
	// 		fmt.Println("Insufficient balance")
	// 	} else {
	// 		accountBalance -= withdrawAmount
	// 		fmt.Println("Your account balance is", accountBalance)
	// 	}
	// 	if withdrawAmount <= 0 {
	// 		fmt.Println("Invalid amount")
	// 		return;
	// 	}
	// } else if userChoice == 4 {
	// 	fmt.Println("Thank you for using Go bank")
	// } else {
	// 	fmt.Println("Invalid choice")
	// }

	// using switch case
	switch userChoice {
	case 1:	fmt.Println("Your account balance is", accountBalance)
	case 2: fmt.Println("Enter the amount to deposit")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Your account balance is", accountBalance)
	case 3: fmt.Println("Enter the amount to withdraw")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient balance")
		} else {
			accountBalance -= withdrawAmount
			fmt.Println("Your account balance is", accountBalance)
		}
		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount")
			return;
		}
	case 4: fmt.Println("Thank you for using Go bank")
	default: fmt.Println("Invalid choice")
	}
}