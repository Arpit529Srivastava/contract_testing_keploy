package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)


type Order struct {
	ID           string    `json:"id"`
	UserEmail    string    `json:"user_email"`
	Product      string    `json:"product"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	PaymentStatus string   `json:"payment_status"`
}

func main() {
	go func (){
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for range ticker.C{
			checkpaymentstatus()
		}
	}()


	http.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		checkpaymentstatus()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string {"message" : "Payment method triggered"})
	})
	fmt.Println("Notification services are running on localhost:8084 😚")
	log.Fatal(http.ListenAndServe(":8084", nil))


}


func checkpaymentstatus() {
	resp, err := http.Get("http://localhost:8081/orders")
	if err != nil {
		log.Println("Error fetching the order from the order database 😭😭")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("Failed to fetch the orders ❌")
		return
	}
	var orders []Order 
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil{
		log.Println("Error Decoding the data 😵‍💫")
		return
	}

	// checking for the completed payments

	for _, check := range orders {
		if check.PaymentStatus == "completed"{
			log.Printf("Payment successful 👏🏻 for order ID: %s, User: %s, Product: %s\n", check.ID, check.UserEmail, check.Product)
		}
	}

}