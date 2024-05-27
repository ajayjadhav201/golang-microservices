package model

import (
	"time"

	"github.com/ajayjadhav201/common"
)

type User struct {
	ID              uint64 `gorm:"primaryKey" json:"ID"`
	FirstName       string `json:"FirstName,omitempty"`
	LastName        string `json:"LastName,omitempty"`
	ProfileImage    string `json:"ProfileImage,omitempty"`
	Email           string `gorm:"unique" json:"Email,omitempty"`
	Password        string `json:"Password,omitempty"`
	MobileNumber    string `gorm:"unique" json:"MobileNumber,omitempty"`
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
