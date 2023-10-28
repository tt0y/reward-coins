package handlers

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListUserBalances(c *fiber.Ctx) error {
	rows, err := config.Database.Query("SELECT id, user_id, coin_type_id, amount FROM user_balances")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var userBalances []models.UserBalance
	for rows.Next() {
		var userBalance models.UserBalance
		err := rows.Scan(&userBalance.ID, &userBalance.UserId, &userBalance.CoinTypeId, &userBalance.Amount)
		if err != nil {
			log.Fatal(err)
		}
		userBalances = append(userBalances, userBalance)
	}

	return c.Status(200).JSON(userBalances)
}

func GetUserBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	var userBalance models.UserBalance

	err := config.Database.QueryRow("SELECT id, user_id, coin_type_id, amount FROM user_balances WHERE id = ?", id).Scan(&userBalance.ID, &userBalance.UserId, &userBalance.CoinTypeId, &userBalance.Amount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.SendStatus(404)
		}
		log.Fatal(err)
	}

	return c.Status(200).JSON(&userBalance)
}

func AddUserBalance(c *fiber.Ctx) error {
	userBalance := new(models.UserBalance)

	if err := c.BodyParser(userBalance); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("INSERT INTO user_balances (user_id, coin_type_id, amount) VALUES (?, ?, ?)",
		userBalance.UserId, userBalance.CoinTypeId, userBalance.Amount)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(201).JSON(userBalance)
}

func UpdateUserBalance(c *fiber.Ctx) error {
	userBalance := new(models.UserBalance)
	id := c.Params("id")

	if err := c.BodyParser(userBalance); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("UPDATE user_balances SET user_id = ?, coin_type_id = ?, amount = ? WHERE id = ?",
		userBalance.UserId, userBalance.CoinTypeId, userBalance.Amount, id)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(userBalance)
}

func RemoveUserBalance(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.Database.Exec("DELETE FROM user_balances WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(200)
}
