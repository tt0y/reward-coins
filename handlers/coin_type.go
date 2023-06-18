package handlers

import (
	"github.com/gofiber/fiber/v2"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListCoinTypes(c *fiber.Ctx) error {
	var coinTypes []models.CoinType

	config.Database.Find(&coinTypes)
	return c.Status(200).JSON(coinTypes)
}

func GetCoinType(c *fiber.Ctx) error {
	id := c.Params("id")
	var coinType models.CoinType

	result := config.Database.Find(&coinType, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&coinType)
}

func AddCoinType(c *fiber.Ctx) error {
	coinType := new(models.CoinType)

	if err := c.BodyParser(coinType); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&coinType)
	return c.Status(201).JSON(coinType)
}

func UpdateCoinType(c *fiber.Ctx) error {
	coinType := new(models.CoinType)
	id := c.Params("id")

	if err := c.BodyParser(coinType); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&coinType)
	return c.Status(200).JSON(coinType)
}

func RemoveCoinType(c *fiber.Ctx) error {
	id := c.Params("id")
	var coinType models.CoinType

	result := config.Database.Delete(&coinType, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
