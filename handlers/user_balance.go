package handlers

import (
	"github.com/gofiber/fiber/v2"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListUserBalances(c *fiber.Ctx) error {
	var userBalances []models.UserBalance

	config.Database.Find(&userBalances)
	return c.Status(200).JSON(userBalances)
}

func GetUserBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	var userBalance models.UserBalance

	result := config.Database.Find(&userBalance, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&userBalance)
}

func AddUserBalance(c *fiber.Ctx) error {
	userBalance := new(models.UserBalance)

	if err := c.BodyParser(userBalance); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&userBalance)
	return c.Status(201).JSON(userBalance)
}

func UpdateUserBalance(c *fiber.Ctx) error {
	userBalance := new(models.UserBalance)
	id := c.Params("id")

	if err := c.BodyParser(userBalance); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&userBalance)
	return c.Status(200).JSON(userBalance)
}

func RemoveUserBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	var userBalance models.UserBalance

	result := config.Database.Delete(&userBalance, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
