package main

import (
	"log"
	"net/http"
	"os"
	"user-services/database"
	"user-services/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables from .env file

	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()
	// Create the users table
	//database.CreateUsersTable() // Pass DB connection
	err = database.CreateUsersTable(db)
	if err != nil {
		log.Fatal("Database setup error:", err)
	}
	// Setup the router
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r, db) // Pass DB connection

	// Get the port from environment variables, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server 🧿
	log.Printf("User service running on port %s 👍", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
