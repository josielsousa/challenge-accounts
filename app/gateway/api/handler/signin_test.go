package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest"
	"github.com/josielsousa/challenge-accounts/app/types"
)

func TestHandler_Signin(t *testing.T) {
	t.Parallel()

	type fields struct {
		authUC authUsecase
	}

	type want struct {
		statusCode int
		body       json.RawMessage
	}

	type args struct {
		input SinginRequest
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "singin success",
			args: args{
				input: SinginRequest{
					Cpf:      "75811508018",
					Password: "p4s5w0rD",
				},
			},
			fields: fields{
				authUC: &authUsecaseMock{
					SigninFunc: func(_ context.Context, _ types.Credentials) (types.Auth, error) {
						return types.Auth{Token: "t0k3n"}, nil
					},
				},
			},
			want: want{
				statusCode: http.StatusOK,
				body:       json.RawMessage(`{"token":"t0k3n"}`),
			},
		},
		{
			name: "give a bad request when cpf is invalid",
			args: args{
				input: SinginRequest{
					Cpf:      "12345678900",
					Password: "p4s5w0rD",
				},
			},
			fields: fields{
				authUC: &authUsecaseMock{
					SigninFunc: func(_ context.Context, _ types.Credentials) (types.Auth, error) {
						return types.Auth{Token: "t0k3n"}, nil
					},
				},
			},
			want: want{
				statusCode: http.StatusBadRequest,
				body:       json.RawMessage(`{"code":"app:bad_request", "message":"cpf must be a valid cpf"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			hand := &Handler{
				authUC: tt.fields.authUC,
			}

			router := chi.NewRouter()
			router.Post("/signin", rest.Handler(hand.Signin))

			bodyBytes, err := json.Marshal(tt.args.input)
			require.NoError(t, err)

			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodPost,
				"/signin",
				bytes.NewReader(bodyBytes),
			)
			require.NoError(t, err)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			respBody, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			got := json.RawMessage(respBody)

			assert.Equal(t, tt.want.statusCode, resp.Code)
			assert.JSONEq(t, string(tt.want.body), string(got))
		})
	}
}
