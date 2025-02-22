package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Collection

func ConnectMongoDB() {
	// Correct the MongoDB URI to use port 27017
	ctx := options.Client().ApplyURI("mongodb://mongodb:27017/")
	client, err := mongo.Connect(context.TODO(), ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ensure the connection is successful
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB 😊")

	// Set DB to use the 'orderDB' database and the collection name
	DB = client.Database("orderDB").Collection("orders")

	// Defer disconnect to close the connection properly when the function exits
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB
}
