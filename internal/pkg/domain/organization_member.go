package domain

import "gorm.io/gorm"

type OrganizationMember struct {
	gorm.Model
	UserID         uint
	OrganizationID uint
	Role           string
	User           User         `gorm:"foreignKey:UserID"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
}
