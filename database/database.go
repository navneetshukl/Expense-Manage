package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/navneetshukl/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

// ! ConnectToDatabase will connect to database
func ConnectToDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_CONNECTION_STRING")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database ", err)
		return nil, err
	}
	fmt.Println("Connected to Database")
	return DB, nil
}

// ! MigrateDatabase will migrate the database
func MigrateDatabase() {
	DB, err := ConnectToDatabase()
	if err != nil {
		log.Fatal("There is error connecting to database ", err)
		return
	}
	DB.AutoMigrate(&models.Grocery{},&models.User{})
}
