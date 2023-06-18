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
	app.Get("/api/v1/transaction-type", handlers.ListTransactionTypes)
	app.Get("/api/v1/transaction-type/:id", handlers.GetTransactionType)
	app.Post("/api/v1/transaction-type", handlers.AddTransactionType)
	app.Put("/api/v1/transaction-type/:id", handlers.UpdateTransactionType)
	app.Delete("/api/v1/transaction-type/:id", handlers.RemoveTransactionType)

	// coin-type api
	app.Get("/api/v1/coin-type", handlers.ListCoinTypes)
	app.Get("/api/v1/coin-type/:id", handlers.GetCoinType)
	app.Post("/api/v1/coin-type", handlers.AddCoinType)
	app.Put("/api/v1/coin-type/:id", handlers.UpdateCoinType)
	app.Delete("/api/v1/coin-type/:id", handlers.RemoveCoinType)

	// exchange-rate api
	app.Get("/api/v1/exchange-rate", handlers.ListExchangeRates)
	app.Get("/api/v1/exchange-rate/:id", handlers.GetExchangeRate)
	app.Post("/api/v1/exchange-rate", handlers.AddExchangeRate)
	app.Put("/api/v1/exchange-rate/:id", handlers.UpdateExchangeRate)
	app.Delete("/api/v1/exchange-rate/:id", handlers.RemoveExchangeRate)
}
