package validator

import (
	"github.com/leebenson/conform"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate - validates user model
func Validate(u interface{}) (errors map[string]string) {
	// Trim, sanitize, and modify struct string fields
	conform.Strings(u)
	errors = make(map[string]string)
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.ActualTag()
		}
	}

	return errors
}

// ValidateVar - validate custom field
func ValidateVar(field interface{}, validationString string) error {
	return validate.Var(field, validationString)
}
