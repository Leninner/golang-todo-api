package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       int
	Username string
	Email    string
	Password string
}

func (t *User) TableName() string {
	return "users"
}
