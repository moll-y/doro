package repository

import (
	"gorm.io/gorm"
	"moll-y.io/doro/internal/pkg/domain"
)

type OrganizationRepository struct {
	DB *gorm.DB
}

func (or *OrganizationRepository) CreateOrganization(name, description string, seatsAvailable, seatsOccupied int) (*domain.Organization, error) {
	org := &domain.Organization{
		Name:           name,
		Description:    description,
		SeatsAvailable: seatsAvailable,
		SeatsOccupied:  seatsOccupied,
	}
	r := or.DB.Create(&org)
	if r.Error != nil {
		return nil, r.Error
	}
	return org, nil
}
