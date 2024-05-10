package main

import (
	"context"
)

type Store struct {
	// add here mongodb instance
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(context.Context) error {
	//
	return nil
}
