package services

import (
	"errors"
	"log"
	"github.com/Arpit529stivastava/payment-services/clients"
)

func ProcessPayment(orderID string, amount float64, paymentMethod string) error { // Changed to string
	order, err := clients.GetOrderDetails(orderID)
	if err != nil {
		return errors.New("Failed to fetch the Order details")
	}

	// Checking if the amount entered matches the order total
	if order.Price != amount {
		return errors.New("Payment amount does not match order total 🥲")
	}

	// Call payment gateway
	if err := clients.MakePayment(amount, paymentMethod); err != nil {
		return errors.New("payment failed")
	}

	log.Println("Payment successful for Order ID:", orderID)
	return nil
}
