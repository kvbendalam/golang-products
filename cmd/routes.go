package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kvbendalam/webservices/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListProducts)
	// app.Post("/fact", handlers.CreateFact)
	// app.Get("/fact/:id", handlers.GetFactById)
	// app.Delete("/fact/:id", handlers.DeleteFact)
}
