package service

import (
	"hactive/assigment-2/config"
	"hactive/assigment-2/entity"
	"hactive/assigment-2/interfaces"
)

type Service struct {
	conf config.Config
	repo interfaces.OrderRepository
}

func NewOrderService(conf config.Config, repo interfaces.OrderRepository) interfaces.OrderService {
	return &Service{
		conf: conf,
		repo: repo,
	}
}

func (s *Service) CreateOrder(order *entity.Order) error {
	return s.repo.CreateOrder(order)
}

func (s *Service) GetOrders() ([]entity.Order, error) {
	return s.repo.GetOrders()
}

func (s *Service) UpdateOrder(id uint, order *entity.Order) error {
	return s.repo.UpdateOrder(id, order)
}

func (s *Service) DeleteOrder(id uint) error {
	return s.repo.DeleteOrder(id)
}
