package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/fiber_example/book"
)

func main() {
	app := fiber.New()

	// Initialize in-memory data
	book.Books = append(book.Books, book.Book{ID: 1, Title: "1984", Author: "George Orwell"})
	book.Books = append(book.Books, book.Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})

	// CRUD routes
	app.Get("/books", book.GetBooks)
	app.Get("/book/:id", book.GetBook)
	app.Post("/book", book.CreateBook)
	app.Put("/book/:id", book.UpdateBook)
	app.Delete("/book/:id", book.DeleteBook)

	app.Listen(":8080")
}

// Handlers
