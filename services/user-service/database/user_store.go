package database

import (
	"user-service/model"

	"golang-microservices/common"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	UserAlreadyRegistered = "User is already registered."
	UserNotFound          = "User not found."
)

type UserStore interface {
	CreateUser(user *model.User) (string, error)
	GetUsers() []*model.User
	GetUserById(id string) (*model.User, error)
	GetUserByEmailorMobile(emailorMobile string) (*model.User, error)
	UpdateUser(id string, user *model.User) (*model.User, error)
	DeleteUser(id string) error
}

type userStore struct {
	Users []*model.User
	DB    *gorm.DB
}

func NewUserStore() *userStore {
	//
	host := common.EnvString("DB_HOST", "localhost")
	port := common.Atoi(common.EnvString("DB_PORT", "5432"), 5432)
	user := common.EnvString("DB_USER", "user")
	password := common.EnvString("DB_PASSWORD", "password")
	dbname := common.EnvString("DB_NAME", "database")

	common.Println(user, " connection to db with password: ", password)
	//
	dsn := common.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	common.Panic(err)

	// defer db.Close()
	db.AutoMigrate(&model.User{})

	common.Println("connected to database")
	return &userStore{
		Users: []*model.User{},
		DB:    db,
	}
}

func (s *userStore) CreateUser(user *model.User) (string, error) {
	//
	common.Println("ajaj adding user to database")
	tx := s.DB.Create(user)
	if tx.Error != nil {
		return "error", common.Error(tx.Error.Error())
	}
	common.Println("ajaj user added to db rows affected: ", tx.RowsAffected)

	return "user created", nil
}

func (s *userStore) GetUsers() []*model.User {
	//
	return nil
}

func (s *userStore) GetUserById(email string) (*model.User, error) {
	//
	return nil, nil
}

func (s *userStore) UpdateUser(id string, user *model.User) (*model.User, error) {
	//
	return nil, nil
}

func (s *userStore) DeleteUser(id string) error {
	//
	return nil
}
