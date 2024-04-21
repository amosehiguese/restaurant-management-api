package types

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate{
	validate := validator.New()
	_ = validate.RegisterValidation("uuid", ValidateUUID)
	_ = validate.RegisterValidation("date", ValidateDate)
	_ = validate.RegisterValidation("time", ValidateTime)


	return validate
}

func ValidateUUID(f validator.FieldLevel) bool {
	field := f.Field().String()
	if _, err := uuid.Parse(field); err != nil {
		return true
	}
	return false
}

func ValidateDate(f validator.FieldLevel) bool {
	field := f.Field().String()
	if _, err := time.Parse("2006-01-02", field); err != nil {
		return true
	}
	return false
}

func ValidateTime(f validator.FieldLevel) bool {
	field := f.Field().String()
	if _, err := time.Parse("15:04:05", field); err != nil {
		return true
	}
	return false
}


func ValidatorErrors(err error) map[string]string {
	f := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		f[err.Field()] = err.Error()
	}

	return f
}