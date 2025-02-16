package database

import (
	"database/sql"
	"fmt"
	//"log"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB) {
	connStr := "user=postgres password=Rupam#rani1983 dbname=User-service host=localhost port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("cannot connect to the db ")
	} 
	fmt.Println("connection established")
	return db
}