package usecase

import (
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/models"
	"gitdeco-api/internal/user"
)

type UserUsecase struct {
	userRepository user.Repository
}

func NewUserUsecase(userRepository user.Repository) user.Usecase {
	return &UserUsecase{userRepository: userRepository}
}

func (uu *UserUsecase) Read(username string) (*models.User, error) {
	user, err := uu.userRepository.Read(username)
	if err != nil {
		panic(&exception.Error{Key: "REPO_ERROR", Data: err.Error()})
	}
	return user, nil
}

func (uu *UserUsecase) ReadAll() ([]*models.User, error) {
	users, err := uu.userRepository.ReadAll()
	if err != nil {
		panic(&exception.Error{Key: "REPO_ERROR", Data: err.Error()})
	}
	return users, nil
}
