package routes

import (
	"database/sql"
	"user-services/handler"

	"github.com/gorilla/mux"
)
func RegisterUserRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/users", handler.CreateUser(db)).Methods("POST")
	r.HandleFunc("/users", handler.GetAllUser(db)).Methods("GET")
	r.HandleFunc("/users/{id}", handler.GetUserByID(db)).Methods("GET")

}