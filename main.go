package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	// godotenv package
	databaseName := goDotEnvVariable("MYSQL_DATABASE")
	fmt.Printf("godotenv : %s = %s \n", "MYSQL_DATABASE", databaseName)

	app.Get("/__health", healthCheck)

	err := app.Listen(3000)
	if err != nil {
		log.Fatalf("Error starting the app")
	}
}

func healthCheck(c *fiber.Ctx) {
	c.Send("status: ok")
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
