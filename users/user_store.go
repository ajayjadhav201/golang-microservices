package main

import (
	"model"
)

type UserStore interface {
	CreateUser()
	GetUserById()
	GetUsers()
	UpdateUser()
	DeleteUser()
}

type userStore struct {
	Users []*model.User
}

func NewUserStore() UserStore {
	return &userStore{
		Users: []*model.User{},
	}
}

func (s *userStore) CreateUser() {
	//
}

func (s *userStore) GetUsers() {
	//
}

func (s *userStore) GetUserById() {
	//
}

func (s *userStore) UpdateUser() {
	//
}

func (s *userStore) DeleteUser() {
	//
}
