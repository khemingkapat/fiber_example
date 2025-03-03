package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/queries"
)

func GetPeopleHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		people, err := queries.QueryPeople(db)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(people)
	}
}
