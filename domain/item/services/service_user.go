package services

import (
	"Campign/domain/item/helper"
	"Campign/domain/item/models"
	"Campign/domain/item/repositories"

	"gorm.io/gorm"
)

type userService struct {
	userRepo repositories.UserRepositories
}

type UserService interface {
	Register(user models.User) helper.Response
	Login(email string) (helper.Response, models.User)
	GetAllUser() helper.Response
}

func NewUserNewService(db *gorm.DB) UserService {
	return &userService{userRepo: repositories.NewUserReositories(db)}
}

func (service *userService) Register(user models.User) helper.Response {
	var respons helper.Response
	if err := service.userRepo.Register(user); err != nil {
		respons.Status = 500
		respons.Message = "Failed to register"
	} else {
		respons.Status = 200
		respons.Message = "Succes registered"
	}
	return respons
}

func (service *userService) Login(email string) (helper.Response, models.User) {
	var respons helper.Response
	data, err := service.userRepo.Login(email)
	if err != nil {
		respons.Status = 500
		respons.Message = "Failed to register"
	} else {
		respons.Status = 200
		respons.Message = "Succes registered"
		respons.Data = data
	}
	return respons, data
}

func (service *userService) GetAllUser() helper.Response {
	var respons helper.Response
	if data, err := service.userRepo.GetAll(); err != nil {
		respons.Status = 500
		respons.Message = "Failed to register"
	} else {
		respons.Status = 200
		respons.Message = "Succes registered"
		respons.Data = data
	}
	return respons
}
