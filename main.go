package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"reward-coins-api/config"
)

func main() {
	err := config.Connect()
	if err != nil {
		log.Printf("Error connecting to database")
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001", // react app
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	SetupRoutes(app)

	err = app.Listen(ApplicationPort)
	if err != nil {
		log.Fatalf("Error starting the app")
	}
}
