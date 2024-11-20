package domain

import "gorm.io/gorm"

type BacklogRepository interface {
	FindBacklogs() ([]*Backlog, error)
}

type Backlog struct {
	gorm.Model
	NumberOfTasksAllowedToCreate int
}
