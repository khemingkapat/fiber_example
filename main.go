package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connection string (hardcoded credentials)
	connStr := "user=admin password=adminpassword dbname=dorm host=localhost port=5430 sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")

	// You can now perform database operations using the 'db' object
	// Example:
	// rows, err := db.Query("SELECT * FROM your_table")
	// ...
}
