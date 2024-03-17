package tasks

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" validate:"required,min=3,max=40"`
	Completed   bool           `json:"completed" gorm:"default:false"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (t *Task) TableName() string {
	return "tasks"
}
