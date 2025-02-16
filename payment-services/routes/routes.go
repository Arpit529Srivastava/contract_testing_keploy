package routes

import (

	"github.com/Arpit529stivastava/payment-services/handler"
	"github.com/gorilla/mux"
)

func RegisterPayment(r *mux.Router) {
	r.HandleFunc("/pay", handler.MakePayment).Methods("POST")
}
