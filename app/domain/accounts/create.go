package accounts

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/josielsousa/challenge-accounts/app/domain/entities"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/cpf"
	"github.com/josielsousa/challenge-accounts/app/domain/vos/hash"
)

func (a Account) Create(ctx context.Context, input AccountInput) error {
	const op = `accounts.Create`

	cpf, err := cpf.NewCPF(input.CPF)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on new instance CPF from input", err)
	}

	secret, err := hash.NewHash(input.Secret)
	if err != nil {
		return fmt.Errorf("%s -> %s: %w", op, "on new hashed secret from input", err)
	}

	acc := entities.Account{
		ID:        uuid.NewString(),
		Name:      input.Name,
		Balance:   input.Balance,
		CPF:       cpf,
		Secret:    secret,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// there are a validation to check if the account already exists
	// using the constraint of the CPF as unique on the database.
	err = a.R.Insert(ctx, acc)
	if err != nil {
		return fmt.Errorf("%s-> %s: %w", op, "on create account", err)
	}

	return nil
}
