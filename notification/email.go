package notification

import (
	"fmt"

	"solid/order"
)

type EmailSender struct{}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

func (e *EmailSender) Send(customer string, order *order.Order) error {
	fmt.Printf("Уведомление %s\n", customer)
	fmt.Printf("Заказ #%d, сумма %.2f руб.\n", order.ID, order.Total)
	fmt.Printf("Товары: %v\n", order.Products)
	fmt.Printf("Статус: %s\n", order.Status)
	return nil
}
