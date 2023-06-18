package handlers

import (
	"github.com/gofiber/fiber/v2"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListProducts(c *fiber.Ctx) error {
	var products []models.Product

	config.Database.Find(&products)
	return c.Status(200).JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	result := config.Database.Find(&product, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&product)
}

func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&product)
	return c.Status(201).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	id := c.Params("id")

	if err := c.BodyParser(product); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&product)
	return c.Status(200).JSON(product)
}

func RemoveProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	result := config.Database.Delete(&product, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
