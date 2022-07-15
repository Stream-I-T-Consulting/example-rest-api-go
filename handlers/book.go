package handlers

import (
	"stream-it-consulting/innovation-team/example-rest-api/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) CreateBook(c *fiber.Ctx) error {
	// Create the book struct
	var book *models.Book

	// Parse the request body into the book struct
	if err := c.BodyParser(&book); err != nil {
		return err
	}

	// Create the book
	newBook, err := h.bookRepo.CreateBook(book)
	if err != nil {
		return err
	}

	// Return the book as JSON
	return c.JSON(newBook)
}
