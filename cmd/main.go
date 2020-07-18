package main

import (
	"fmt"

	"github.com/josielsousa/challenge-accounts/providers/http"
)

func main() {
	fmt.Println("API inicializando...")

	router := &http.RouterProvider{}
	router.Init()
}
