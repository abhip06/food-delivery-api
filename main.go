package main

import (
	"os"

	"github.com/abhip06/food-delivery-api/database"
	"github.com/abhip06/food-delivery-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routes.SetupRoutes(app)

	app.Listen(":" + port)
}
