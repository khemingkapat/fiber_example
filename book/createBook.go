// book/create_book.go
package book

import "github.com/gofiber/fiber/v2"

func CreateBook(c *fiber.Ctx) error {
	title := c.FormValue("title")
	author := c.FormValue("author")

	if title == "" || author == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Title and author are required")
	}

	newBook := Book{
		ID:     len(Books) + 1,
		Title:  title,
		Author: author,
	}

	Books = append(Books, newBook)

	return c.JSON(newBook)
}
