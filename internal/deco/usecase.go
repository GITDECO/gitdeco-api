package deco

import (
	"gitdeco-api/internal/deco/dto"
	"gitdeco-api/internal/models"
)

type Usecase interface {
	Create(username string, request *dto.DecoRequest) (*models.Deco, error)
	Read(ID uint) (*models.Deco, error)
	ReadAll() ([]*models.Deco, error)
	Update(ID uint, request *dto.DecoRequest) (*models.Deco, error)
	Delete(ID uint) error
}
