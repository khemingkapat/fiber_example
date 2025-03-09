package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/auth"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(auth.User)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		err := auth.CreateUser(db, user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.JSON(fiber.Map{"message": "User Registered"})
	}
}
