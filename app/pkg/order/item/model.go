package item

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID     uint
	ProductID   uint
	Sku         string
	ProductName string
	Qty         uint
	Price       float32
	TotalPrice  float32 // Qty * Price
}
