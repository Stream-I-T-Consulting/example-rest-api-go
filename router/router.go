package router

import (
	"database/sql"
	"stream-it-consulting/innovation-team/example-rest-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func HTTPRootRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to My REST API",
		})
	})
}

func HTTPBookRoutes(app *fiber.App, dbConn *sql.DB) {
	handler := handlers.NewHandler(dbConn)

	app.Post("/books", handler.CreateBook)
}
