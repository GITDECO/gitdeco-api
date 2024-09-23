package auth

import "gitdeco-api/internal/models"

type Repository interface {
	Create(user *models.User) (*models.User, error)
	ExistsByUsername(username string) (bool, error)
}
