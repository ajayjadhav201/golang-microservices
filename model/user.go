package model

import (
	"time"

	"github.com/ajayjadhav201/common"
)

type User struct {
	ID              uint   `gorm:"primaryKey" json:"omitempty"`
	UserID          string `json:"ID,omitempty"`
	FirstName       string `json:"FirstName,omitempty"`
	LastName        string `json:"LastName,omitempty"`
	ProfileImage    string `json:"ProfileImage,omitempty"`
	Email           string `json:"Email,omitempty"`
	Password        string `json:"Password,omitempty"`
	MobileNumber    string `json:"MobileNumber,omitempty"`
	Address         string `json:"Address,omitempty"`
	ShippingAddress string `json:"ShippingAddress,omitempty"`
	Token           string `json:"Token,omitempty"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `gorm:"index"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) CopyWith(newUser *User) error {
	b, e := common.MarshalJSON(newUser)
	if e != nil {
		return e
	}
	e = common.UnmarshalJSON(b, u)
	if e != nil {
		return e
	}
	return nil
}
