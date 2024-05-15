package main

import "time"

type UserStore interface {
	CreateUser()
	GetUserById()
	GetUsers()
	UpdateUser()
	DeleteUser()
}

type User struct {
	UserID       string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	MobileNumber string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type userStore struct {
	Users []*User
}

func NewUserStore() UserStore {
	return &userStore{
		Users: []*User{},
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
