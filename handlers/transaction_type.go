package handlers

import (
	"github.com/gofiber/fiber/v2"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListTransactionTypes(c *fiber.Ctx) error {
	var transactionTypes []models.TransactionType

	config.Database.Find(&transactionTypes)
	return c.Status(200).JSON(transactionTypes)
}

func GetTransactionType(c *fiber.Ctx) error {
	id := c.Params("id")
	var transactionType models.TransactionType

	result := config.Database.Find(&transactionType, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&transactionType)
}

func AddTransactionType(c *fiber.Ctx) error {
	transactionType := new(models.TransactionType)

	if err := c.BodyParser(transactionType); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&transactionType)
	return c.Status(201).JSON(transactionType)
}

func UpdateTransactionType(c *fiber.Ctx) error {
	transactionType := new(models.TransactionType)
	id := c.Params("id")

	if err := c.BodyParser(transactionType); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&transactionType)
	return c.Status(200).JSON(transactionType)
}

func RemoveTransactionType(c *fiber.Ctx) error {
	id := c.Params("id")
	var transactionType models.TransactionType

	result := config.Database.Delete(&transactionType, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
