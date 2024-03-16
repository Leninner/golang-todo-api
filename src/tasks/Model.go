package tasks

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	ID          int
	Title       string
	Completed   bool
	Description string
}

func (t *Task) TableName() string {
	return "tasks"
}
