package types

//ErrorResponse - Estrutura padrão para respostas de solicitações que lançaram erro.
type ErrorResponse struct {
	Error interface{} `json:"error"`
}

//SuccessResponse - Estrutura padrão para respostas de solicitações com sucesso.
type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}
