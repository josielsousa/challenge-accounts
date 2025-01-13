package http_test

import (
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
	t.Parallel()

	t.Run("Teste Get Params com sucesso", func(t *testing.T) {
		t.Parallel()

		id := "35"
		vars := map[string]string{"id": id}

		// Criando nova fake request para teste
		url := "http://localhost:3000/accounts/" + vars["id"]

		mockReq := httptest.NewRequest(http.MethodGet, url, nil)
		mockReq = mux.SetURLVars(mockReq, vars)

		hlp := httpHelper.NewHelper()
		params := hlp.GetParams(mockReq)

		if id != params["id"] {
			t.Errorf(ErrorGetParams, params["id"], id)
		}
	})
}

func TestHelperThrowError(t *testing.T) {
	t.Parallel()

	t.Run("Teste Throw Error", func(t *testing.T) {
		t.Parallel()

		// Criando nova fake request e mockRps para teste
		rec := httptest.NewRecorder()

		hlp := httpHelper.NewHelper()
		hlp.ThrowError(rec, http.StatusInternalServerError, "Error")

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf(ErrorOnThrowError, res.StatusCode)
		}
	})
}

func TestHelperThrowSuccess(t *testing.T) {
	t.Parallel()

	t.Run("Teste Throw Success", func(t *testing.T) {
		t.Parallel()

		// Criando nova fake request e rec para teste
		rec := httptest.NewRecorder()

		hlp := httpHelper.NewHelper()
		hlp.ThrowSuccess(rec, http.StatusOK, "Success")

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf(ErrorOnThrowSuccess, res.StatusCode)
		}
	})
}
