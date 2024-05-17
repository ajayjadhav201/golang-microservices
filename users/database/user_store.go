package database

import (
	"model"
)

type UserStore interface {
	CreateUser(user model.User) (*model.User, error)
	GetUsers() *[]model.User
	GetUserById(id string) (*model.User, error)
	UpdateUser(id string, user model.User) (*model.User, error)
	DeleteUser(id string) error
}

type userStore struct {
	Users []*model.User
}

func NewUserStore() UserStore {
	return &userStore{
		Users: []*model.User{},
	}
}

func (s *userStore) CreateUser(user model.User) (*model.User, error) {
	//
	return nil, nil
}

func (s *userStore) GetUsers() *[]model.User {
	//
	return nil
}

func (s *userStore) GetUserById(id string) (*model.User, error) {
	//
	return nil, nil
}

func (s *userStore) UpdateUser(id string, user model.User) (*model.User, error) {
	//
	return nil, nil
}

func (s *userStore) DeleteUser(id string) error {
	//
	return nil
}
