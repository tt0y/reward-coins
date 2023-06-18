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

	// user api
	app.Get("/api/v1/user", handlers.ListUsers)
	app.Get("/api/v1/user/:id", handlers.GetUser)
	app.Post("/api/v1/user", handlers.AddUser)
	app.Put("/api/v1/user/:id", handlers.UpdateUser)
	app.Delete("/api/v1/user/:id", handlers.RemoveUser)

	// product api
	app.Get("/api/v1/product", handlers.ListProducts)
	app.Get("/api/v1/product/:id", handlers.GetProduct)
	app.Post("/api/v1/product", handlers.AddProduct)
	app.Put("/api/v1/product/:id", handlers.UpdateProduct)
	app.Delete("/api/v1/product/:id", handlers.RemoveProduct)

	// user balance api
	app.Get("/api/v1/balance", handlers.ListUserBalances)
	app.Get("/api/v1/balance/:id", handlers.GetUserBalance)
	app.Post("/api/v1/balance", handlers.AddUserBalance)
	app.Put("/api/v1/balance/:id", handlers.UpdateUserBalance)
	app.Delete("/api/v1/balance/:id", handlers.RemoveUserBalance)
}
