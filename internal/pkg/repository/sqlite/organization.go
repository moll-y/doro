package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"errors"
	"moll-y.io/doro/internal/pkg/domain"
)

type OrganizationRepository struct {
	DB *gorm.DB
}

func (or *OrganizationRepository) CreateOrganization(name, description string, seatsAvailable, seatsOccupied int) (*domain.Organization, error) {
  o := domain.Organization{}
  or.DB.Preload("User").Preload(clause.Associations).Find(&o, 4)
  for _, m := range o.Members {
    log.Printf("---------> %+v\n", *m)
  }
  return nil, errors.New("test")
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
