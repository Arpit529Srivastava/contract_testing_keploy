package handler

import (
	"encoding/json"
	"net/http"
	"github.com/Arpit529stivastava/payment-services/services"
)

type PaymentRequest struct {
	OrderID       string  `json:"order_id"`  // Changed to string
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"method"`
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Not a Valid Request 😔", http.StatusBadRequest)
		return
	}

	// Process the Payment
	if err := services.ProcessPayment(req.OrderID, req.Amount, req.PaymentMethod); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment successful😎😎"})
}
