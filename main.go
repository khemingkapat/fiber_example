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

	// Register and Login routes
	app.Post("/register", handlers.CreateUserHandler(db))
	app.Post("/login", handlers.LoginUserHandler(db))

	app.Use(auth.JWTMiddleware)

	// Group for /people paths (only accessible by managers)
	people := app.Group("/people", auth.ManagerOnlyMiddleware)
	people.Get("/", handlers.GetUsersHandler(db))         // Get all users
	people.Get("/:id", handlers.GetUserHandler(db))       // Get a specific user by ID
	people.Put("/:id", handlers.UpdateUserHandler(db))    // Update user by ID
	people.Delete("/:id", handlers.DeleteUserHandler(db)) // Delete user by ID

	rooms := app.Group("/rooms")

	rooms.Get("/", func(c *fiber.Ctx) error {
		if auth.IsManager(c) {
			return handlers.GetRoomsHandler(db)(c)
		}
		return handlers.GetRoomByTenantIDHandler(db)(c)
	})

	rooms.Get("/tenant", auth.ManagerOnlyMiddleware, handlers.GetRoomWithTenantHandler(db)) // Get rooms with tenant (for managers)
	rooms.Get("/:id", auth.ManagerOnlyMiddleware, handlers.GetRoomHandler(db))              // Get a specific room by ID (for managers)

	// Start the server
	app.Listen(":8080")
}
