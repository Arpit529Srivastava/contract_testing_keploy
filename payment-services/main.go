package main

import (
	"log"
	"net/http"

	"github.com/Arpit529stivastava/payment-services/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register routes
	routes.RegisterPayment(r)

	// Start the server
	log.Println("Payment Service is running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", r))
}
