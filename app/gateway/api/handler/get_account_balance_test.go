package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest"
)

func TestGetAccountBalance(t *testing.T) {
	t.Parallel()

	type fields struct {
		accUC accUsecase
	}

	type want struct {
		statusCode int
		body       json.RawMessage
	}

	tests := []struct {
		name      string
		accountID string
		fields    fields
		want      want
	}{
		{
			name:      "get account balance success",
			accountID: "1",
			fields: fields{
				accUC: &accUsecaseMock{
					GetAccountBalanceFunc: func(_ context.Context, _ string) (int, error) {
						return 1000, nil
					},
				},
			},
			want: want{
				statusCode: http.StatusOK,
				body: json.RawMessage(`
					{
						"balance": 1000
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
			router.Get("/accounts/{account_id}/balance", rest.Handler(hand.GetAccountBalance))

			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodGet,
				fmt.Sprintf("/accounts/%s/balance", tt.accountID),
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
