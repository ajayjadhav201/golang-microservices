package main

import "context"

type Service struct {
	Store OrderStore
}

func NewService(Store OrderStore) *Service {
	return &Service{
		Store: Store,
	}
}

func (s *Service) CreateOrder(context.Context) {

}
