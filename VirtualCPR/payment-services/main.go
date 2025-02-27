package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create a router
	router := http.NewServeMux()

	// Payment handler
	router.HandleFunc("/pay", func(w http.ResponseWriter, r *http.Request) {
		var paymentRequest struct {
			ID      string `json:"id"`
			Payment string `json:"payment"`
		}

		// Decode the incoming JSON request
		if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
			http.Error(w, "Invalid Payload", http.StatusBadRequest)
			return
		}

		// Forward the request to the order-service
		requestBody, err := json.Marshal(paymentRequest)
		if err != nil {
			http.Error(w, "Error encoding payment request", http.StatusInternalServerError)
			return
		}

		// Send the request to the order-service
		resp, err := http.Post("http://order_services_container:8081/update-payment", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			http.Error(w, "Error contacting order-service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Check the response from the order-service
		if resp.StatusCode == http.StatusOK {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Payment status updated successfully 😎"})
			fmt.Println("Payment SuccessFull 💰")
		} else {
			w.WriteHeader(resp.StatusCode)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update payment status 😢"})
		}
	})

	// Start the payment-service
	fmt.Println("Payment Service running on port 8082 😎...")
	log.Fatal(http.ListenAndServe(":8082", router))
}