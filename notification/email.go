package notification

import (
	"fmt"
	"main/order"
)

type EmailSender struct{}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

func (e *EmailSender) Send(customer string, order *order.Order) error {
	fmt.Printf("Детали отправленного заказа: ID=%d, Сумма=%.2f, Статус=%s\n",
		order.ID, order.Total, order.Status)
	return nil
}
