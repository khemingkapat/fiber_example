package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/queries"
	"gorm.io/gorm"
)

func GetPeopleHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(queries.GetPeople(db))
	}
}
