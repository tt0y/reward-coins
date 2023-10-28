package handlers

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListTransactionTypes(c *fiber.Ctx) error {
	rows, err := config.Database.Query("SELECT id, name FROM transaction_types")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var transactionTypes []models.TransactionType
	for rows.Next() {
		var transactionType models.TransactionType
		err := rows.Scan(&transactionType.ID, &transactionType.Name)
		if err != nil {
			log.Fatal(err)
		}
		transactionTypes = append(transactionTypes, transactionType)
	}

	return c.Status(200).JSON(transactionTypes)
}

func GetTransactionType(c *fiber.Ctx) error {
	id := c.Params("id")
	var transactionType models.TransactionType

	err := config.Database.QueryRow("SELECT id, name FROM transaction_types WHERE id = ?", id).Scan(&transactionType.ID, &transactionType.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.SendStatus(404)
		}
		log.Fatal(err)
	}

	return c.Status(200).JSON(&transactionType)
}

func AddTransactionType(c *fiber.Ctx) error {
	transactionType := new(models.TransactionType)

	if err := c.BodyParser(transactionType); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("INSERT INTO transaction_types (name) VALUES (?)", transactionType.Name)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(201).JSON(transactionType)
}

func UpdateTransactionType(c *fiber.Ctx) error {
	transactionType := new(models.TransactionType)
	id := c.Params("id")

	if err := c.BodyParser(transactionType); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("UPDATE transaction_types SET name = ? WHERE id = ?", transactionType.Name, id)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(transactionType)
}

func RemoveTransactionType(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.Database.Exec("DELETE FROM transaction_types WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(200)
}
