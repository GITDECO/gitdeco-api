package auth

import (
	"gitdeco-api/internal/models"
	"gitdeco-api/pkg"
)

type Usecase interface {
	Auth(user *models.User) (*pkg.Token, error)
}
