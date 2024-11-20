package repository

import (
	"errors"
	"gorm.io/gorm"
	"moll-y.io/doro/internal/pkg/domain"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) CreateUser(name, email, password string) (*domain.User, error) {
	user := &domain.User{Name: name, Email: email, Password: password}
	r := ur.DB.Create(&user)
	if r.Error != nil {
		return nil, r.Error
	}
	return user, nil
}

func (ur *UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	r := ur.DB.First(user, "email = ?", email)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if r.Error != nil {
		return nil, r.Error
	}
	return user, nil
}
