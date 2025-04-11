// package main

// import "fmt"

// type PaymentGateway interface {
// 	pay()
// }

// type payment struct {
// 	ID        string `json:"id"`
// 	Amount    int    `json:"amount"`
// 	Currency  string `json:"currency"`
// 	Status    string `json:"status"`
// 	CreatedAt string `json:"created_at"`
// }

// func NewPayment() payment {
// 	return payment{
// 		ID:        "12345",
// 		Amount:    100,
// 		Currency:  "USD",
// 		Status:    "pending",
// 		CreatedAt: "2023-08-10",
// 	}
// }

// func (p payment) pay(gateway PaymentGateway) {
// 	// paymentInfo := NewBkash()
// 	// paymentInfo.payWithBkash()
// 	// paymentInfo := NewNagad()
// 	// paymentInfo.payWithNagad()
// 	gateway.pay()
// }

// type bkash struct {
// 	ID        string `json:"id"`
// 	Amount    int    `json:"amount"`
// 	Currency  string `json:"currency"`
// 	Status    string `json:"status"`
// 	CreatedAt string `json:"created_at"`
// }

// func NewBkash() bkash {
// 	return bkash{
// 		ID:        "12345",
// 		Amount:    100,
// 		Currency:  "USD",
// 		Status:    "pending",
// 		CreatedAt: "2023-08-10",
// 	}
// }

// type nagad struct {
// 	ID        string `json:"id"`
// 	Amount    int    `json:"amount"`
// 	Currency  string `json:"currency"`
// 	Status    string `json:"status"`
// 	CreatedAt string `json:"created_at"`
// }

// func NewNagad() nagad {
// 	return nagad{
// 		ID:        "12345",
// 		Amount:    100,
// 		Currency:  "USD",
// 		Status:    "pending",
// 		CreatedAt: "2023-08-10",
// 	}
// }

// func (n nagad) pay() {
// 	fmt.Println("Payment created by nagad!")
// }

// func (b bkash) pay() {
// 	fmt.Println("Payment created by bkash!")
// }

// func main() {

// 	payment := NewPayment()
// 	gateway := NewNagad()
// 	payment.pay(gateway)
// }

// refactored version =============

package main

import (
	"fmt"
	"time"
)

// PaymentGateway defines the interface for payment processors
type PaymentGateway interface {
	Pay() error // Changed to return error for better error handling
}

// Payment represents a generic payment transaction
// Renamed from 'payment' to follow Go naming conventions (exported types start with capital letter)
type Payment struct {
	ID        string    `json:"id"`
	Amount    int       `json:"amount"`
	Currency  string    `json:"currency"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"` // Changed to time.Time for proper date handling
}

// NewPayment creates a new Payment instance with default values
func NewPayment() *Payment {
	// Return a pointer to avoid unnecessary copying
	return &Payment{
		ID:        "12345",
		Amount:    100,
		Currency:  "USD",
		Status:    "pending",
		CreatedAt: time.Now().UTC(), // Use actual current time
	}
}

// ProcessPayment processes a payment using the provided gateway
// Renamed from 'pay' to be more descriptive
func (p *Payment) ProcessPayment(gateway PaymentGateway) error {
	return gateway.Pay() // Return any error from the payment process
}

// BkashGateway represents the Bkash payment processor
// Renamed from 'bkash' to be more descriptive and follow naming conventions
type BkashGateway struct {
	ProviderID string    `json:"provider_id"`
	CreatedAt  time.Time `json:"created_at"`
	// Removed duplicate fields that were already in Payment
}

// NewBkashGateway creates a new Bkash payment gateway
func NewBkashGateway() *BkashGateway {
	return &BkashGateway{
		ProviderID: "bkash-provider-123",
		CreatedAt:  time.Now().UTC(),
	}
}

// Pay implements the PaymentGateway interface for BkashGateway
func (b *BkashGateway) Pay() error {
	fmt.Println("Payment processed by Bkash!")
	return nil // Return nil error for successful payment
}

// NagadGateway represents the Nagad payment processor
// Renamed from 'nagad' to be more descriptive and follow naming conventions
type NagadGateway struct {
	ProviderID string    `json:"provider_id"`
	CreatedAt  time.Time `json:"created_at"`
	// Removed duplicate fields that were already in Payment
}

// NewNagadGateway creates a new Nagad payment gateway
func NewNagadGateway() *NagadGateway {
	return &NagadGateway{
		ProviderID: "nagad-provider-456",
		CreatedAt:  time.Now(),
	}
}

// Pay implements the PaymentGateway interface for NagadGateway
func (n *NagadGateway) Pay() error {
	fmt.Println("Payment processed by Nagad!")
	return nil // Return nil error for successful payment
}

func main() {
	// Create a new payment
	payment := NewPayment()
	
	// Process with Nagad gateway
	nagadGateway := NewNagadGateway()
	if err := payment.ProcessPayment(nagadGateway); err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}
	
	// Process with Bkash gateway (added to demonstrate multiple gateways)
	bkashGateway := NewBkashGateway()
	if err := payment.ProcessPayment(bkashGateway); err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}
}
