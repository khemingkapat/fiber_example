package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/auth"
	object "github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(object.User)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		err := auth.CreateUser(db, user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"message": "User Created",
		})
	}
}
