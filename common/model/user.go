package model

import "time"

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
