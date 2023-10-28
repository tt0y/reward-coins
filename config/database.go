package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Database *sql.DB

func getDataSource() string {
	dbName := goDotEnvVariable("MYSQL_DATABASE")
	user := goDotEnvVariable("MYSQL_USER")
	password := goDotEnvVariable("MYSQL_PASSWORD")

	log.Println(dbName)
	log.Println(user)
	log.Println(password)

	return fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbName)
}

func Connect() error {
	var err error

	Database, err = sql.Open("mysql", getDataSource())

	if err != nil {
		log.Fatal(err)
	}

	err = Database.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return nil
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
