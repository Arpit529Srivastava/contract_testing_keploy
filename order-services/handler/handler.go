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

	// Set default payment status and creation time
	order.PaymentStatus = "pending" // Default payment status
	order.EmailStatus = "not sent"
	order.CreatedAt = time.Now()

	// Insert order into MongoDB
	collection := database.GetCollection("orders")
	_, err = collection.InsertOne(context.TODO(), order)
	if err != nil {
		http.Error(w, "Error inserting data into MongoDB", http.StatusInternalServerError)
		log.Println("MongoDB Insert Error:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order placed successfully 😎"})
	fmt.Println("Order placed succesfully")
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
	vars := mux.Vars(r)
	orderID := vars["id"]

	if orderID == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}
func GetOrderByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderEmail := vars["user_email"]
	if orderEmail == "" {
		http.Error(w, "Order email is required", http.StatusBadRequest)
		return
	}
	collection := database.GetCollection("orders")
	var order models.Order
	err := collection.FindOne(context.TODO(), bson.M{"user_email": orderEmail}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Order not found", http.StatusNotFound)
			log.Fatal("error finding the user_email")
		} else {
			http.Error(w, "Error fetching order", http.StatusInternalServerError)
			log.Println("MongoDB Query Error:", err)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// Update Payment Status Handler
func UpdatePaymentStatus(w http.ResponseWriter, r *http.Request) {
	var paymentRequest struct {
		ID      string `json:"id"`
		Payment string `json:"payment"`
	}

	// Decode the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	// Update payment status in MongoDB
	collection := database.GetCollection("orders")
	filter := bson.M{"_id": paymentRequest.ID}
	update := bson.M{"$set": bson.M{"payment_status": "completed"}} // Update status to "completed"


	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, "Error updating payment status", http.StatusInternalServerError)
		log.Println("MongoDB Update Error:", err)
		return
	}

	if result.ModifiedCount == 0 {
		http.Error(w, "Order not found or payment status already updated", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment status updated successfully 😎"})
}

func UpdateEmailStatus(w http.ResponseWriter, r *http.Request) {
	var EmailRequest struct {
		ID string `json:"id"` // ID is a string
	}

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&EmailRequest); err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("orders")

	// 🔹 Match by string "_id" instead of ObjectID
	filter := bson.M{"_id": EmailRequest.ID} 
	update := bson.M{"$set": bson.M{"email_status": "sent"}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, "Error updating email status: "+err.Error(), http.StatusInternalServerError)
		log.Println("MongoDB Update Error:", err)
		return
	}

	if result.ModifiedCount == 0 {
		http.Error(w, "Order not found or email status already updated", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email status updated successfully 😎"})
}
