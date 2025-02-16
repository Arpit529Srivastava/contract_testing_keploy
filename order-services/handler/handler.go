package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Arpit529stivastava/order-services/database"
	"github.com/Arpit529stivastava/order-services/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Check if user ID exists in database by calling User Service
func CheckUserID(userID int) bool {
	url := fmt.Sprintf("http://localhost:8080/users/%d", userID) 
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error contacting User Service:", err)
		return false
	}
	defer resp.Body.Close()

	var result map[string]bool
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error decoding response from User Service:", err)
		return false
	}

	return result["exists"]
}

// Create Order Handler
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	// Decode the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	// Convert UserID to an integer
	userID, err := strconv.Atoi(order.UserID)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	// Validate User ID in PostgreSQL
	if !CheckUserID(userID) {
		http.Error(w, "User doesn't exist in the Database 🙁. Please create an account first. 👍", http.StatusUnauthorized)
		return
	}

	// Insert order into MongoDB
	order.CreatedAt = time.Now()
	collection := database.GetCollection("orders") // Ensure consistent collection name

	_, err = collection.InsertOne(context.TODO(), order)
	if err != nil {
		http.Error(w, "Error inserting data into MongoDB", http.StatusInternalServerError)
		log.Println("MongoDB Insert Error:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order placed successfully 😎"})
}

// Get All Orders Handler
func GetOrders(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("orders") 
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching orders", http.StatusInternalServerError)
		log.Println("MongoDB Query Error:", err)
		return
	}
	defer cursor.Close(context.TODO())

	var orders []models.Order
	if err := cursor.All(context.TODO(), &orders); err != nil {
		http.Error(w, "Error reading orders", http.StatusInternalServerError)
		log.Println("MongoDB Cursor Error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

// Get Order By ID Handler
func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	// Extract the order ID from the URL
	vars := mux.Vars(r)
	orderID := vars["id"] // Extract ID from the URL path

	if orderID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	// Fetch the order from MongoDB by ID (MongoDB's _id is used for querying)
	collection := database.GetCollection("orders")
	var order models.Order
	err := collection.FindOne(context.TODO(), bson.M{"_id": orderID}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching order", http.StatusInternalServerError)
			log.Println("MongoDB Query Error:", err)
		}
		return
	}

	// Return the order details
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

