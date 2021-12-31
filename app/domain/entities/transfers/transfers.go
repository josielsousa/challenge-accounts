package transfers

import "time"

// Transfer - Estrutura da entidade `transfer`
type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               float64
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
	DeletedAt            *time.Time
}
