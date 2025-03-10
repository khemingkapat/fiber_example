package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/auth"
	object "github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func LoginUserHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(object.UserLogin)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		token, err := auth.LoginUser(db, user)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 3),
			HTTPOnly: true,
		})

		return c.JSON(fiber.Map{
			"message": "Login Succeeded",
		})
	}
}
