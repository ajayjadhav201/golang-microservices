package model

import "time"

type User struct {
	ID              string
	FirstName       string
	LastName        string
	Email           string
	Password        string
	MobileNumber    string
	Address         string
	ShippingAddress string
	Token           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
