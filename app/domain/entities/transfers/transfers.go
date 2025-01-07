package transfers

import (
	"time"
)

// Transfer - Estrutura da entidade `transfer`.
type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type AccountData struct {
	ID      string
	Balance int
}

type TransferData struct {
	Transfer
	AccountOrigin      AccountData
	AccountDestination AccountData
}
