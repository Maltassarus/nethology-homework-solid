package order

import (
	"errors"
)

type RepositoryWriter interface {
	SaveOrder(order *Order) error
}

type Notifier interface {
	Send(customer string, order *Order) error
}

type OrderService struct {
	repo     RepositoryWriter
	notifier Notifier
}

func NewOrderService(repo RepositoryWriter, notifier Notifier) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	if customer == "" {
		return errors.New("customer name cannot be empty")
	}
	if len(products) == 0 {
		return errors.New("products list cannot be empty")
	}
	if total <= 0 {
		return errors.New("total must be positive")
	}

	order := &Order{
		Customer: customer,
		Products: products,
		Total:    total,
		Status:   "pending",
	}

	if err := s.repo.SaveOrder(order); err != nil {
		return err
	}

	if err := s.notifier.Send(customer, order); err != nil {
		return err
	}

	return nil
}
