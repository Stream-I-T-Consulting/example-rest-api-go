package handlers

import (
	"strconv"
	"stream-it-consulting/innovation-team/example-rest-api/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBooks(c *fiber.Ctx) error {
	// Get the books
	books, err := h.bookRepo.GetBooks()
	if err != nil {
		return err
	}

	return c.JSON(books)
}

func (h handler) GetBookByID(c *fiber.Ctx) error {
	// Get the book ID
	id, err := strconv.Atoi(c.Params("id"))
	if id == 0 || err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid book ID",
		})
	}

	// Get the book
	book, err := h.bookRepo.GetBookByID(id)
	if err != nil {
		return err
	}

	return c.JSON(book)
}

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
