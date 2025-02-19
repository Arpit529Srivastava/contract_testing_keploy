package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/Arpit529stivastava/user-services/usermodels"
	"github.com/gorilla/mux"
)

// CreateUser - Inserts a new user into the database
func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			log.Println("Error decoding request body:", err)
			return
		}

		err := db.QueryRow(`INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`, user.Name, user.Email).Scan(&user.ID)
		if err != nil {
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
			log.Println("Database error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
		fmt.Println("User Created Successfully 🙌🏻")
	}
}

// GetUserByID - Retrieves a user by ID
// GetUserByID - Retrieves a user by ID and checks if the user exists
// GetUserByID - Checks if a user exists by their ID
func GetUserByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["id"]

		// Variable to hold the result of the query
		var exists bool

		// Query to check if the user exists
		err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`, userID).Scan(&exists)
		
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			log.Println("Database error:", err)
			return
		}

		// Return response with boolean indicating whether the user exists
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"exists": exists})
	}
}



// GetAllUser - Retrieves all users from the database
func GetAllUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT id, name, email FROM users`)
		if err != nil {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
			log.Println("Database query error:", err)
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
				log.Println("Error scanning row:", err)
				http.Error(w, "Error reading users", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

