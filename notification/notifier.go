package notification

import "main/order"

type Notifier interface {
	Send(customer string, order *order.Order) error
}
