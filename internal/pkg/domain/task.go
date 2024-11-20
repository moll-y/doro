package domain

import "gorm.io/gorm"

type TaskRepository interface {
	UpdateTask(task *Task) error
	FindTasks() ([]*Task, error)
}

type Task struct {
	gorm.Model
	Name  string
	State string
}
