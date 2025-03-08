package queries

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connStr string) (*sql.DB, error) { // Corrected function signature
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
