package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handler.ListProducts)
	// app.Post("/fact", handlers.CreateFact)
	// app.Get("/fact/:id", handlers.GetFactById)
	// app.Delete("/fact/:id", handlers.DeleteFact)
}
