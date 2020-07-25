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
	srvAuth     *service.AuthService
	srvAccount  *service.AccountService
	srvTransfer *service.TransferService
}

//NewRouter - Instância o novo provider com as dependências `mux, log` inicializadas.
func NewRouter(srvAuth *service.AuthService, srvAccount *service.AccountService, srvTransfer *service.TransferService, log types.APILogProvider) *RouterProvider {
	return &RouterProvider{
		logger:      log,
		mux:         mux.NewRouter(),
		srvAuth:     srvAuth,
		srvAccount:  srvAccount,
		srvTransfer: srvTransfer,
	}
}

//Init - Inicializa as rotas da API
func (rp *RouterProvider) ServeHTTP() {
	rp.mux.HandleFunc("/", homeHandler).Methods("GET")

	//Inicia as rotas de login e informa qual método interno vai receber as requisições.
	rp.mux.HandleFunc("/login", rp.srvAuth.Login).Methods("POST")

	//Inicia as rotas de accounts e informa qual método interno vai receber as requisições.
	rp.mux.HandleFunc("/accounts", rp.srvAccount.GetAllAccounts).Methods("GET")
	rp.mux.HandleFunc("/accounts", rp.srvAccount.InsertAccount).Methods("POST")
	rp.mux.HandleFunc("/accounts/{id}/balance", rp.srvAccount.GetAccountBalance).Methods("GET")

	//Inicia as rotas de transfer e informa qual método interno vai receber as requisições.
	rp.mux.HandleFunc("/transfers", rp.srvAuth.ValidateToken(rp.srvTransfer.GetAllTransfers)).Methods("GET")
	rp.mux.HandleFunc("/transfers", rp.srvAuth.ValidateToken(rp.srvTransfer.DoTransfer)).Methods("POST")

	rp.logger.Info("API disponibilizada na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", rp.mux))
}

//homeHandler - Função utilizada para a rota principal da API `/`
func homeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Desafio técnico accounts."))
}
