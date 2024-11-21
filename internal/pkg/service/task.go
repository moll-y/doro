package service

import (
	"moll-y.io/doro/internal/pkg/domain"
)

type TaskService struct {
	TaskRepository domain.TaskRepository
}

func (ts *TaskService) UpdateTask() (*domain.Task, error) {
	return ts.TaskRepository.UpdateTask()
}

func (ts *TaskService) FindTasks() (*[]domain.Task, error) {
	return ts.TaskRepository.FindTasks()
}
