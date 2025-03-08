package main

import (
	"errors"
	"fmt"
	"log"

	// "github.com/gofiber/fiber/v2"
	// "github.com/khemingkapat/fiber_example/handlers"
	// "github.com/khemingkapat/fiber_example/objects"
	"github.com/khemingkapat/fiber_example/queries"
)

func main() {
	conn_str, err := connStrSelect("localhost")
	if err != nil {
		panic("invalid mode")
	}
	db, err := queries.InitDB(conn_str)
	if err != nil {
		log.Fatal(err)
	}

	currPerson := queries.GetPerson(db, 1)
	fmt.Println(currPerson)

	people := queries.GetPeople(db)
	fmt.Println(people)

	currPerson.Firstname = "New Name"
	queries.UpdatePerson(db, currPerson)

	queries.DeletePerson(db, 5)
}

func connStrSelect(mode string) (string, error) {
	var host string
	var port int

	if mode == "postgres" {
		host = "postgres"
		port = 5432
	} else if mode == "localhost" {
		host = "localhost"
		port = 5430
	} else {
		return "", errors.New("invalid mode")
	}

	result := fmt.Sprintf("user=admin password=admin dbname=dorm host=%s port=%d sslmode=disable", host, port)
	return result, nil
}
