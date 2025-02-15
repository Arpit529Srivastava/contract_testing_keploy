package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("postgres", "postgres://user:password@localhost:5432/notifications?sslmode=disable")

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	log.Println("Connected to PostgreSQL")
}
