package log

import (
	"github.com/sirupsen/logrus"
)

// APILogProvider - Implementação do provider de logs.
type APILogProvider struct {
	log *logrus.Logger
}

// NewLogger - Instância o novo provider com a dependência `logrus` inicializada.
func NewLogger() *APILogProvider {
	return &APILogProvider{
		log: logrus.StandardLogger(),
	}
}

// Info - Log a informação passadas.
func (api *APILogProvider) Info(info string) {
	api.log.Info(info)
}

// Error - Log o erro informado.
func (api *APILogProvider) Error(trace string, erro error) {
	api.log.Error(trace, erro)
}
