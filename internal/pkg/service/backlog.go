package service

import (
	"moll-y.io/doro/internal/pkg/domain"
)

type BacklogService struct {
	BacklogRepository domain.BacklogRepository
}

func (us *BacklogService) FindBacklogs() (*[]domain.Backlog, error) {
	return us.BacklogRepository.FindBacklogs()
}
