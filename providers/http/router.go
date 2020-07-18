package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josielsousa/challenge-accounts/types"
)

// RouterProvider - Implementação do provider de rotas.
type RouterProvider struct {
	mux    *mux.Router
	logger types.APILogProvider
}

//NewRouter - Instância o novo provider com as dependências `mux, log` inicializadas.
func NewRouter(log types.APILogProvider) *RouterProvider {
	return &RouterProvider{
		logger: log,
		mux:    mux.NewRouter(),
	}
}

//Init - Inicializa as rotas da API
func (rp *RouterProvider) ServeHTTP() {
	rp.mux.HandleFunc("/", homeHandler).Methods("GET")

	rp.logger.Info("API disponibilizada na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", rp.mux))
}

//homeHandler - Função utilizada para a rota principal da API `/`
func homeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Desafio técnico accounts."))
}
