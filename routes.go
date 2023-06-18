package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func startListening(app *fiber.App) {
	app.Get("/__health", func(c *fiber.Ctx) error {
		return c.SendString("{status: 'ok'}")
	})

	err := app.Listen(ApplicationPort)
	if err != nil {
		log.Fatalf("Error starting the app")
	}
}
