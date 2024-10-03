package service

import (
	"log"
	"moll-y.io/doro/internal/pkg/domain"
)

type OrganizationService struct {
	UserRepository         domain.UserRepository
	OrganizationRepository domain.OrganizationRepository
}

func (os *OrganizationService) CreateOrganization(actor int, name, description string) (*domain.Organization, error) {
	org, err := os.OrganizationRepository.CreateOrganization(name, description, 0, 0)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return org, nil
}
