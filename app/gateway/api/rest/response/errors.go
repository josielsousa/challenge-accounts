package response

import (
	"errors"
	"net/http"

	"github.com/josielsousa/challenge-accounts/app/domain/erring"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type AppErr struct {
	StatusCode int
	Error      Error
}

var UnauthorizedErr = Error{
	Code:    "app:unauthorized",
	Message: erring.ErrUnauthorized.Error(),
}

var appErrors = map[error]AppErr{
	erring.ErrAccountNotFound: {
		StatusCode: http.StatusNotFound,
		Error: Error{
			Code:    "app:not_found",
			Message: erring.ErrAccountNotFound.Error(),
		},
	},

	erring.ErrInvalidAmount: {
		StatusCode: http.StatusBadRequest,
		Error: Error{
			Code:    "app:invalid_amount",
			Message: erring.ErrInvalidAmount.Error(),
		},
	},

	erring.ErrUpdateAccountNotPerformed: {
		StatusCode: http.StatusUnprocessableEntity,
		Error: Error{
			Code:    "app:unprocessable_entity",
			Message: erring.ErrUpdateAccountNotPerformed.Error(),
		},
	},

	erring.ErrAccountOriginNotFound: {
		StatusCode: http.StatusNotFound,
		Error: Error{
			Code:    "app:not_found",
			Message: erring.ErrAccountOriginNotFound.Error(),
		},
	},

	erring.ErrAccountDestinationNotFound: {
		StatusCode: http.StatusNotFound,
		Error: Error{
			Code:    "app:not_found",
			Message: erring.ErrAccountDestinationNotFound.Error(),
		},
	},

	erring.ErrAccountAlreadyExists: {
		StatusCode: http.StatusConflict,
		Error: Error{
			Code:    "app:invalid_amount",
			Message: erring.ErrAccountAlreadyExists.Error(),
		},
	},

	erring.ErrInsufficientFunds: {
		StatusCode: http.StatusBadRequest,
		Error: Error{
			Code:    "app:insufficient_funds",
			Message: erring.ErrInsufficientFunds.Error(),
		},
	},

	erring.ErrExpiredToken: {
		StatusCode: http.StatusUnauthorized,
		Error: Error{
			Code:    "app:invalid_amount",
			Message: erring.ErrExpiredToken.Error(),
		},
	},

	erring.ErrEmptyToken: {
		StatusCode: http.StatusForbidden,
		Error: Error{
			Code:    "app:fobidden",
			Message: erring.ErrEmptyToken.Error(),
		},
	},

	erring.ErrInvalidToken: {
		StatusCode: http.StatusUnauthorized,
		Error: Error{
			Code:    "app:invalid_token",
			Message: erring.ErrInvalidToken.Error(),
		},
	},

	erring.ErrMalformedToken: {
		StatusCode: http.StatusUnauthorized,
		Error: Error{
			Code:    "app:malformed_token",
			Message: erring.ErrMalformedToken.Error(),
		},
	},

	erring.ErrParseTokenWithClaims: {
		StatusCode: http.StatusForbidden,
		Error: Error{
			Code:    "app:forbidden",
			Message: erring.ErrParseTokenWithClaims.Error(),
		},
	},

	erring.ErrSignatureKeyInvalid: {
		StatusCode: http.StatusInternalServerError,
		Error: Error{
			Code:    "app:internal_server_error",
			Message: "internal server error",
		},
	},

	erring.ErrUnexpected: {
		StatusCode: http.StatusInternalServerError,
		Error: Error{
			Code:    "app:internal_server_error",
			Message: "an unexpected internal server error",
		},
	},

	erring.ErrUnauthorized: {
		StatusCode: http.StatusUnauthorized,
		Error:      UnauthorizedErr,
	},

	erring.ErrRecordNotFound: {
		StatusCode: http.StatusNotFound,
		Error: Error{
			Code:    "app:not_found",
			Message: erring.ErrRecordNotFound.Error(),
		},
	},
}

func ToAppErr(err error) AppErr {
	for value, appErr := range appErrors {
		if errors.Is(value, err) {
			return appErr
		}
	}

	return AppErr{
		StatusCode: http.StatusNotImplemented,
		Error: Error{
			Code:    "app:not_implemented",
			Message: "not implemented",
		},
	}
}
