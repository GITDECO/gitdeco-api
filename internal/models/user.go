package models

type User struct {
	Username string `json:"username" gorm:"type:varchar(255);primaryKey"`
	Email    string `json:"email" gorm:"null"`
	Name     string `json:"name" gorm:"null"`
	Bio      string `json:"bio" gorm:"null"`
	Avatar   string `json:"avatar" gorm:"null"`

	Decos []Deco `gorm:"foreignKey:Username"`
}

func (User) TableName() string {
	return "users"
}
