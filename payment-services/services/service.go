package services

import (
	"errors"
	"log"
	"sync"
)

// In-memory payment store
var payments = make(map[int]string)
var paymentID = 1
var mu sync.Mutex

func ProcessPayment(orderID int, amount float64) (int, error) {
	mu.Lock()
	defer mu.Unlock()

	// Simulate payment processing
	log.Println("Processing payment for order:", orderID, "Amount:", amount)

	payments[paymentID] = "completed"
	log.Println("Payment successful with ID:", paymentID)

	paymentID++
	return paymentID - 1, nil
}

func GetPaymentStatus(id int) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	status, exists := payments[id]
	if !exists {
		return "", errors.New("payment not found")
	}

	return status, nil
}
