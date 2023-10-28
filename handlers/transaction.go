package handlers

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListTransactions(c *fiber.Ctx) error {
	rows, err := config.Database.Query("SELECT id, status, transaction_type_id, coin_type_id, amount, user_id_from, user_id_to, product_id FROM transactions")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.Status, &transaction.TransactionTypeId, &transaction.CoinTypeId, &transaction.Amount, &transaction.UserIdFrom, &transaction.UserIdTo, &transaction.ProductId)
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, transaction)
	}

	return c.Status(200).JSON(transactions)
}

func GetTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	var transaction models.Transaction

	err := config.Database.QueryRow("SELECT id, status, transaction_type_id, coin_type_id, amount, user_id_from, user_id_to, product_id FROM transactions WHERE id = ?", id).Scan(&transaction.ID, &transaction.Status, &transaction.TransactionTypeId, &transaction.CoinTypeId, &transaction.Amount, &transaction.UserIdFrom, &transaction.UserIdTo, &transaction.ProductId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.SendStatus(404)
		}
		log.Fatal(err)
	}

	return c.Status(200).JSON(&transaction)
}

func AddTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("INSERT INTO transactions (status, transaction_type_id, coin_type_id, amount, user_id_from, user_id_to, product_id) VALUES (?, ?, ?, ?, ?, ?, ?)",
		transaction.Status, transaction.TransactionTypeId, transaction.CoinTypeId, transaction.Amount, transaction.UserIdFrom, transaction.UserIdTo, transaction.ProductId)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(201).JSON(transaction)
}

func UpdateTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)
	id := c.Params("id")

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("UPDATE transactions SET status = ?, transaction_type_id = ?, coin_type_id = ?, amount = ?, user_id_from = ?, user_id_to = ?, product_id = ? WHERE id = ?",
		transaction.Status, transaction.TransactionTypeId, transaction.CoinTypeId, transaction.Amount, transaction.UserIdFrom, transaction.UserIdTo, transaction.ProductId, id)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(transaction)
}

func RemoveTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.Database.Exec("DELETE FROM transactions WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(200)
}
