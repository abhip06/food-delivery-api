package models

import (
	"time"
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
	ID                string       `json:"id" gorm:"type:varchar(255);primaryKey"`
	ShippingInfoRefer int          `json:"shipping_info"`
	ShippingInfo      ShippingInfo `gorm:"foreignkey:ShippingInfoRefer"`
	ItemRefer         string       `json:"item_id"`
	OrderItems        Item         `gorm:"foreignkey:ItemRefer"`
	UserRefer         string       `json:"user_id"`
	User              User         `gorm:"foreignkey:UserRefer"`
	ItemsPrice        uint         `json:"itemsprice"`
	TaxPrice          uint         `json:"taxprice"`
	DeliveryCharges   *uint        `json:"delivery_charges" gorm:"default:0"`
	TotalPrice        uint         `json:"totalprice"`
	OrderStatus       string       `json:"orderstatus" gorm:"default:Processing"`
	DeliveredAt       time.Time
	OrderedAt         time.Time
}
