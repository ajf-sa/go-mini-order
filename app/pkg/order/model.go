package order

import (
	"github.com/alfuhigi/micro-order-api/pkg/order/item"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ClientID       string
	UserID         uint
	OrderItems     []*item.OrderItem
	OrderStatus    []*OrderStatus
	PaymentOptions uint
	PickUp         string
	DropOff        string
	DeliveryFees   float32
	Notes          string
	Phone          string
	Email          string
	TotalPrice     float32 //SUM(...OrderItems.TotalPrice)
	// Payment_options
	// Address string
	// Promo_code
}

type OrderStatus struct {
	gorm.Model
	OrderID     uint
	Name        string
	Description string
}
