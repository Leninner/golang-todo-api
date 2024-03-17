package utils

import "github.com/go-playground/validator/v10"

func IsInteger(s string) bool {
	var validateInteger = validator.New()
	err := validateInteger.Var(s, "numeric")

	return err == nil
}
