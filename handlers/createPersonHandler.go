package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/objects"
	"github.com/khemingkapat/fiber_example/queries"
	"gorm.io/gorm"
)

func CreatePersonHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		person := new(object.Person)
		if err := c.BodyParser(person); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Error Parsing Person")
		}

		err := queries.CreatePerson(db, person)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.JSON(fiber.Map{"message": "Person Created"})
	}
}
