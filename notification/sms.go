package notification

import (
	"fmt"
	"main/order"
)

type SMSSender struct {
	phoneNumber string
}

func NewSMSSender(phoneNumber string) *SMSSender {
	return &SMSSender{
		phoneNumber: phoneNumber,
	}
}

func (s *SMSSender) Send(customer string, order *order.Order) error {
	fmt.Printf("Детали отправленного SMS уведомления заказа: ID=%d, Сумма=%.2f\n", order.ID, order.Total)
	return nil
}
