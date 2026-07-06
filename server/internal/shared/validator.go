package shared

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var ValidateInstance *validator.Validate

func InitValidator() {
	ValidateInstance = validator.New()

	ValidateInstance.RegisterValidation("calculation_expression", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`^[0-9+\-x/s^%().]+$`)
		return re.MatchString(fl.Field().String())
	})
}

type Validator interface {
	Validate() error
}
