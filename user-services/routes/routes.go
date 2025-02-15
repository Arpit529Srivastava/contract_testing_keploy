package routes

import (
	"database/sql"

	"github.com/Arpit529stivastava/user-services/handler"
	"github.com/gorilla/mux"
)
func RegisterUserRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/users", handler.CreateUser(db)).Methods("POST")
	r.HandleFunc("/users", handler.GetAllUser(db)).Methods("GET")
}