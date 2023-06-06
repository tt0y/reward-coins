package main

import (
	"github.com/gofiber/fiber"
	"log"
)

func startListening(app *fiber.App) {
	err := app.Listen(ApplicationPort)
	if err != nil {
		log.Fatalf("Error starting the app")
	}

	app.Get("/__health", healthCheck)
}
