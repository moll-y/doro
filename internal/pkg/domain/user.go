package domain

import "gorm.io/gorm"

type UserRepository interface {
	FindUserByEmail(email string) (*User, error)
}

type User struct {
	gorm.Model
	Email    string
	Password string
	Role     Role
}
