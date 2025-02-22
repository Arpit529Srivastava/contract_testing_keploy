package main

import (
	"fmt"
	"log"
	"net/http"
	"order-services/database"
	"order-services/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to MongoDB
	database.ConnectMongoDB()
	database.CreateOrdersCollection()

	// Create Router
	router := mux.NewRouter()

	// Register Routes
	routes.RegisterRoutes(router)

	fmt.Println("Order Service running on port 8081 😎...")
	log.Fatal(http.ListenAndServe(":8081", router))
}