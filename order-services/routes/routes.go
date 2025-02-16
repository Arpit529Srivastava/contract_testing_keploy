package routes

import (
	"github.com/Arpit529stivastava/order-services/handler"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/orders", handler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", handler.GetOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", handler.GetOrderByID).Methods("GET")
}
