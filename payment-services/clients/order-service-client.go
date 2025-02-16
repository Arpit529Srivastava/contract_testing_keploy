package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Order struct {
	ID         string  `json:"_id"`
	UserEmail  string  `json:"user_email"`
	Product    string  `json:"product"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	CreatedAt  string  `json:"created_at"`
}

func GetOrderDetails(orderID string) (*Order, error) {
	url := fmt.Sprintf("http://localhost:8081/order/%s", orderID) // Ensure URL is correct

	resp, err := http.Get(url)
	if err != nil {
		// Log the error if the order service is unreachable
		log.Printf("Error while reaching order service: %v", err)
		return nil, errors.New("failed to reach order service")
	}
	defer resp.Body.Close()

	// Log status code for further debugging
	log.Printf("Received status code %d from order service for orderID %s", resp.StatusCode, orderID)

	if resp.StatusCode != http.StatusOK {
		// Log the response body in case of errors
		var responseMessage map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&responseMessage)
		log.Printf("Order service responded with error: %v", responseMessage)
		return nil, errors.New("order not found")
	}

	var order Order
	if err := json.NewDecoder(resp.Body).Decode(&order); err != nil {
		// Log decoding error
		log.Printf("Error decoding order details: %v", err)
		return nil, err
	}

	return &order, nil
}
