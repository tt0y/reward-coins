package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New()
	db := getDb()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Fiber API")
	})

	startListening(app)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing the db connection")
		}
	}(db)
}

func getDataSource() string {
	dbName := goDotEnvVariable("MYSQL_DATABASE")
	user := goDotEnvVariable("MYSQL_USER")
	password := goDotEnvVariable("MYSQL_PASSWORD")

	return fmt.Sprintf("%s:%s@/%s", user, password, dbName)
}

func getDb() *sql.DB {
	db, err := sql.Open("mysql", getDataSource())

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * ConnectionsMaxLifetimeMin)
	db.SetMaxOpenConns(MaxOpenConnections)
	db.SetMaxIdleConns(MaxIdleConnections)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return db
}

// use godot package to load/read the .env file and return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
