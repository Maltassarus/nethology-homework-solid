package order

import (
	"errors"
	"main/repository"
)

type Order struct {
	ID       int
	Customer string
	Products []string
	Total    float64
	Status   string
}

type OrderService struct {
	repo     repository.RepositoryWriter
	notifier notification.Notifier
}

func NewOrderService(repo repository.RepositoryWriter, notifier notification.Notifier) *OrderService {
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
