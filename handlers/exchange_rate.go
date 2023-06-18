package handlers

import (
	"github.com/gofiber/fiber/v2"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListExchangeRates(c *fiber.Ctx) error {
	var exchangeRates []models.ExchangeRate

	config.Database.Find(&exchangeRates)
	return c.Status(200).JSON(exchangeRates)
}

func GetExchangeRate(c *fiber.Ctx) error {
	id := c.Params("id")
	var exchangeRate models.ExchangeRate

	result := config.Database.Find(&exchangeRate, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&exchangeRate)
}

func AddExchangeRate(c *fiber.Ctx) error {
	exchangeRate := new(models.ExchangeRate)

	if err := c.BodyParser(exchangeRate); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&exchangeRate)
	return c.Status(201).JSON(exchangeRate)
}

func UpdateExchangeRate(c *fiber.Ctx) error {
	exchangeRate := new(models.ExchangeRate)
	id := c.Params("id")

	if err := c.BodyParser(exchangeRate); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&exchangeRate)
	return c.Status(200).JSON(exchangeRate)
}

func RemoveExchangeRate(c *fiber.Ctx) error {
	id := c.Params("id")
	var exchangeRate models.ExchangeRate

	result := config.Database.Delete(&exchangeRate, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
