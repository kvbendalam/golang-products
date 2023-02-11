package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/kvbendalam/webservices/database"
	"github.com/kvbendalam/webservices/models"
)

func ListProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.DB.Db.Find(&products)
	return c.Status(200).JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(product.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&product)
	return c.Status(200).JSON(product)
}

// func GetFactById(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var facts models.Fact

// 	result := database.DB.Db.Find(&facts, id)

// 	if result.RowsAffected == 0 {
// 		return c.SendStatus(404)
// 	}

// 	return c.Status(200).JSON(&facts)
// }

// func DeleteFact(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var fact models.Fact

// 	if result := database.DB.Db.First(&fact, id); result.Error != nil {
// 		fmt.Println(result.Error)
// 	}

// 	database.DB.Db.Delete(&fact)

// 	return c.Status(200).JSON(&fact)
// }
