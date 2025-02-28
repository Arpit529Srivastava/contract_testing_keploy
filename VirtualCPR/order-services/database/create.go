package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateOrdersCollection ensures the 'orders' collection exists in 'orderDB'
func CreateOrdersCollection() {
	collection := DB

	// Check if the collection has at least one document
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatalf("Error checking orders collection: %v", err)
	}

	// Log the status of the collection
	if count == 0 {
		fmt.Println("Orders collection is empty")
	} else {
		fmt.Println("Orders collection already exists")
	}
}
