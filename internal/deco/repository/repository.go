package repository

import (
	"gitdeco-api/internal/deco"
	"gitdeco-api/internal/models"

	"gorm.io/gorm"
)

type DecoRepository struct {
	db *gorm.DB
}

func NewDecoRepository(db *gorm.DB) deco.Repository {
	return &DecoRepository{
		db: db,
	}
}

func (dr *DecoRepository) Create(deco *models.Deco) (*models.Deco, error) {
	err := dr.db.Create(&deco).Error
	return deco, err
}

func (dr *DecoRepository) Read(ID uint) (*models.Deco, error) {
	deco := new(models.Deco)
	err := dr.db.Where("ID = ?", ID).First(&deco).Error
	return deco, err
}

func (dr *DecoRepository) ReadAll() ([]*models.Deco, error) {
	decos := []*models.Deco{}
	err := dr.db.Find(&decos).Error
	return decos, err
}

func (dr *DecoRepository) Update(deco *models.Deco) (*models.Deco, error) {
	err := dr.db.Save(&deco).Error
	return deco, err
}

func (dr *DecoRepository) Delete(deco *models.Deco) error {
	err := dr.db.Delete(&deco).Error
	return err
}
