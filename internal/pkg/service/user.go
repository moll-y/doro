package service

import (
	"fmt"
	"log"
	"moll-y.io/doro/internal/pkg/domain"
)

type UserService struct {
	UserRepository domain.UserRepository
}

func (us *UserService) CreateUser(name, email, password string) (*domain.User, error) {
	user, err := us.UserRepository.FindUserByEmail(name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if user != nil {
		log.Printf(`User with email "%s" already exists.\n`, email)
		return nil, fmt.Errorf(`User with email "%s" already exists.`, email)
	}
	user, err = us.UserRepository.CreateUser(name, email, password)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
