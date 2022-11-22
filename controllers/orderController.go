package controllers

import (
	"github.com/abhip06/food-delivery-api/database"
	"github.com/abhip06/food-delivery-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Order(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.JSON(err)
	}

	order.ID = uuid.New().String()

	database.DB.Save(&order)

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"order":   order,
	})
}

func GetAllOrders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Find(&orders)

	if len(orders) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": "false",
			"message": "No Orders Found.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"orders":  orders,
	})
}

// Get order by id
func GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order

	result := database.DB.Find(&order, "id = ?", id)

	if result.Error != nil {
		return c.Status(200).JSON(fiber.Map{
			"success": "false",
			"message": "No Order Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"order":   order,
	})
}

func CancelOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order

	result := database.DB.First(&order, "id = ?", id)

	if result.Error != nil {
		return c.Status(200).JSON(fiber.Map{
			"success": "false",
			"message": "No Order Found",
		})
	}

	database.DB.Delete(&order, "id = ?", id)

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"message": "Order Cancelled.",
	})
}

func UpdateOrderStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order

	result := database.DB.First(&order, "id = ?", id)

	if result.Error != nil {
		return c.Status(200).JSON(fiber.Map{
			"success": "false",
			"message": "No Order Found",
		})
	}

	type UpdateOrder struct {
		OrderStatus string `json:"orderstatus" gorm:"default:Processing"`
	}

	var updateData UpdateOrder

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(200).JSON(fiber.Map{
			"success": "false",
			"message": "Invalid Input fields.",
		})
	}

	order.OrderStatus = updateData.OrderStatus

	database.DB.Save(&order)

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"message": "Order status updated.",
		"order":   order,
	})
}
