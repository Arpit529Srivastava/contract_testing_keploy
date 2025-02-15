package database

import (
	"database/sql"
	"log"
)


func InitDB() *sql.DB {
	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/orderservice"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MySQL Database")
	return db
}