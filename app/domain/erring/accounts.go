package erring

import "errors"

var (
	ErrInvalidAmount              = errors.New("invalid amount")
	ErrAccountNotFound            = errors.New("account not found")
	ErrUpdateAccountNotPerformed  = errors.New("update account not performed")
	ErrAccountOriginNotFound      = errors.New("account origin not found")
	ErrAccountDestinationNotFound = errors.New("account destination not found")
	ErrAccountAlreadyExists       = errors.New("account already exists")
	ErrInsufficientFunds          = errors.New("insufficient amount")
)
