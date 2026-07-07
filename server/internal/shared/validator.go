package shared

import (
	"github.com/go-playground/validator/v10"
)

var ValidateInstance *validator.Validate

func InitValidator() {
	ValidateInstance = validator.New()
}

type Validator interface {
	Validate() error
}
