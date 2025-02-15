package main

import (
	"log"
	"net/http"

	"github.com/Arpit529stivastava/order-services/database"
	"github.com/Arpit529stivastava/order-services/routes"
	"github.com/gorilla/mux"
)


func main(){
	db := database.InitDB()
	defer db.Close()


	/// setup rotuer

	r := mux.NewRouter()
	routes.RegisterOrderRoutes(r,db)

	// SERVER STARTED

	log.Println(("Order service started"))
	log.Fatal(http.ListenAndServe(":8081", r))
}