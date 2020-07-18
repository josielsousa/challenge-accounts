package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RouterProvider - Implementação do provider de rotas.
type RouterProvider struct {
}

//Init - Inicializa as rotas da API
func (rp *RouterProvider) Init() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")

	fmt.Println("API disponibilizada na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

//homeHandler - Função utilizada para a rota principal da API `/`
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Desafio técnico accounts."))
}
