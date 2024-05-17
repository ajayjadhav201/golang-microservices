package model

import "time"

type User struct {
	ID              string `json:"ID"`
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	ProfileImage    string `json:"ProfileImage"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	MobileNumber    string `json:"MobileNumber"`
	Address         string `json:"Address"`
	ShippingAddress string `json:"ShippingAddress"`
	Token           string `json:"Token"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
