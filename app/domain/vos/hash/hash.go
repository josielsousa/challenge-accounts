package hash

import (
	"errors"

	"github.com/josielsousa/challenge-accounts/app/common"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidSecret     = errors.New("invalid secret")
	ErrScanInvalidSecret = errors.New("scan invalid secret")
	ErrGenerateHash      = errors.New("unexpected error generating hash")
)

type Hash struct {
	hashedValue string
}

func NewHash(secret string) (Hash, error) {
	h := Hash{}
	hs, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	if err != nil {
		return Hash{}, ErrGenerateHash
	}

	h.hashedValue = string(hs)
	return h, nil
}

// Compare - verifica se a senha informada é a mesma salva na account.
func (h Hash) Compare(secret string) error {
	err := bcrypt.CompareHashAndPassword([]byte(h.hashedValue), []byte(secret))
	if err != nil {
		return ErrInvalidSecret
	}

	return nil
}

func (h Hash) Value() string {
	return h.hashedValue
}

// Scan implements the database/sql/driver Scanner interface
func (h *Hash) Scan(value interface{}) error {
	if value == nil {
		*h = Hash{}
		return common.ErrScanValueNil
	}

	if value, ok := value.(string); ok {
		if _, err := bcrypt.Cost([]byte(value)); err != nil {
			return ErrScanInvalidSecret
		}

		*h = Hash{hashedValue: value}
		return nil
	}

	return common.ErrScanValueIsNotString
}
