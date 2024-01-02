package types

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate{
	validate := validator.New()
	_ = validate.RegisterValidation("uuid", func(f validator.FieldLevel) bool {
		field := f.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

func ValidatorErrors(err error) map[string]string {
	f := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		f[err.Field()] = err.Error()
	}

	return f
}