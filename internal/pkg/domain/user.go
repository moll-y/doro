package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string
	Password string
	Role     Role
}
