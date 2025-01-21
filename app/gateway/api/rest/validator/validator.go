package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	onceValidator   sync.Once
	globalValidator Validator
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) ValidateStructModel(structModel any) error {
	return handleValidationError(v.validator.Struct(structModel))
}

func GetGlobalValidator() Validator {
	if globalValidator.validator == nil {
		setGlobalValidator()
	}

	return globalValidator
}
