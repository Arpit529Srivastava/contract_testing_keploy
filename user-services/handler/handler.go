package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Arpit529stivastava/user-services/models"
)


// CREATES THE USER

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)
		
		err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)
		if err != nil {
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

// GET ALL THE USER

func GetAllUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name ,email FROM users")
		if err != nil {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var users []models.User
		for rows.Next() {
			var user models.User
			rows.Scan(&user.ID, &user.Name, &user.Email)
			users = append(users, user)
		}
		json.NewEncoder(w).Encode(users)
	}
}
