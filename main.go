package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/auth"
	"github.com/khemingkapat/fiber_example/handlers"
	"github.com/khemingkapat/fiber_example/queries"
)

func main() {
	conn_str := connStrSelect("localhost")
	db := queries.InitDB(conn_str)

	app := fiber.New()
	app.Use("/people", auth.AuthRequired)
	app.Get("/people", handlers.GetPeopleHandler(db))

	app.Get("/people/:id", handlers.GetPersonHandler(db))

	app.Post("/people", handlers.CreatePersonHandler(db))

	app.Put("/people/:id", handlers.UpdatePersonHandler(db))

	app.Delete("/people/:id", handlers.DeletePersonHandler(db))

	app.Post("/register", handlers.CreateUserHandler(db))

	app.Post("/login", handlers.LoginUserHandler(db))

	app.Listen(":8080")
}
