package repository

import (
	"gorm.io/gorm"
	"moll-y.io/doro/internal/pkg/domain"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (tr *TaskRepository) UpdateTask() (*domain.Task, error) {
	var task domain.Task
	tr.DB.First(&task)
	task.State = "completed"
	tr.DB.Save(&task)
	return &task, nil
}

func (tr *TaskRepository) FindTasks() (*[]domain.Task, error) {
	var tasks []domain.Task
	tr.DB.Find(&tasks)
	return &tasks, nil
}
