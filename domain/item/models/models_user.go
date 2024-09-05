package models

type User struct {
	ID       string      `json:"id" gorm:"column:id;primaryKey"`
	Name     string      `json:"name" gorm:"column:name"`
	Email    string      `json:"email" gorm:"column:email;unique"`
	Password string      `json:"password" gorm:"column:password"`
	Campaign []Campaigns `json:"campaign" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "user"
}
