package handler

import (
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

func TestListAccounts(t *testing.T) {
	t.Parallel()

	type fields struct {
		accUC accUsecase
	}

	type want struct {
		statusCode int
		body       json.RawMessage
	}

	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "list accounts success",
			fields: fields{
				accUC: &accUsecaseMock{
					GetAllAccountsFunc: func(_ context.Context) ([]accounts.AccountOutput, error) {
						return []accounts.AccountOutput{
							{
								Name:      "Joshep",
								Balance:   1000,
								CPF:       "75811508018",
								CreatedAt: time.Date(2025, time.January, 22, 18, 0o5, 0, 0, time.UTC),
							},
						}, nil
					},
				},
			},
			want: want{
				statusCode: http.StatusOK,
				body: json.RawMessage(`
					{
						"data": [
							{
								"name": "Joshep",
								"balance": 1000,
								"cpf": "75811508018",
								"created_at": "2025-01-22T18:05:00Z"
							}
						],
						"success": true
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
			router.Get("/accounts", rest.Handler(hand.ListAccounts))

			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodGet,
				"/accounts",
				nil,
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
