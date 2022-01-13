package transfers

import "time"

// Transfer - Estrutura da entidade `transfer`
type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
