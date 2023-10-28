package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListCoinTypes(c *fiber.Ctx) error {
	rows, err := config.Database.Query("SELECT id, name, description FROM coin_types")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var coinTypes []models.CoinType
	for rows.Next() {
		var coinType models.CoinType
		err := rows.Scan(&coinType.ID, &coinType.Name, &coinType.Description)
		if err != nil {
			log.Fatal(err)
		}
		coinTypes = append(coinTypes, coinType)
	}

	return c.Status(200).JSON(coinTypes)
}

func GetCoinType(c *fiber.Ctx) error {
	id := c.Params("id")
	var coinType models.CoinType

	err := config.Database.QueryRow("SELECT id, name, description FROM coin_types WHERE id = ?", id).Scan(&coinType.ID, &coinType.Name, &coinType.Description)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.SendStatus(404)
		}
		log.Fatal(err)
	}

	return c.Status(200).JSON(&coinType)
}

func AddCoinType(c *fiber.Ctx) error {
	coinType := new(models.CoinType)

	if err := c.BodyParser(coinType); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("INSERT INTO coin_types (name, description) VALUES (?, ?)", coinType.Name, coinType.Description)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(201).JSON(coinType)
}

func UpdateCoinType(c *fiber.Ctx) error {
	coinType := new(models.CoinType)
	id := c.Params("id")

	if err := c.BodyParser(coinType); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("UPDATE coin_types SET name = ?, description = ? WHERE id = ?", coinType.Name, coinType.Description, id)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(coinType)
}

func RemoveCoinType(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.Database.Exec("DELETE FROM coin_types WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(200)
}
