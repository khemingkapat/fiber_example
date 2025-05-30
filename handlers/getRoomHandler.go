package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/queries"
	"gorm.io/gorm"
)

func GetRoomHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error Getting ID")
		}
		room, err := queries.GetRoom(db, uint(id))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.JSON(room)
	}
}
