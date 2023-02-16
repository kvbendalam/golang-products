package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kvbendalam/golangproducts/database"
	"github.com/kvbendalam/golangproducts/models"
)

func setupRoutes(app *fiber.App) {
	app.Get("/product", ListProducts)
	app.Post("/product", CreateProduct)
	app.Get("/product/:id", GetProductById)
	app.Delete("/product/:id", DeleteProduct)
}

func ListProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Db.Find(&products)
	return c.Status(200).JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&product)
	return c.Status(200).JSON(product)
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")
	var products models.Product

	result := database.DB.Db.Find(&products, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&products)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if result := database.DB.Db.First(&product, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	database.DB.Db.Delete(&product)

	return c.Status(200).JSON(&product)
}
