package services

import (
	"Campign/domain/item/helper"
	"Campign/domain/item/models"
	"Campign/domain/item/repositories"
	"fmt"

	"gorm.io/gorm"
)

type campaignServices struct {
	campaignRepo repositories.CampaignRepositories
}

type CampignServices interface {
	Create(campaign models.Campaigns) helper.Response
	Update(id string, campaign models.Campaigns) helper.Response
	Delete(id string) helper.Response
	GetById(id string) helper.Response
	GetAll(id string) helper.Response
}

func NewCampaignNewServices(db *gorm.DB) CampignServices {
	return &campaignServices{campaignRepo: repositories.NewCampaignRepositories(db)}
}

func (service *campaignServices) Create(campaign models.Campaigns) helper.Response {
	var respons helper.Response
	if err := service.campaignRepo.Create(campaign); err != nil {
		respons.Status = 500
		respons.Message = "Failed to create a new Item"
	} else {
		respons.Status = 200
		respons.Message = "Succes to create a new Item"
	}
	return respons
}
func (service *campaignServices) Delete(id string) helper.Response {
	var respons helper.Response
	if err := service.campaignRepo.Delete(id); err != nil {
		respons.Status = 500
		respons.Message = "Failed to delete an Item"
	} else {
		respons.Status = 200
		respons.Message = "Succes to delete an Item"
	}
	return respons
}
func (service *campaignServices) GetAll(id string) helper.Response {
	var respons helper.Response
	data, err := service.campaignRepo.GetAll(id)
	if err != nil {
		respons.Status = 500
		respons.Message = "Failed to get data"
	} else {
		respons.Status = 200
		respons.Message = "Succes to get all Item"
		respons.Data = data
	}
	return respons
}

func (service *campaignServices) GetById(id string) helper.Response {
	var respons helper.Response
	data, err := service.campaignRepo.GetById(id)
	if err != nil {
		respons.Status = 500
		respons.Message = "Failed to get by id"
	} else {
		respons.Status = 200
		respons.Message = "Success to get by id"
		respons.Data = data
	}
	return respons
}
func (service *campaignServices) Update(id string, campaign models.Campaigns) helper.Response {
	var respons helper.Response
	err := service.campaignRepo.Update(id, campaign)
	if err != nil {
		respons.Status = 500
		respons.Message = fmt.Sprintf("failed to update %d", id)
	} else {
		respons.Status = 200
		respons.Message = fmt.Sprintf("Succes to update %d", id)
	}
	return respons
}
