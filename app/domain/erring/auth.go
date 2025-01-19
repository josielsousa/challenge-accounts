package erring

import "errors"

var (
	ErrEmptyToken           = errors.New("token is empty")
	ErrExpiredToken         = errors.New("o token está expirado")
	ErrInvalidToken         = errors.New("token inválido")
	ErrMalformedToken       = errors.New("token mal formado")
	ErrParseTokenWithClaims = errors.New("o token não pode ser analisado")
	ErrSignatureKeyInvalid  = errors.New("a chave de assinatura do token é inválida")

	ErrUnexpected   = errors.New("erro Inesperado")
	ErrUnauthorized = errors.New("não autenticado")

	ErrRecordNotFound = errors.New("record not found")
)
