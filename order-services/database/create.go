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

	// If collection is empty, insert a placeholder document to initialize it
	if count == 0 {
		_, err := collection.InsertOne(ctx, bson.M{"init": "placeholder"})
		if err != nil {
			log.Fatalf("Error initializing orders collection: %v", err)
		}
		fmt.Println("Orders collection initialized successfully")
	} else {
		fmt.Println("Orders collection already exists")
	}
}
