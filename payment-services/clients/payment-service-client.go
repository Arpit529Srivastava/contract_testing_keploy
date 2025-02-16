package clients

import (
	"errors"
	"log"
)

// MakePayment simulates an external payment gateway call
func MakePayment(amount float64, method string) error {
	// Simulating a payment process
	log.Printf("Processing payment of $%.2f using %s...\n", amount, method)

	// Simulate successful payment
	if amount > 0 {
		log.Println("Payment successful!")
		return nil
	}

	return errors.New("payment failed")
}
