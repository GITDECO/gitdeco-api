package deco

import "gitdeco-api/internal/models"

type Repository interface {
	Create(deco *models.Deco) (*models.Deco, error)
	Read(ID uint) (*models.Deco, error)
	ReadAll() ([]*models.Deco, error)
	Update(deco *models.Deco) (*models.Deco, error)
	Delete(deco *models.Deco) error
}
