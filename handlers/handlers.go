package handlers

import (
	"database/sql"
	"stream-it-consulting/innovation-team/example-rest-api/repositories"

	"github.com/gofiber/fiber/v2"
)

type (
	handler struct {
		bookRepo repositories.IBookRepository
	}
	Handler interface {
		// GetBooks(c *fiber.Ctx) error
		// GetBookByID(c *fiber.Ctx) error
		CreateBook(c *fiber.Ctx) error
		// UpdateBook(c *fiber.Ctx) error
		// DeleteBook(c *fiber.Ctx) error
	}
)

// NewHandler creates a new handler
func NewHandler(dbConn *sql.DB) handler {
	return handler{
		bookRepo: repositories.NewBookRepository(dbConn),
	}
}
