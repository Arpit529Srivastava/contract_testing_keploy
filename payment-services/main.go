package main

import (
	"log"
	"net/http"

	"github.com/Arpit529stivastava/payment-services/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set up routes
	routes.PaymentRoutes(r)

	// Start server on port 8003
	log.Println("Payment Service running on port 8003")
	if err := r.Run(":8003"); err != nil {
		log.Fatal("Failed to start server:", http.StatusInternalServerError)
	}
}
