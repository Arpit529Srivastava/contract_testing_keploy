package routes

import (
	"database/sql"

	"github.com/Arpit529stivastava/order-services/handler"
	"github.com/gorilla/mux"
)


func RegisterOrderRoutes (r *mux.Router, db *sql.DB) {
	r.HandleFunc("/orders", handler.CreatOrder(db)).Methods("POST")
	r.HandleFunc("/orders", handler.GetAllOrders(db)).Methods("GET")
}