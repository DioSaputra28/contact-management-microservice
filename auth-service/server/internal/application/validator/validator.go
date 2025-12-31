package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s any) error {
	err := validate.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationErrors {
				errorMessages = append(errorMessages, formatFieldError(fieldError))
			}
			return fmt.Errorf("validation failed: %s", strings.Join(errorMessages, ", "))
		}
		return err
	}
	return nil
}


func formatFieldError(fe validator.FieldError) string {
	field := fe.Field()

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, fe.Param())
	case "alphanum":
		return fmt.Sprintf("%s must contain only alphanumeric characters", field)
	default:
		return fmt.Sprintf("%s failed validation on '%s'", field, fe.Tag())
	}
}
