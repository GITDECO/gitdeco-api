package usecase

import (
	"gitdeco-api/internal/auth"
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/models"
	"gitdeco-api/pkg"
)

type AuthUsecase struct {
	authRepository auth.Repository
}

func NewAuthUsecase(authRepository auth.Repository) auth.Usecase {
	return &AuthUsecase{authRepository: authRepository}
}

func (au *AuthUsecase) Auth(user *models.User) (*pkg.Token, error) {
	exists, err := au.authRepository.ExistsByUsername(user.Username)
	if err != nil {
		panic(&exception.Error{Key: "REPO_ERROR", Data: err.Error()})
	}
	if exists {
		return pkg.GenerateToken(user.Username, false), nil
	}
	newUser, err := au.authRepository.Create(user)
	if err != nil {
		panic(&exception.Error{Key: "REPO_ERROR", Data: err.Error()})
	}
	return pkg.GenerateToken(newUser.Username, false), nil
}
