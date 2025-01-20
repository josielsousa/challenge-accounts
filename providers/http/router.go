package http

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/josielsousa/challenge-accounts/service"
)

// RouterProvider - Implementação do provider de rotas.
type RouterProvider struct {
	mux         *mux.Router
	srvAuth     *service.AuthService
	srvAccount  *service.AccountService
	srvTransfer *service.TransferService
}

// NewRouter - Instância o novo provider com as dependências `mux, log` inicializadas.
func NewRouter(
	srvAuth *service.AuthService,
	srvAccount *service.AccountService,
	srvTransfer *service.TransferService,
) *RouterProvider {
	return &RouterProvider{
		mux:         mux.NewRouter(),
		srvAuth:     srvAuth,
		srvAccount:  srvAccount,
		srvTransfer: srvTransfer,
	}
}

// Init - Inicializa as rotas da API.
//
// TODO: use chi as router
func (rp *RouterProvider) ServeHTTP() {
	rp.mux.HandleFunc("/", homeHandler).Methods("GET")

	// Inicia as rotas de login e informa qual método interno vai receber as requisições.

	// Inicia as rotas de accounts e informa qual método interno vai receber as requisições.
	rp.mux.HandleFunc("/accounts", rp.srvAccount.GetAllAccounts).Methods("GET")
	rp.mux.HandleFunc("/accounts", rp.srvAccount.InsertAccount).Methods("POST")
	rp.mux.HandleFunc("/accounts/{id}/balance", rp.srvAccount.GetAccountBalance).Methods("GET")

	// Inicia as rotas de transfer e informa qual método interno vai receber as requisições.
	rp.mux.HandleFunc("/transfers", rp.srvTransfer.GetAllTransfers).Methods("GET")
	rp.mux.HandleFunc("/transfers", rp.srvTransfer.DoTransfer).Methods("POST")

	slog.Info("API disponibilizada na porta 3000")

	//nolint:gosec
	if err := http.ListenAndServe(":3000", rp.mux); err != nil {
		slog.Error("on start server: ", slog.Any("error", err))
	}
}

// homeHandler - Função utilizada para a rota principal da API `/`.
func homeHandler(w http.ResponseWriter, _ *http.Request) {
	if _, err := w.Write([]byte("Desafio técnico accounts.")); err != nil {
		slog.Error("on write response: ", slog.Any("error", err))
	}
}
