package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Collection

func ConnectMongoDB() {

	// Create a new MongoDB client with a proper context and timeout
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017/").
		SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ensure the connection is successful
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB 😊")

	// Select the database and collection
	DB = client.Database("orderDB").Collection("orders")
}

// GetCollection returns the requested collection
func GetCollection(collectionName string) *mongo.Collection {
	return DB
}
