package order

type OService interface {
	CreateOrder(o *Order) (*Order, error)
	FindOneOrderByID(orderID uint) (*Order, error)
	FindAllOrderByUserID(userID uint) ([]*Order, error)
	FindAllOrderByClientID(clientID uint) ([]*Order, error)
	FindAllOrder() ([]*Order, error)
}

type oService struct {
	repo ORepository
}

func NewOService(repo ORepository) OService {
	return &oService{repo}
}
func (s *oService) CreateOrder(o *Order) (*Order, error) {
	set, err := s.repo.Create(o)
	return set, err

}
func (s *oService) FindOneOrderByID(orderID uint) (*Order, error) {
	return nil, nil

}
func (s *oService) FindAllOrderByUserID(userID uint) ([]*Order, error) {
	get, err := s.repo.FindAllByUserID(userID)
	return get, err

}
func (s *oService) FindAllOrderByClientID(clientID uint) ([]*Order, error) {
	get, err := s.repo.FindAllByClientID(clientID)
	return get, err

}
func (s *oService) FindAllOrder() ([]*Order, error) {
	return nil, nil
}
