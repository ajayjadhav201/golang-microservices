package database

import (
	"user-service/model"

	"golang-microservices/common"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProductStore interface {
	CreateUser(user model.User) (*model.User, error)
	GetUsers() *[]model.User
	GetUserById(id string) (*model.User, error)
	UpdateUser(id string, user model.User) (*model.User, error)
	DeleteUser(id string) error
}

type productStore struct {
	Users []*model.User
	DB    *gorm.DB
}

func NewProductStore() ProductStore {
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
	return &productStore{
		Users: []*model.User{},
		DB:    db,
	}
}

func (s *productStore) CreateUser(user model.User) (*model.User, error) {
	//
	return nil, nil
}

func (s *productStore) GetUsers() *[]model.User {
	//
	return nil
}

func (s *productStore) GetUserById(id string) (*model.User, error) {
	//
	return nil, nil
}

func (s *productStore) UpdateUser(id string, user model.User) (*model.User, error) {
	//
	return nil, nil
}

func (s *productStore) DeleteUser(id string) error {
	//
	return nil
}
