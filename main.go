package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/handlers"
	"github.com/khemingkapat/fiber_example/queries"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Successfully connected to the database!")
	db, err := queries.InitDB()
	if err != nil {
		log.Println(err)
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("fiber example for dorm management")
	})
	app.Get("/people", handlers.GetPeopleHandler(db))

	log.Fatal(app.Listen(":8080"))
}
