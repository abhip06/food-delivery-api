package main

import (
	"os"

	"github.com/abhip06/food-delivery-api/database"
	"github.com/abhip06/food-delivery-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.SetupRoutes(app)

	app.Listen(":" + port)
}
