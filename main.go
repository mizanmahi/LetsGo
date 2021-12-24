package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Go-Ticket-Seller!")

	var totalTickets = 50
	var availableTickets = 50
	var userName string
	var ticketPurchased int
	var userEmail string

	var bookings []string

	for {

		fmt.Println("Please enter your name:")
		fmt.Scan(&userName)

		bookings = append(bookings, userName)

		fmt.Println("Please enter your email:")
		fmt.Scan(&userEmail)

		fmt.Println("Hello, " + userName + "! How many tickets would you like to purchase?")
		fmt.Scan(&ticketPurchased)

		if ticketPurchased > availableTickets {
			fmt.Printf("Sorry, we have only %v tickets available. Please try again. \n", availableTickets)
			continue
		}

		fmt.Printf("Hello %v, you have purchased %v tickets. A confirmation email is sent to this address %v \n", userName, ticketPurchased, userEmail)

		availableTickets = availableTickets - ticketPurchased

		fmt.Printf("There are %v tickets available out of %v.\n", availableTickets, totalTickets)
		fmt.Println(bookings)

		for _, booking := range bookings {
			// foreach loop
			fmt.Println(booking)
		}

		if availableTickets == 0 {
			fmt.Println("There are no tickets left. Please try again later.")
			// end program
			break
		}
	}

}
