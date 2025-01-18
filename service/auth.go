package service

import (
	"net/http"
)

// AuthService - Implementação do service para autenticação.
type AuthService struct {
	// authHlp      *auth.Helper
	// validatorHlp *validator.Helper
	// stgAccount   model.AccountStorage
}

// NewAuthService - Instância o service.
func NewAuthService() *AuthService {
	return nil
}

// Login - Realiza a autenticação do usuário na API.
//
//	200: Quando a autenticação for bem sucedida.
//	401: Quando o `secret` fornecido for diferente do secret armazenado.
//	404: Quando não encontrar a account.
//	500: Erro inesperado durante o processamento da requisição
func (s *AuthService) Login(_ http.ResponseWriter, _ *http.Request) {
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
// func (s *AuthService) ValidateToken(next func(http.ResponseWriter, *http.Request, *types.Claims)) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		tokenHeader := req.Header.Get("Access-Token")
// 		claims := &types.Claims{}
//
// 		getJwtKey := func(_ *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		}
//
// 		if len(strings.Trim(tokenHeader, " ")) == 0 || tokenHeader == "null" {
// 			s.logger.Info("Error on parser token claims ")
// 			s.httpHlp.ThrowError(w, http.StatusBadRequest, InfoTokenEmpty)
//
// 			return
// 		}
//
// 		jwtToken, err := jwt.ParseWithClaims(tokenHeader, claims, getJwtKey)
// 		if err != nil {
// 			tokenMalFormed, _ := regexp.MatchString(ErrorTokenMalformed, err.Error())
//
// 			tokenSignatureInvalid, _ := regexp.MatchString(ErrorTokenSignatureInvalid, err.Error())
// 			if errors.Is(err, jwt.ErrSignatureInvalid) || tokenSignatureInvalid || tokenMalFormed {
// 				s.logger.Error("Error on parser token claims: ", err)
// 				s.httpHlp.ThrowError(w, http.StatusUnauthorized, ErrorSignatureKeyInvalid)
//
// 				return
// 			}
//
// 			s.logger.Error("Error unexpected on parser token claims: ", err)
// 			s.httpHlp.ThrowError(w, http.StatusInternalServerError, types.ErrorUnexpected)
//
// 			return
// 		}
//
// 		if !jwtToken.Valid {
// 			s.logger.Error("Error on parser token claims: ", err)
// 			s.httpHlp.ThrowError(w, http.StatusUnauthorized, ErrorTokenInvalid)
//
// 			return
// 		}
//
// 		minutesElapsedLastAuth := time.Until(time.Unix(claims.ExpiresAt, 0))
//
// 		if minutesElapsedLastAuth <= 0 {
// 			s.logger.Info("Token expired")
// 			s.httpHlp.ThrowError(w, http.StatusUnauthorized, InfoTokenExpired)
//
// 			return
// 		}
//
// 		// Invoca a próxima requisição.
// 		next(w, req, claims)
// 	}
// }
