package notification

import (
	"fmt"

	"solid/order"
)

type SMSSender struct {
	phoneNumber string
}

func NewSMSSender(phoneNumber string) *SMSSender {
	return &SMSSender{
		phoneNumber: phoneNumber,
	}
}

// Send реализует интерфейс order.Notifier
func (s *SMSSender) Send(customer string, order *order.Order) error {
	fmt.Printf("SMS уведомление %s\n", customer)
	fmt.Printf("Номер телефона: %s\n", s.phoneNumber)
	fmt.Printf("Заказ #%d, сумма %.2f руб.\n", order.ID, order.Total)
	fmt.Printf("Товары: %v\n", order.Products)
	return nil
}
