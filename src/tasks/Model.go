package tasks

import (
	"todo-api/src/core/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" validate:"required,min=3,max=40"`
	Completed   bool   `json:"completed" gorm:"default:false"`
	Description string `json:"description"`
}

func (t *Task) TableName() string {
	return "tasks"
}

func (t *Task) Validate() []*utils.ErrorResponse {

	var errors []*utils.ErrorResponse
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error := &utils.ErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			}

			errors = append(errors, error)
		}
	}

	return errors

}
