package item

type IService interface {
	CreateItem(i *OrderItem) (*OrderItem, error)
	FindAllItemByOrderID(orderID uint) ([]*OrderItem, error)
}

type iservice struct {
	repo IRepository
}

func NewIService(repo IRepository) IService {
	return &iservice{repo}
}

func (s *iservice) CreateItem(i *OrderItem) (*OrderItem, error) {
	set, err := s.repo.Create(i)
	return set, err

}

func (s *iservice) FindAllItemByOrderID(orderID uint) ([]*OrderItem, error) {
	get, err := s.repo.FindAllByOrderID(orderID)
	return get, err
}
