package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	mutex           sync.Mutex
	globalValidator Validator
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) ValidateStructModel(structModel any) error {
	return handleValidationError(v.validator.Struct(structModel))
}

func GetGlobalValidator() Validator {
	mutex.Lock()
	defer mutex.Unlock()

	if globalValidator.validator == nil {
		setGlobalValidator()
	}

	return globalValidator
}
