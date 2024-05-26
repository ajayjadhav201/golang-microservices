package database

import (
	"model"

	"github.com/ajayjadhav201/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	UserAlreadyRegistered = "User is already registered."
	UserNotFound          = "User not found."
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
	DB    *gorm.DB
}

func NewUserStore() UserStore {
	//
	if common.Test {
		return &userStore{
			Users: []*model.User{},
		}
	}
	//
	//
	host := common.EnvString("DB_HOST", "localhost")
	port := common.Atoi(common.EnvString("DB_PORT", "5432"), 5432)
	user := common.EnvString("DB_USER", "user")
	password := common.EnvString("DB_PASSWORD", "password")
	dbname := common.EnvString("DB_NAME", "database")

	//
	dsn := common.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	common.Panic(err)

	// defer db.Close()
	db.AutoMigrate(&model.User{})

	newUser := model.User{
		FirstName: "Ajay",
		LastName:  "Jadhav",
		Email:     "jadhavaj201@gmail.com",
		Password:  common.Encrypt("Jadhavaj20@"),
	}
	tx := db.Create(&newUser)
	if tx.Error != nil {
		common.Println("transaction error ", tx.Error.Error())
	}

	common.Println("transaction", tx.Statement)

	common.Println("connected to database")
	return &userStore{
		Users: []*model.User{},
		DB:    db,
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
