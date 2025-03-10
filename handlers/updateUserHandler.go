package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/objects"
	"github.com/khemingkapat/fiber_example/queries"
	"gorm.io/gorm"
)

func UpdateUserHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error Getting ID")
		}

		person := new(object.User)
		if err = c.BodyParser(person); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error Parsing Person")
		}

		person.ID = uint(id)

		err = queries.UpdateUser(db, person)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.JSON(person)
	}
}
