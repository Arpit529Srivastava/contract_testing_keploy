package database
import(
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "user=postgres password=yourpassword dbname=userservice sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} 
	return db
}