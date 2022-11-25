package routes

import (
	"github.com/abhip06/food-delivery-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// User Routes
	app.Post("/api/v1/register", controllers.Register)
	app.Post("/api/v1/login", controllers.Login)
	app.Get("/api/v1/user", controllers.User)
	app.Get("/api/v1/logout", controllers.Logout)

	// Item Routes
	app.Get("/api/v1/items", controllers.GetAllItems)
	app.Get("/api/v1/item/:id", controllers.GetItem)
	app.Get("/api/v1/featured", controllers.GetFeaturedItem)    // Get Featured Item
	app.Get("/api/v1/search/name", controllers.SearchByName)         // Search by name
	app.Get("/api/v1/search/category", controllers.SearchByCategory) // Search by category
	app.Post("/api/v1/admin/item", controllers.CreateItem)           // Admin
	app.Put("/api/v1/admin/item/:id", controllers.UpdateItem)        // Admin
	app.Delete("/api/v1/admin/item/:id", controllers.DeleteItem)     // Admin

	// Order Routes
	app.Get("/api/v1/order/:id", controllers.GetOrder)
	app.Post("/api/v1/order", controllers.Order)
	app.Delete("/api/v1/order/:id", controllers.CancelOrder)
	app.Get("/api/v1/admin/orders", controllers.GetAllOrders)         // Admin
	app.Put("/api/v1/admin/order/:id", controllers.UpdateOrderStatus) // Admin
}
