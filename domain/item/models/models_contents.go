package models

type Contents struct {
	ID          string `json:"id" gorm:"column:id;primaryKey"`
	Campaign_id string `json:"campaign_id" gorm:"column:campaign_id"`
	Title       string `json:"title" gorm:"column:title"`
	Content     string `json:"content" gorm:"column:content"`
}

func (Contents) TableName() string {
	return "Contents"
}
