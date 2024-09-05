package repositories

import (
	"Campign/domain/item/models"
	"fmt"

	"gorm.io/gorm"
)

type dbCampaign struct {
	Conn *gorm.DB
}

type CampaignRepositories interface {
	Create(campaign models.Campaigns) error
	Update(id string, campaign models.Campaigns) error
	Delete(id string) error
	GetById(id string) (models.Campaigns, error)
	GetAll(id string) ([]models.Campaigns, error)
}

func NewCampaignRepositories(Conn *gorm.DB) CampaignRepositories {
	return &dbCampaign{Conn: Conn}
}

func (db *dbCampaign) Create(campaign models.Campaigns) error {
	return db.Conn.Create(&campaign).Error
}
func (db *dbCampaign) Delete(id string) error {
	fmt.Println(id)
	return db.Conn.Delete(&models.Campaigns{ID: id}).Error
}
func (db *dbCampaign) GetAll(id string) ([]models.Campaigns, error) {
	var data []models.Campaigns
	result := db.Conn.Where("user_id", id).Find(&data)
	return data, result.Error
}

func (db *dbCampaign) GetById(id string) (models.Campaigns, error) {
	var data models.Campaigns
	result := db.Conn.Where("id", id).First(&data)
	return data, result.Error
}

func (db *dbCampaign) Update(id string, campaign models.Campaigns) error {
	return db.Conn.Where("id", id).Updates(campaign).Error
}
