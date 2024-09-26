package service

import (
	"log"
	"moll-y.io/doro/internal/pkg/domain"
)

type UserService struct {
	UserRepository domain.UserRepository
}

func (us *UserService) FindUserByEmail(email string) (*domain.User, error) {
	user, err := us.UserRepository.FindUserByEmail(email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
