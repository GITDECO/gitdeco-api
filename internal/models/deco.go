package models

import (
	"time"
)

type Deco struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"unique,not null"`
	Markdown  string `json:"markdown" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"not null"`
	UpdatedAt string `json:"updated_at" gorm:"not null"`

	Username string `gorm:"type:varchar(255)"`
}

func (deco *Deco) Update(name string, markdown string) *Deco {
	deco.Title = name
	deco.Markdown = markdown
	deco.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return deco
}

func (Deco) TableName() string {
	return "decos"
}
