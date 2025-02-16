package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found or couldn't be loaded")
	}

	// Fetching environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// Construct connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		user, password, dbname, host, port, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Cannot connect to the database:", err)
		return nil, err
	}

	fmt.Println("Connection established, let's go! 👌👌")
	return db, nil
}
