package repository

import (
	"dropbox-clone/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// Create user into db if you want
func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

// Get the user by google id from the db
func (r *UserRepository) GetByGoogleID(googleID string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("google_id = ?", googleID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	result := r.DB.Where("user_id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
