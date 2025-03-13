package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// fmt is a package that provides I/O functions

// there is basically two types of files in go executable(can be built and produces exe file, must need to have "package main" and a main func) and reusable

func main() {
	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Using default balance of 1000")

		// alternatively we can use panic
		// panic("Could not read balance from file")
	}

	// in GO we have only for loop and while loop is not there
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Enter your pin: ", i)
	// }

	// infinite loop, since there is no condition
	fmt.Println("Welcome to Go bank")
	 for {

	 	fmt.Println("Choose an action")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var userChoice int
		fmt.Scan(&userChoice)
		fmt.Printf("You have selected %d\n", userChoice)


	// using switch case
	switch userChoice {
		case 1:	fmt.Println("Your account balance is", accountBalance)
		case 2: fmt.Println("Enter the amount to deposit")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			saveBalanceToFile(accountBalance)
			fmt.Println("Your account balance is", accountBalance)
		case 3: fmt.Println("Enter the amount to withdraw")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance")
				continue;
			} else {
				accountBalance -= withdrawAmount
				saveBalanceToFile(accountBalance)
				fmt.Println("Your account balance is", accountBalance)
			}
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount")
				continue;
			}
		case 4: fmt.Println("Thank you for using Go bank")
		   
			fmt.Println("Your account balance is", accountBalance)
			 return;
		default: fmt.Println("Invalid choice")
			return;
		}

	}

	
}

func saveBalanceToFile(balance float64) {

	balanceString :=  fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceString), 0644)
}

func readBalanceFromFile() (float64, error) {
	balance, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000, errors.New("file not found")
	}
	balanceString := string(balance)
	return strconv.ParseFloat(balanceString, 64)

}