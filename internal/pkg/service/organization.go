package service

import (
	"log"
	"moll-y.io/doro/internal/pkg/domain"
)

type OrganizationService struct {
	UserRepository         domain.UserRepository
	OrganizationRepository domain.OrganizationRepository
}

func (os *OrganizationService) CreateOrganization(actor, name, description string) (*domain.Organization, error) {
	user, err := os.UserRepository.FindUserByID(actor)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if user == nil {
	}
	org, err := os.OrganizationRepository.CreateOrganization(name, description, 0, 0)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return org, nil
}

func (os *OrganizationService) ImportOrganization(actor, source string) (*domain.Organization, error) {
	org, err := os.OrganizationRepository.CreateOrganization(source, "left right left", 0, 0)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return org, nil
}
