package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/Arpit529stivastava/user-services/database"
	"github.com/Arpit529stivastava/user-services/routes"
	"github.com/gorilla/mux"
)

func main() {
	db := database.InitDB()
	defer db.Close()

// setup the router
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r,db)

	// start the server 🧿
	log.Println("User service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
