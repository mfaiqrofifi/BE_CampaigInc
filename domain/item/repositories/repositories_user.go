package repositories

import (
	"Campign/domain/item/models"

	"gorm.io/gorm"
)

type dbUser struct {
	Conn *gorm.DB
}

type UserRepositories interface {
	Register(user models.User) error
	Login(email string) (models.User, error)
	GetAll() ([]models.User, error)
}

func NewUserReositories(Conn *gorm.DB) UserRepositories {
	return &dbUser{Conn: Conn}
}

func (db dbUser) Register(user models.User) error {
	return db.Conn.Create(&user).Error
}

func (db dbUser) Login(email string) (models.User, error) {
	var data models.User
	result := db.Conn.Where("email=?", email).First(&data)
	return data, result.Error
}

func (db dbUser) GetAll() ([]models.User, error) {
	var data []models.User
	err := db.Conn.Preload("Campaign").Find(&data)
	return data, err.Error
}
