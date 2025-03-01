package book

import (
	"slices"
	"strconv"

	"github.com/gofiber/fiber/v2" //
)

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, book := range Books {
		if book.ID == id {
			Books = slices.Delete(Books, i, i+1)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
