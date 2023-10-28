package handlers

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func ListUsers(c *fiber.Ctx) error {
	rows, err := config.Database.Query("SELECT id, name, email, phone, password, is_admin FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.IsAdmin)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	err := config.Database.QueryRow("SELECT id, name, email, phone, password, is_admin FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.IsAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.SendStatus(404)
		}
		log.Fatal(err)
	}

	return c.Status(200).JSON(&user)
}

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("INSERT INTO users (name, email, phone, password, is_admin) VALUES (?, ?, ?, ?, ?)",
		user.Name, user.Email, user.Phone, user.Password, user.IsAdmin)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(models.User)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, err := config.Database.Exec("UPDATE users SET name = ?, email = ?, phone = ?, password = ?, is_admin = ? WHERE id = ?",
		user.Name, user.Email, user.Phone, user.Password, user.IsAdmin, id)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(user)
}

func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := config.Database.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	return c.SendStatus(200)
}
