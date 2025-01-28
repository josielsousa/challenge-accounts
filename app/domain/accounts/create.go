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

func (u Usecase) Create(
	ctx context.Context, input AccountInput,
) (AccountOutput, error) {
	const op = `accounts.Create`

	cpf, err := cpf.NewCPF(input.CPF)
	if err != nil {
		return AccountOutput{}, fmt.Errorf("%s-> new cpf: %w", op, err)
	}

	secret, err := hash.NewHash(input.Secret)
	if err != nil {
		return AccountOutput{}, fmt.Errorf("%s -> hash secret: %w", op, err)
	}

	now := time.Now()

	acc := entities.Account{
		ID:        uuid.NewString(),
		Name:      input.Name,
		Balance:   input.Balance,
		CPF:       cpf,
		Secret:    secret,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// there are a validation to check if the account already exists
	// using the constraint of the CPF as unique on the database.
	err = u.R.Insert(ctx, acc)
	if err != nil {
		return AccountOutput{}, fmt.Errorf(
			"%s-> on create account: %w", op, err,
		)
	}

	return AccountOutput{
		ID:        acc.ID,
		Name:      acc.Name,
		Balance:   acc.Balance,
		CPF:       acc.CPF.Value(),
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
	}, nil
}
