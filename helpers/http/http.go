package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/josielsousa/challenge-accounts/types"
)

// Helper - Define a struct para o helper HTTP.
type Helper struct{}

// NewHelper - Instância helper HTTP.
func NewHelper() *Helper {
	return &Helper{}
}

// GetParams - Encapsula a recuperação de paramêtros.
func (h *Helper) GetParams(req *http.Request) map[string]string {
	return mux.Vars(req)
}

// ThrowError - Retorna um Erro de acordo com o formato passado e com o statusCode informado.
func (h *Helper) ThrowError(w http.ResponseWriter, code int, Error interface{}) {
	h.StatusCode(w, code)
	h.JSONEncoder(w, types.ErrorResponse{Error: Error})
}

// ThrowSuccess - Retorna Sucesso com o Data informado.
func (h *Helper) ThrowSuccess(w http.ResponseWriter, code int, Data interface{}) {
	h.StatusCode(w, code)
	h.JSONEncoder(w, Data)
}

// StatusCode - Seta o statusCode no response.
func (h *Helper) StatusCode(w http.ResponseWriter, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
}

// JSONEncoder - Seta um Novo Encoder no Writer.
func (h *Helper) JSONEncoder(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
