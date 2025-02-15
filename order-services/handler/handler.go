package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Arpit529stivastava/order-services/models"
)

// CREATE ALL THE ORDERS

func CreatOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order models.Order
		json.NewDecoder(r.Body).Decode(&order)

		result, err := db.Exec("orders (user_id, amount, status) VALUES (?, ?, ?)",
			order.UserID, order.Amount, "Pending")
		if err != nil {
			http.Error(w, "Failed to create order", http.StatusInternalServerError)
			return
		}

		orderID, _ := result.LastInsertId()
		order.ID = int(orderID)
		json.NewEncoder(w).Encode(order)
	}
}


// GET ALL THE ORDERS

func GetAllOrders(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, user_id, amount, status FROM orders")
		if err != nil {
			http.Error(w, "Error fetching orders", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var orders []models.Order
		for rows.Next() {
			var order models.Order
			rows.Scan(&order.ID, &order.UserID, &order.Amount, &order.Status)
			orders = append(orders, order)
		}

		json.NewEncoder(w).Encode(orders)
	}
}
