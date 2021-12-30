package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
)

const (
	ErrorGetParams      = "Valor retornado %s diferente do esperado %s"
	ErrorOnThrowError   = "Deveria retornar status 500; retornou %d"
	ErrorOnThrowSuccess = "Deveria retornar status 200; retornou %d"
)

func TestHelperGetParams(t *testing.T) {
	t.Run("Teste Get Params com sucesso", func(t *testing.T) {
		id := "35"
		vars := map[string]string{"id": id}

		// Criando nova fake request para teste
		url := fmt.Sprintf("http://localhost:3000/accounts/%s", vars["id"])
		mockReq := httptest.NewRequest(http.MethodGet, url, nil)
		mockReq = mux.SetURLVars(mockReq, vars)

		hlp := httpHelper.NewHelper()
		params := hlp.GetParams(mockReq)

		if id != params["id"] {
			t.Error(fmt.Sprintf(ErrorGetParams, params["id"], id))
		}
	})
}

func TestHelperThrowError(t *testing.T) {
	t.Run("Teste Throw Error", func(t *testing.T) {
		// Criando nova fake request e mockRps para teste
		mockRps := httptest.NewRecorder()

		hlp := httpHelper.NewHelper()
		hlp.ThrowError(mockRps, http.StatusInternalServerError, "Error")

		if mockRps.Result().StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorOnThrowError, mockRps.Result().StatusCode)
		}
	})
}

func TestHelperThrowSuccess(t *testing.T) {
	t.Run("Teste Throw Success", func(t *testing.T) {
		// Criando nova fake request e mockRps para teste
		mockRps := httptest.NewRecorder()

		hlp := httpHelper.NewHelper()
		hlp.ThrowSuccess(mockRps, http.StatusOK, "Success")

		if mockRps.Result().StatusCode != http.StatusOK {
			t.Errorf(ErrorOnThrowSuccess, mockRps.Result().StatusCode)
		}
	})
}
