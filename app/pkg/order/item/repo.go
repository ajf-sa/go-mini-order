package item

type IRepository interface {
	Create(i *OrderItem) (*OrderItem, error)
	FindAllByOrderID(orderID uint) ([]*OrderItem, error)
}
