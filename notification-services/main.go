package main

import (
	"log"

	"github.com/Arpit529stivastava/notification-services/config"
	"github.com/Arpit529stivastava/notification-services/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to Database
	config.ConnectDB()

	// Initialize Router
	r := gin.Default()

	// Setup Routes
	routes.SetupRoutes(r)

	// Start Server
	log.Println("Notification Service is running on port 8004")
	if err := r.Run(":8004"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
