package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" validate:"required,min=3,max=40"`
	Email     string         `json:"email" validate:"required,email"`
	Role      string         `json:"role" validate:"required,oneof=admin user"`
	Password  string         `json:"password" validate:"required,min=6"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
