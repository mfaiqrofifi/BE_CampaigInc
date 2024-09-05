package models

type Campaigns struct {
	ID          string     `json:"id" gorm:"column:id;primaryKey"`
	UserID      string     `json:"user_id" gorm:"column:user_id"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Content     []Contents `json:"content" gorm:"foreignKey:CampaignID"`
}

func (Campaigns) TableName() string {
	return "campaigns"
}
