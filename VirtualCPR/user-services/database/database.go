package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// Global DB instance

func InitDB() (*sql.DB, error) {
	// Load environment variables from .env file

	// Fetching environment variables
	user := "postgres"
	password := "Rupam#rani1983"
	dbname := "user_services"
	host := "postgres"
	port := "5432"
	sslmode := "disable"

	// Construct connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		user, password, dbname, host, port, sslmode)

	db, err := sql.Open("postgres", connStr) // Assign to global DB variable
	if err != nil {
		log.Println("Cannot connect to the database:", err)
		return nil, err
	}

	fmt.Println("Connection established, let's go! 👌👌")
	return db, nil
}

// CreateUsersTable explicitly creates the 'users' table if it doesn't exist
func CreateUsersTable(db *sql.DB) error {
	query := `
	DROP TABLE IF EXISTS users;
    CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
    );`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	log.Println("Users table checked/created successfully ✅")
	return nil
}
