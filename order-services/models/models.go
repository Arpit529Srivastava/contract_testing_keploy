package models

import "time"

type Order struct {
    UserID        string       `json:"id" bson:"_id,omitempty"` // Change string to int
    UserEmail     string    `json:"user_email" bson:"user_email"`
    Product       string    `json:"product" bson:"product"`
    Quantity      int       `json:"quantity" bson:"quantity"`
    Price         float64   `json:"price" bson:"price"`
    CreatedAt     time.Time `json:"created_at" bson:"created_at"`
}

