package queries

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) { // Corrected function signature
	connStr := "user=admin password=admin dbname=dorm host=localhost port=5430 sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
