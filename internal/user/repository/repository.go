package repository

import (
	"gitdeco-api/internal/models"
	"gitdeco-api/internal/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Read(username string) (*models.User, error) {
	user := new(models.User)
	err := ur.db.Where("Username = ?", username).Preload("Decos").First(&user).Error
	return user, err
}

func (ur *UserRepository) ReadAll() ([]*models.User, error) {
	users := []*models.User{}
	err := ur.db.Find(&users).Error
	return users, err
}
