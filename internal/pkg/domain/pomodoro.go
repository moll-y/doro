package domain

import "gorm.io/gorm"

type PomodoroRepository interface {
	GetPomodoroByID(id int) (*Pomodoro, error)
	UpdatePomodoro(pomodoro *Pomodoro) error
}

type Pomodoro struct {
	gorm.Model
	State string
}
