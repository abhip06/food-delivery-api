package controllers

import (
	"github.com/abhip06/food-delivery-api/database"
	"github.com/abhip06/food-delivery-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateItem(c *fiber.Ctx) error {
	item := new(models.Item)

	if err := c.BodyParser(&item); err != nil {
		return c.JSON(err)
	}

	// var user *models.User
	// database.DB.Find(&user, "id = ?", item.UserRefer)

	// fmt.Println(user)

	item.ID = uuid.New().String()

	database.DB.Save(&item)

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"message": "Item added successfully.",
		"item":    item,
	})
}

func GetAllItems(c *fiber.Ctx) error {
	items := []models.Item{}

	database.DB.Find(&items)

	if len(items) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": "false",
			"message": "No Items Found.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"items":   items,
	})
}

func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item

	result := database.DB.Find(&item, "id = ?", id)

	if result.Error != nil {
		return c.Status(200).JSON(fiber.Map{
			"success": "false",
			"message": "No Item Found",
		})
	}

	return c.Status(200).JSON(item)
}

func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item

	result := database.DB.First(&item, "id = ?", id)

	if result.Error != nil {
		return c.Status(200).JSON(fiber.Map{
			"success": "false",
			"message": "No Item Found",
		})
	}

	database.DB.Delete(&item, "id = ?", id)

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"message": "Item deleted successfully.",
	})
}

func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")

	item := new(models.Item)

	result := database.DB.First(&item, "id = ?", id)

	if result.Error != nil {
		return c.Status(200).JSON(fiber.Map{
			"success": "false",
			"message": "No Item Found",
		})
	}

	type UpdateItem struct {
		Name        string `json:"name"`
		Price       uint   `json:"price"`
		ShopAddress string `json:"shop_address"`
		Category    string `json:"category"`
		IsAvailable bool   `json:"is_available" gorm:"default:true"`
	}

	var updateData UpdateItem

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(200).JSON(fiber.Map{
			"success": "false",
			"message": "Invalid Input fields.",
		})
	}

	item.Name = updateData.Name
	item.Price = updateData.Price
	item.ShopAddress = updateData.ShopAddress
	item.Category = updateData.Category
	item.IsAvailable = updateData.IsAvailable

	database.DB.Save(&item)

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"message": "Item updated successfully.",
		"item":    item,
	})
}

// Search Item by Name
func SearchByName(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	var items []models.Item

	database.DB.Where("name = ?", keyword).Find(&items)

	if len(items) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": "false",
			"message": "No Items Found.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"items":   items,
	})
}

// Search Item by Category
func SearchByCategory(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	var items []models.Item

	database.DB.Where("category = ?", keyword).Find(&items)

	if len(items) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": "false",
			"message": "No Items Found.",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": "true",
		"items":   items,
	})
}
