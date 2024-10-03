package domain

import "gorm.io/gorm"

type OrganizationRepository interface {
	CreateOrganization(name, description string, seatsAvailable, seatsOccupied int) (*Organization, error)
}

type Organization struct {
	gorm.Model
	Name           string
	Description    string
	SeatsAvailable int
	SeatsOccupied  int
}
