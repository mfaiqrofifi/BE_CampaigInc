package controller

import (
	"Campign/domain/item/helper"
	"Campign/domain/item/middleware"
	"Campign/domain/item/models"
	"Campign/domain/item/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	userService services.UserService
}

func NewControllerUser(db *gorm.DB) UserController {
	service := services.NewUserNewService(db)
	controllers := UserController{
		userService: service,
	}
	return controllers
}

func (controller UserController) Register(c echo.Context) error {
	newUUID := uuid.New()
	type payload struct {
		ID       string `json:"id"`
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return err
	}
	err, password := helper.HashPassword(payloadValidator.Password)
	if err != nil {
		return err
	}
	result := controller.userService.Register(
		models.User{
			ID:       newUUID.String(),
			Name:     payloadValidator.Name,
			Email:    payloadValidator.Email,
			Password: password,
		},
	)
	return c.JSON(http.StatusOK, result)
}

func (controller UserController) GetAllUse(c echo.Context) error {
	result := controller.userService.GetAllUser()
	return c.JSON(http.StatusOK, result)
}

func (controller UserController) Login(c echo.Context) error {
	type payload struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return err
	}
	result, data := controller.userService.Login(payloadValidator.Email)
	if err := helper.CheckPassword(data.Password, payloadValidator.Password); err != nil {
		return err
	}
	token, err := middleware.CreateToken(data.ID, data.Email)
	if err != nil {
		return err
	}
	type newPayload struct {
		Nama  string `json:"Nama"`
		Email string `json:"email"`
		Token string `json:"token"`
	}
	result.Data = newPayload{
		Nama:  data.Name,
		Email: data.Email,
		Token: token,
	}
	return c.JSON(http.StatusOK, result)
}
