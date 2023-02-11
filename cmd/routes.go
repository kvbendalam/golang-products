package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kvbendalam/webservices/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/product", handlers.ListProducts)
	app.Post("/products", handlers.CreateProduct)
	// app.Get("/fact/:id", handlers.GetFactById)
	// app.Delete("/fact/:id", handlers.DeleteFact)
}
