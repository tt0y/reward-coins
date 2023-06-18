package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"reward-coins-api/config"
)

func main() {
	err := config.Connect()
	if err != nil {
		log.Printf("Error connecting to database")
	}
	app := *fiber.New()
	SetupRoutes(&app)

	err = app.Listen(ApplicationPort)
	if err != nil {
		log.Fatalf("Error starting the app")
	}
}
