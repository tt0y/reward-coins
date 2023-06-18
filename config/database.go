package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"reward-coins-api/models"
)

var Database *gorm.DB

func getDataSource() string {
	dbName := goDotEnvVariable("MYSQL_DATABASE")
	user := goDotEnvVariable("MYSQL_ROOT_USER")
	password := goDotEnvVariable("MYSQL_ROOT_PASSWORD")

	return fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbName)
}
func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(getDataSource()), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&models.TransactionType{})

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
