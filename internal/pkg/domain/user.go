package domain

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(name, email, password string) (*User, error)
	FindUserByEmail(email string) (*User, error)
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
