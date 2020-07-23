package model

import "time"

//Constantes para trabalhar com o model `transfers`.
const (
	TransferTablename = "transfers"
)

//Transfer - Estrutura da entidade `transfer`
type Transfer struct {
	ID                   string     `json:"id,omitempty"`
	AccountOriginID      string     `json:"account_origin_id"`
	AccountDestinationID string     `json:"account_destination_id"`
	Amount               float64    `json:"amount"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"-"`
	DeletedAt            *time.Time `json:"-"`
}

//TransferStorage - Interface que define as assinaturas para o storage de transfers.
type TransferStorage interface {
	GetAllTransfers() ([]Transfer, error)
	Insert(transfer Transfer) (*Transfer, error)
}
