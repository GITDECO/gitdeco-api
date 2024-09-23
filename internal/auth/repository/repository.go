package repository

import (
	"gitdeco-api/internal/auth"
	"gitdeco-api/internal/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) Create(user *models.User) (*models.User, error) {
	err := ar.db.Create(&user).Error
	return user, err
}

func (ar *AuthRepository) ExistsByUsername(username string) (bool, error) {
	user := new(models.User)
	if err := ar.db.Where("Username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
