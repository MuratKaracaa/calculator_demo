package calculator

import (
	"regexp"

	"server/internal/shared"

	"github.com/go-playground/validator/v10"
)

var calculationExpressionRegex = regexp.MustCompile(`^[0-9+\-x/s^%().]+$`)

func registerValidations() {
	shared.ValidateInstance.RegisterValidation("calculation_expression", func(fl validator.FieldLevel) bool {
		return calculationExpressionRegex.MatchString(fl.Field().String())
	})
}
