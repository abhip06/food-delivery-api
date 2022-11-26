package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID       string `json:"id" gorm:"type:varchar(255);primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	Role     string `json:"role" gorm:"default:user"`
}

// Items Model
type Item struct {
	ID          string   `json:"id" gorm:"type:varchar(255);primaryKey"`
	Name        string   `json:"name"`
	Price       uint     `json:"price"`
	ShopAddress string   `json:"shop_address"`
	ItemImage   string   `json:"item_image"`
	Ratings     *float64 `json:"ratings" gorm:"default:0"`
	Category    string   `json:"category"`
	IsAvailable bool     `json:"is_available" gorm:"default:true"`
	IsFeatured  bool     `json:"is_featured" gorm:"default:false"`
	UserRefer   string   `json:"user_id"`
	User        User     `gorm:"foreignkey:UserRefer"`
	CreatedAt   time.Time
}

type ShippingInfo struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Address string `json:"address"`
	City    string `json:"city"`
	PinCode uint   `json:"pincode"`
	PhoneNo uint   `json:"phoneno"`
}

// Order Model
type Order struct {
	ID              primitive.ObjectID `json:"_id" bson"_id"`
	ShippingInfo    ShippingInfo       `json:"shipping_info" bson:"shipping_info"`
	OrderItems      []Item             `json:"order_items" bson:"order_items"`
	User            User               `json:"user_refer" bson:"user_refer"`
	ItemsPrice      uint               `json:"itemsprice"`
	TaxPrice        uint               `json:"taxprice"`
	DeliveryCharges *uint              `json:"delivery_charges" bson:"delivery_charges"`
	TotalPrice      uint               `json:"totalprice"`
	OrderStatus     string             `json:"order_status" bson:"order_status"`
	DeliveredAt     time.Time          `json:"delivered_at" bson:"delivered_at"`
	OrderedAt       time.Time          `json:"ordered_at" bson:"ordered_at"`
}
