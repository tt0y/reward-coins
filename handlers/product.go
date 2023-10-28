package handlers

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListProducts(c *fiber.Ctx) error {
	rows, err := config.Database.Query("SELECT name, description, cost, coin_type_id, images, active FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Name, &product.Description, &product.Cost, &product.CoinTypeId, &product.Images, &product.Active)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	return c.Status(200).JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	err := config.Database.QueryRow("SELECT name, description, cost, coin_type_id, images, active FROM products WHERE id = ?", id).Scan(&product.Name, &product.Description, &product.Cost, &product.CoinTypeId, &product.Images, &product.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.SendStatus(404)
		}
		log.Fatal(err)
	}

	return c.Status(200).JSON(&product)
}

func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("INSERT INTO products (name, description, cost, coin_type_id, images, active) VALUES (?, ?, ?, ?, ?, ?)",
		product.Name, product.Description, product.Cost, product.CoinTypeId, product.Images, product.Active)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(201).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	id := c.Params("id")

	if err := c.BodyParser(product); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("UPDATE products SET name = ?, description = ?, cost = ?, coin_type_id = ?, images = ?, active = ? WHERE id = ?",
		product.Name, product.Description, product.Cost, product.CoinTypeId, product.Images, product.Active, id)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(product)
}

func RemoveProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.Database.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(200)
}
