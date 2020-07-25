package validation

import (
	"github.com/thedevsaddam/govalidator"
)

// Mensagens utilizadas no contexto de validação
const (
	ErrorFieldNotString = "The %s should be a string"
)

// Define as regras para validação dos campos de `account`
var validateRulesAccount = map[string][]string{
	"cpf":      {"required", "string"},
	"name":     {"required", "string"},
	"secret":   {"required", "string"},
	"ballance": {"required", "numeric_between:0.01,999999999.99"},
}

// Define as mensagens para validação dos campos de `account`
var validateMessagesAccount = govalidator.MapData{
	"cpf":      {"required:O campo é obrigatório", "string: O campo tem formato inválido"},
	"name":     {"required:O campo é obrigatório", "string: O campo tem formato inválido"},
	"secret":   {"required:O campo é obrigatório", "string: O campo tem formato inválido"},
	"ballance": {"required:O campo é obrigatório", "numeric_between: O campo deve possuir valor entre 0,01 e 999.999.999,99"},
}

// Define as regras para validação dos campos de `transfer`
var validateRulesTransfer = map[string][]string{
	"account_destination_id": {"required", "string"},
	"amount":                 {"required", "numeric_between:0.01,999999999.99"},
}

// Define as mensagens para validação dos campos de `account`
var validateMessagesTransfer = govalidator.MapData{
	"account_destination_id": {"required:O campo é obrigatório", "string: O campo tem formato inválido"},
	"amount":                 {"required:O campo é obrigatório", "numeric_between: O campo deve possuir valor entre 0,01 e 999.999.999,99"},
}

// Define as regras para validação dos campos de `login`
var validateRulesLogin = map[string][]string{
	"cpf":    {"required", "string"},
	"secret": {"required", "string"},
}

// Define as mensagens para validação dos campos de `login`
var validateMessagesLogin = govalidator.MapData{
	"cpf":    {"required:O campo é obrigatório", "string: O campo tem formato inválido"},
	"secret": {"required:O campo é obrigatório", "string: O campo tem formato inválido"},
}
