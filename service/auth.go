package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/josielsousa/challenge-accounts/helpers/auth"
	httpHelper "github.com/josielsousa/challenge-accounts/helpers/http"
	"github.com/josielsousa/challenge-accounts/repo/model"
	"github.com/josielsousa/challenge-accounts/types"
)

//Constantes utilizadas no serviço de autenticação.
const (
	MaxTimeToExpiration      = 3
	InfoTokenExpired         = "Token expirado."
	ErrorTokenInvalid        = "Token inválido."
	ErrorSignatureKeyInvalid = "A chave de assinatura do token é inválida."
)

// JWT string chave utilizada para geração do token.
var jwtKey = []byte("api-challenge-accounts")

// AuthService - Implementação do service para autenticação.
type AuthService struct {
	authHlp    *auth.Helper
	httpHlp    *httpHelper.Helper
	logger     types.APILogProvider
	stgAccount model.AccountStorage
}

//NewAuthService - Instância o service.
func NewAuthService(stgAccount model.AccountStorage, log types.APILogProvider) *AuthService {
	return &AuthService{
		logger:     log,
		stgAccount: stgAccount,
		authHlp:    auth.NewHelper(),
		httpHlp:    httpHelper.NewHelper(),
	}
}

//Login - Realiza a autenticação do usuário na API.
// 	200: Quando a autenticação for bem sucedida.
//	401: Quando o `secret` fornecido for diferente do secret armazenado.
//	404: Quando não encontrar a account.
//	500: Erro inesperado durante o processamento da requisição
func (s *AuthService) Login(w http.ResponseWriter, req *http.Request) {
	credential, err := s.JSONDecoder(req)
	if err != nil {
		s.logger.Error("Error on decode credentials: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	acc, err := s.stgAccount.GetAccountByCPF(credential.CPF)
	if err != nil {
		s.logger.Error("Error on recovery account by `cpf`: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	successAcc := acc != nil && len(acc.ID) > 0
	if !successAcc {
		s.httpHlp.ThrowError(w, http.StatusNotFound, types.ErroAccountNotFound)
		return
	}

	//Verifica se o secret informado na autenticação, é o mesmo armazenado na `account`.
	err = s.authHlp.VerifySecret(acc.Secret, credential.Secret)
	if err != nil {
		s.logger.Error("Error on recovery account by `cpf`: ", err)
		s.httpHlp.ThrowError(w, http.StatusUnauthorized, types.ErrorUnauthorized)
		return
	}

	//Tempo de expiração do token
	expirationTime := time.Now().Add(MaxTimeToExpiration * time.Minute)

	//JWT Claims - Payload contendo o CPF do usuário e a data de expiração do token
	claims := &types.Claims{
		AccountID: acc.ID,
		Username:  acc.Cpf,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		s.logger.Error("Error on generate new token by `cpf`: ", err)
		s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
		return
	}

	authToken := &types.Auth{Token: tokenString}
	s.httpHlp.ThrowSuccess(w, http.StatusOK, types.SuccessResponse{Success: true, Data: authToken})
}

//ValidateToken - Realiza a validação do token enviado.
// 	200: Quando a validação for bem sucedida.
//	401: Quando o `token` fornecido for inválido.
//	500: Erro inesperado durante o processamento da requisição
func (s *AuthService) ValidateToken(next func(http.ResponseWriter, *http.Request, *types.Claims)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		tokenHeader := req.Header.Get("Access-Token")
		claims := &types.Claims{}
		getJwtKey := func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		}

		jwtToken, err := jwt.ParseWithClaims(tokenHeader, claims, getJwtKey)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
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

		minutesToExpiration := MaxTimeToExpiration * time.Minute
		minutesElapsedLastAuth := time.Unix(claims.ExpiresAt, 0).Sub(time.Now())
		if minutesElapsedLastAuth > minutesToExpiration {
			s.logger.Info("Token expired: ")
			s.httpHlp.ThrowError(w, http.StatusUnauthorized, InfoTokenExpired)
			return
		}

		//Invoca a próxima requisição.
		next(w, req, claims)
	}
}

//JSONDecoder - Realiza o parser do body recebido da request.
func (s *AuthService) JSONDecoder(req *http.Request) (types.Credentials, error) {
	credential := types.Credentials{}
	err := json.NewDecoder(req.Body).Decode(&credential)
	return credential, err
}
