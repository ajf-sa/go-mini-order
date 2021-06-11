package order

import (
	"github.com/alfuhigi/micro-order-api/pkg/order/item"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ClientID    string
	UserID      uint
	OrderItems  []*item.OrderItem
	OrderStatus []*OrderStatus
	Address     string
	Phone       string
	Email       string
}

type OrderStatus struct {
	gorm.Model
	OrderID     uint
	Name        string
	Description string
}
