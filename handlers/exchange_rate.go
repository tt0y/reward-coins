package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListExchangeRates(c *fiber.Ctx) error {
	rows, err := config.Database.Query("SELECT id, coin_type_id_from, coin_type_id_to, rate FROM exchange_rates")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var exchangeRates []models.ExchangeRate
	for rows.Next() {
		var exchangeRate models.ExchangeRate
		err := rows.Scan(&exchangeRate.ID, &exchangeRate.CoinTypeIdFrom, &exchangeRate.CoinTypeIdTo, &exchangeRate.Rate)
		if err != nil {
			log.Fatal(err)
		}
		exchangeRates = append(exchangeRates, exchangeRate)
	}

	return c.Status(200).JSON(exchangeRates)
}

func GetExchangeRate(c *fiber.Ctx) error {
	id := c.Params("id")
	var exchangeRate models.ExchangeRate

	err := config.Database.QueryRow("SELECT id, coin_type_id_from, coin_type_id_to, rate FROM exchange_rates WHERE id = ?", id).Scan(&exchangeRate.ID, &exchangeRate.CoinTypeIdFrom, &exchangeRate.CoinTypeIdTo, &exchangeRate.Rate)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.SendStatus(404)
		}
		log.Fatal(err)
	}

	return c.Status(200).JSON(&exchangeRate)
}

func AddExchangeRate(c *fiber.Ctx) error {
	exchangeRate := new(models.ExchangeRate)

	if err := c.BodyParser(exchangeRate); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("INSERT INTO exchange_rates (coin_type_id_from, coin_type_id_to, rate) VALUES (?, ?, ?)", exchangeRate.CoinTypeIdFrom, exchangeRate.CoinTypeIdTo, exchangeRate.Rate)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(201).JSON(exchangeRate)
}

func UpdateExchangeRate(c *fiber.Ctx) error {
	exchangeRate := new(models.ExchangeRate)
	id := c.Params("id")

	if err := c.BodyParser(exchangeRate); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("UPDATE exchange_rates SET coin_type_id_from = ?, coin_type_id_to = ?, rate = ? WHERE id = ?", exchangeRate.CoinTypeIdFrom, exchangeRate.CoinTypeIdTo, exchangeRate.Rate, id)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(exchangeRate)
}

func RemoveExchangeRate(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.Database.Exec("DELETE FROM exchange_rates WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(200)
}
