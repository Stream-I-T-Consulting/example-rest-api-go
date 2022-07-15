package main

import (
	"stream-it-consulting/innovation-team/example-rest-api/config"
	"stream-it-consulting/innovation-team/example-rest-api/database"
	"stream-it-consulting/innovation-team/example-rest-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//  Load config from .env file
	config.LoadConfig()

	//  Initialize database
	dbConn := database.Initialize()
	defer dbConn.Close()

	//  Initialize Fiber Framework
	app := fiber.New()

	// Root route
	router.HTTPRootRoute(app)
	// Book routes
	router.HTTPBookRoutes(app, dbConn)

	// Start server
	app.Listen(":3000")
}
