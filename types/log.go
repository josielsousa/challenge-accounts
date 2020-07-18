package types

//APILogProvider - Interface que responsável por inicializar o provider de log
type APILogProvider interface {
	Info(info string)
	Error(trace string, erro error)
}
