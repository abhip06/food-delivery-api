package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/abhip06/food-delivery-api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func MongoDBConnect() *mongo.Client {

	db_uri := os.Getenv("MONGODB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(db_uri))
	if err != nil {
		log.Fatal("Enable to connect to mongodb. ERROR: ", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	return client
}

var Client *mongo.Client = MongoDBConnect()
