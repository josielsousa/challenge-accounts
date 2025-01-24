package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/domain/accounts"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest"
)

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	type fields struct {
		accUC accUsecase
	}

	type want struct {
		statusCode int
		body       json.RawMessage
	}

	type args struct {
		input CreateAccountRequest
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "create account success",
			args: args{
				input: CreateAccountRequest{
					Name:     "Joshep",
					Balance:  10_00,
					CPF:      "75811508018",
					Password: "p4s5w0rD",
				},
			},
			fields: fields{
				accUC: &accUsecaseMock{
					CreateFunc: func(_ context.Context, _ accounts.AccountInput) (accounts.AccountOutput, error) {
						return accounts.AccountOutput{
							ID:        "1",
							Name:      "Joshep",
							Balance:   10_00,
							CPF:       "75811508018",
							CreatedAt: time.Date(2025, time.January, 22, 18, 0o5, 0, 0, time.UTC),
							UpdatedAt: time.Date(2025, time.January, 22, 18, 0o5, 0, 0, time.UTC),
						}, nil
					},
				},
			},
			want: want{
				statusCode: http.StatusCreated,
				body: json.RawMessage(`
					{
						"id": "1",
						"name": "Joshep",
						"balance": 1000,
						"cpf": "75811508018",
						"created_at": "2025-01-22T18:05:00Z"
					}
				`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			hand := &Handler{
				accUC: tt.fields.accUC,
			}

			router := chi.NewRouter()
			router.Post("/accounts", rest.Handler(hand.CreateAccount))

			bodyBytes, err := json.Marshal(tt.args.input)
			require.NoError(t, err)

			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodPost,
				"/accounts",
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
