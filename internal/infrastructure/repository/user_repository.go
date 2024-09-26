package repository

import (
	"gorm.io/gorm"
	"moll-y.io/doro/internal/domain"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	r := ur.DB.First(user, "email = ?", email)
	if r.Error != nil {
		return nil, r.Error
	}
	return user, nil
}
