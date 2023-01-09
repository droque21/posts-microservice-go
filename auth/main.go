package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Middlewares
	app.Use(logger.New())

	// Routes
	GenerateRoutes(app)

	// Start server
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
