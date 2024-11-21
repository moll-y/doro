package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"moll-y.io/doro/internal/pkg/domain"
)

type BacklogRepository struct {
	DB *gorm.DB
}

func (ur *BacklogRepository) FindBacklogs() (*[]domain.Backlog, error) {
	var backlogs []domain.Backlog
	r := ur.DB.Find(&backlogs)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("%+v\n", r)
	}
	if r.Error != nil {
		fmt.Printf("%+v\n", r)
	}
	return &backlogs, nil
}
