package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josielsousa/challenge-accounts/service"
	"github.com/josielsousa/challenge-accounts/types"
)

// RouterProvider - Implementação do provider de rotas.
type RouterProvider struct {
	mux         *mux.Router
	logger      types.APILogProvider
	srvAccount  *service.AccountService
	srvTransfer *service.TransferService
}

//NewRouter - Instância o novo provider com as dependências `mux, log` inicializadas.
func NewRouter(srvAccount *service.AccountService, srvTransfer *service.TransferService, log types.APILogProvider) *RouterProvider {
	return &RouterProvider{
		logger:      log,
		mux:         mux.NewRouter(),
		srvAccount:  srvAccount,
		srvTransfer: srvTransfer,
	}
}

//Init - Inicializa as rotas da API
func (rp *RouterProvider) ServeHTTP() {
	rp.mux.HandleFunc("/", homeHandler).Methods("GET")

	//Inicia as rotas de accounts e informa qual método interno vai receber a REQUEST
	rp.mux.HandleFunc("/accounts", rp.srvAccount.GetAllAccounts).Methods("GET")
	rp.mux.HandleFunc("/accounts", rp.srvAccount.InsertAccount).Methods("POST")
	rp.mux.HandleFunc("/accounts/{id}", rp.srvAccount.GetAccount).Methods("GET")
	rp.mux.HandleFunc("/accounts/{id}", rp.srvAccount.UpdateAccount).Methods("PUT")

	rp.logger.Info("API disponibilizada na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", rp.mux))
}

//homeHandler - Função utilizada para a rota principal da API `/`
func homeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Desafio técnico accounts."))
}
