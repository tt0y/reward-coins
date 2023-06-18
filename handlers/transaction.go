package handlers

import (
	"github.com/gofiber/fiber/v2"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListTransactions(c *fiber.Ctx) error {
	var transactions []models.Transaction

	config.Database.Find(&transactions)
	return c.Status(200).JSON(transactions)
}

func GetTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	var transaction models.Transaction

	result := config.Database.Find(&transaction, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&transaction)
}

func AddTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&transaction)
	return c.Status(201).JSON(transaction)
}

func UpdateTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)
	id := c.Params("id")

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&transaction)
	return c.Status(200).JSON(transaction)
}

func RemoveTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	var transaction models.Transaction

	result := config.Database.Delete(&transaction, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
