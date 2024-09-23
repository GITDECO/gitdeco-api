package user

import (
	"gitdeco-api/internal/models"
)

type Usecase interface {
	Read(username string) (*models.User, error)
	ReadAll() ([]*models.User, error)
}
