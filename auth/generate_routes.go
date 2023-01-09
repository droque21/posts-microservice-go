package main

import (
	"github.com/droque21/posts-microservice-go/auth/handlers"
	"github.com/gofiber/fiber/v2"
)

func GenerateRoutes(app *fiber.App) {
	api := app.Group("/api/v1/auth")

	api.Post("/signup", handlers.Create)
}
