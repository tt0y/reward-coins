package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	app := fiber.New()
	db := getDb()
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

func healthCheck(c *fiber.Ctx) {
	c.Send("status: ok")
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
