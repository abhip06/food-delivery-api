package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/abhip06/food-delivery-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URI")
	
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to Database")
	}


	DB = connection

	connection.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Order{},
		&models.ShippingInfo{},
	)
}
