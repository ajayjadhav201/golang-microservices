package main

import "context"

type OrderService interface {
	CreateOrder(context.Context)
}

type OrderStore interface {
	Create(context.Context) error
}
