package service

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/types"
)

// Constantes utilizadas no serviço de autenticação.
//
//nolint:gosec
const (
	MaxTimeToExpiration        = 5
	InfoTokenEmpty             = "Token vazio."
	InfoTokenExpired           = "Token expirado."
	ErrorTokenInvalid          = "Token inválido."
	ErrorTokenMalformed        = "token is malformed"
	ErrorTokenSignatureInvalid = "token signature is invalid"
	ErrorSignatureKeyInvalid   = "A chave de assinatura do token é inválida."
)

// JWT string chave utilizada para geração do token.
var jwtKey = []byte("api-challenge-accounts")

// AuthService - Implementação do service para autenticação.
type AuthService struct {
	// authHlp      *auth.Helper
	httpHlp *httpHelper.Helper
	// validatorHlp *validator.Helper
	logger types.APILogProvider
	// stgAccount   model.AccountStorage
}

// NewAuthService - Instância o service.
func NewAuthService() *AuthService {
	// return &AuthService{
	// 	logger:       log,
	// 	stgAccount:   stgAccount,
	// 	authHlp:      auth.NewHelper(),
	// 	httpHlp:      httpHelper.NewHelper(),
	// 	validatorHlp: validator.NewHelper(),
	// }

	return nil
}

// Login - Realiza a autenticação do usuário na API.
//
//	200: Quando a autenticação for bem sucedida.
//	401: Quando o `secret` fornecido for diferente do secret armazenado.
//	404: Quando não encontrar a account.
//	500: Erro inesperado durante o processamento da requisição
func (s *AuthService) Login(w http.ResponseWriter, req *http.Request) {
	// credential := s.validatorHlp.ValidateDataLogin(w, req)
	// if credential == nil {
	// 	return
	// }

	// s.httpHlp.ThrowSuccess(w, http.StatusOK, types.SuccessResponse{Success: true, Data: authToken})
}

// ValidateToken - Realiza a validação do token enviado.
//
//	Quando a validação for bem sucedida, executa a próxima rota informada.
//	401: Quando o `token` fornecido for inválido.
//	400: Quando o token estiver vazio / nulo
//	500: Erro inesperado durante o processamento da requisição
//
// TODO: moves this method to a middleware package.
//
//nolint:varnamelen
func (s *AuthService) ValidateToken(next func(http.ResponseWriter, *http.Request, *types.Claims)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		tokenHeader := req.Header.Get("Access-Token")
		claims := &types.Claims{}

		getJwtKey := func(_ *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		}

		if len(strings.Trim(tokenHeader, " ")) == 0 || tokenHeader == "null" {
			s.logger.Info("Error on parser token claims ")
			s.httpHlp.ThrowError(w, http.StatusBadRequest, InfoTokenEmpty)

			return
		}

		jwtToken, err := jwt.ParseWithClaims(tokenHeader, claims, getJwtKey)
		if err != nil {
			tokenMalFormed, _ := regexp.MatchString(ErrorTokenMalformed, err.Error())

			tokenSignatureInvalid, _ := regexp.MatchString(ErrorTokenSignatureInvalid, err.Error())
			if errors.Is(err, jwt.ErrSignatureInvalid) || tokenSignatureInvalid || tokenMalFormed {
				s.logger.Error("Error on parser token claims: ", err)
				s.httpHlp.ThrowError(w, http.StatusUnauthorized, ErrorSignatureKeyInvalid)

				return
			}

			s.logger.Error("Error unexpected on parser token claims: ", err)
			s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)

			return
		}

		if !jwtToken.Valid {
			s.logger.Error("Error on parser token claims: ", err)
			s.httpHlp.ThrowError(w, http.StatusUnauthorized, ErrorTokenInvalid)

			return
		}

		minutesElapsedLastAuth := time.Until(time.Unix(claims.ExpiresAt, 0))

		if minutesElapsedLastAuth <= 0 {
			s.logger.Info("Token expired")
			s.httpHlp.ThrowError(w, http.StatusUnauthorized, InfoTokenExpired)

			return
		}

		// Invoca a próxima requisição.
		next(w, req, claims)
	}
}
