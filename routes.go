package main

import (
	"github.com/gofiber/fiber"
	"log"
)

func startListening(app *fiber.App) {
	app.Get("/__health", healthCheck)

	err := app.Listen(ApplicationPort)
	if err != nil {
		log.Fatalf("Error starting the app")
	}
}
