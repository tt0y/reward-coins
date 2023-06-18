package main

import (
	"github.com/gofiber/fiber/v2"
	"reward-coins-api/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Reward-coin API")
	})

	app.Get("/__health", func(c *fiber.Ctx) error {
		return c.SendString("{status: 'ok'}")
	})

	// transaction-type api
	app.Get("/api/v1/transaction-types", handlers.ListTransactionTypes)
	app.Get("/api/v1/transaction-types/:id", handlers.GetTransactionType)
	app.Post("/api/v1/transaction-types", handlers.AddTransactionType)
	app.Put("/api/v1/transaction-types", handlers.UpdateTransactionType)
	app.Delete("/api/v1/transaction-types/:id", handlers.RemoveTransactionType)
}
