package repository

import "example/solid/order"

type RepositoryWriter interface {
	SaveOrder(order *order.Order) error
}

type RepositoryInitializer interface {
	Init() error
}

type CombinedRepository interface {
	RepositoryWriter
	RepositoryInitializer
}
