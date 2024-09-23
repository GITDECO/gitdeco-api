package user

import "gitdeco-api/internal/models"

type Repository interface {
	Read(username string) (*models.User, error)
	ReadAll() ([]*models.User, error)
}
