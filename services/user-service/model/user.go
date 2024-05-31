package model

import (
	"golang-microservices/common"
	"time"
)

/*
type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName    string `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty" validate:"required"`
	LastName     string `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty" validate:"required"`
	Email        string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty" validate:"required,email"`
	Password     string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty" validate:"required"`
	MobileNumber string `protobuf:"bytes,5,opt,name=MobileNumber,proto3" json:"MobileNumber,omitempty" validate:"required,mobile"`
	Address      string `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty" `
	City         string `protobuf:"bytes,7,opt,name=City,proto3" json:"City,omitempty"`
	State        string `protobuf:"bytes,8,opt,name=State,proto3" json:"State,omitempty"`
	Country      string `protobuf:"bytes,9,opt,name=Country,proto3" json:"Country,omitempty"`
	ZipCode      string `protobuf:"bytes,10,opt,name=ZipCode,proto3" json:"ZipCode,omitempty"`
	ProfileImage string `protobuf:"bytes,11,opt,name=ProfileImage,proto3" json:"ProfileImage,omitempty"`
}
*/

type User struct {
	ID              int64     `gorm:"primaryKey" json:"ID"`
	FirstName       string    `json:"FirstName,omitempty"`
	LastName        string    `json:"LastName,omitempty"`
	ProfileImage    string    `json:"ProfileImage,omitempty"`
	Email           string    `gorm:"unique" json:"Email,omitempty"`
	Password        string    `json:"Password,omitempty"`
	MobileNumber    string    `gorm:"unique" json:"MobileNumber,omitempty"`
	Address         string    `json:"Address,omitempty"`
	ShippingAddress string    `json:"ShippingAddress,omitempty"`
	Token           string    `json:"Token,omitempty"`
	CreatedAt       time.Time `json:"CreatedAt,omitempty"`
	UpdatedAt       time.Time `json:"UpdatedAt,omitempty"`
	// DeletedAt       *time.Time `gorm:"index"`
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

// if Settings string `gorm:"type:varchar(100)"`
// func (u *User) BeforeCreate(scope *gorm.Scope) error {
// 	settingsJSON, err := json.Marshal(u.Settings)
// 	if err != nil {
// 		return err
// 	}
// 	scope.SetColumn("Settings", settingsJSON)
// 	return nil
// }
// func (u *User) AfterFind(scope *gorm.Scope) error {
// 	settingsJSON := scope.DB().Get("Settings")
// 	if settingsJSON != nil {
// 		return json.Unmarshal([]byte(settingsJSON.(string)), &u.Settings)
// 	}
// 	return nil
// }

//  if Settings string `gorm:"type:text"`
// func (u *User) GetSettings() (map[string]string, error) {
//     var settings map[string]string
//     err := json.Unmarshal([]byte(u.Settings), &settings)
//     return settings, err
// }
// func (u *User) SetSettings(settings map[string]string) error {
//     settingsJSON, err := json.Marshal(settings)
//     if err != nil {
//         return err
//     }
//     u.Settings = string(settingsJSON)
//     return nil
// }
