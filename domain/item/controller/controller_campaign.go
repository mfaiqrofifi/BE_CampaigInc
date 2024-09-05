package controller

import (
	"Campign/domain/item/middleware"
	"Campign/domain/item/models"
	"Campign/domain/item/services"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CampaignController struct {
	campaignServices services.CampignServices
}

func NewControllerCampaign(db *gorm.DB) CampaignController {
	service := services.NewCampaignNewServices(db)
	controller := CampaignController{
		campaignServices: service,
	}
	return controller
}

func (controller CampaignController) Create(c echo.Context) error {
	type payload struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return err
	}
	newUUID := uuid.New()
	result := controller.campaignServices.Create(
		models.Campaigns{
			ID:          newUUID.String(),
			UserID:     middleware.GetUserNameFromToken(c),
			Title:       payloadValidator.Title,
			Description: payloadValidator.Description,
		},
	)
	return c.JSON(http.StatusOK, result)
}
func (controller CampaignController) Update(c echo.Context) error {
	fmt.Println("running")
	type payload struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return err
	}
	IdItem := c.Param("id")
	result := controller.campaignServices.Update(IdItem,
		models.Campaigns{
			Title:       payloadValidator.Title,
			Description: payloadValidator.Description,
		},
	)
	return c.JSON(http.StatusOK, result)
}
func (controller CampaignController) Delete(c echo.Context) error {
	IdItem := c.Param("id")
	fmt.Println(IdItem)
	result := controller.campaignServices.Delete(IdItem)
	return c.JSON(http.StatusOK, result)
}
func (controller CampaignController) GetAll(c echo.Context) error {
	id := middleware.GetUserNameFromToken(c)
	result := controller.campaignServices.GetAll(id)
	return c.JSON(http.StatusOK, result)
}
func (controller CampaignController) GetById(c echo.Context) error {
	IdItem := c.QueryParam("id")
	result := controller.campaignServices.GetById(IdItem)
	return c.JSON(http.StatusOK, result)
}
