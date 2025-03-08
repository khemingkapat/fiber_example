package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/handlers"
	"github.com/khemingkapat/fiber_example/queries"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=admin password=admin dbname=dorm host=postgres port=5432 sslmode=disable"
	db, err := queries.InitDB(connStr)
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
