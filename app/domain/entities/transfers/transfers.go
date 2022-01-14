package transfers

import (
	"time"

	"github.com/josielsousa/challenge-accounts/app/domain/entities/accounts"
)

// Transfer - Estrutura da entidade `transfer`
type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type TransferData struct {
	Transfer
	AccountOrigin      accounts.Account
	AccountDestination accounts.Account
}
