package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
)

type RulerError struct {
	Code    string
	Message string
}

func (v RulerError) Error() string {
	return v.Message
}

func (v RulerError) addField(field string) RulerError {
	v.Code = fmt.Sprintf("%s:%s", field, v.Code)
	v.Message = fmt.Sprintf("%s %s", field, v.Message)

	return v
}

type Ruler struct {
	fn     validator.Func
	vError RulerError
}

var validations = map[string]Ruler{
	"cpf": {
		fn: func(field validator.FieldLevel) bool {
			if _, err := cpf.NewCPF(field.Field().String()); err != nil {
				return false
			}

			return true
		},
		vError: RulerError{"cpf", "must be a valid cpf"},
	},
}

func setGlobalValidator() {
	const maxSplitTags = 2

	validate := validator.New(validator.WithRequiredStructEnabled())

	// This code sets the field value as the json tag.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", maxSplitTags)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}

		return name
	})

	for key, val := range validations {
		_ = validate.RegisterValidation(key, val.fn)
	}

	globalValidator = Validator{
		validator: validate,
	}
}

func handleValidationError(err error) error {
	var vErrors validator.ValidationErrors
	if errors.As(err, &vErrors) {
		for _, vError := range vErrors {
			validation, ok := validations[vError.Tag()]
			if !ok {
				continue
			}

			return validation.vError.addField(vError.Field())
		}
	}

	return err
}
