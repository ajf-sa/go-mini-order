package order

type ORepository interface {
	Create(o *Order) (*Order, error)
	FindOneByID(OrderID uint) (*Order, error)
	FindAllByUserID(userID uint) ([]*Order, error)
	FindAllByClientID(clientID uint) ([]*Order, error)
}
