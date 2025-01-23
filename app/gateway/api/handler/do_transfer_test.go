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

	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest"
	"github.com/josielsousa/challenge-accounts/app/types"
)

func TestDoTransfer(t *testing.T) {
	t.Parallel()

	type fields struct {
		trfUC trfUsecase
	}

	type want struct {
		statusCode int
		body       json.RawMessage
	}

	type args struct {
		accountID       string
		claimsAccountID string
		authenticated   bool
		input           DoTransferRequest
	}

	tests := []struct {
		name   string
		args   args
		fields fields
		want   want
	}{
		{
			name: "do transfer success",
			args: args{
				accountID:       "acc-id-001",
				claimsAccountID: "acc-id-001",
				authenticated:   true,
				input: DoTransferRequest{
					AccountDestinationID: "acc-id-002",
					Amount:               4_50,
				},
			},
			fields: fields{
				trfUC: &trfUsecaseMock{
					DoTransferFunc: func(_ context.Context, _ transfers.TransferInput) (transfers.TransferOutput, error) {
						return transfers.TransferOutput{
							ID:                   "trf-id-001",
							AccountOriginID:      "acc-id-001",
							AccountDestinationID: "acc-id-002",
							Amount:               4_50,
							CreatedAt:            time.Date(2025, time.January, 22, 18, 0o5, 0, 0, time.UTC),
						}, nil
					},
				},
			},
			want: want{
				statusCode: http.StatusOK,
				body: json.RawMessage(`
					{
						"id": "trf-id-001",
						"account_origin_id": "acc-id-001",
						"account_destination_id": "acc-id-002",
						"amount": 450,
						"created_at": "2025-01-22T18:05:00Z"
					}
				`),
			},
		},
		{
			name: "fails when not authenticated",
			args: args{
				accountID:       "acc-id-001",
				claimsAccountID: "acc-id-001",
				authenticated:   false,
				input: DoTransferRequest{
					AccountDestinationID: "acc-id-002",
					Amount:               4_50,
				},
			},
			fields: fields{
				trfUC: &trfUsecaseMock{
					DoTransferFunc: func(_ context.Context, _ transfers.TransferInput) (transfers.TransferOutput, error) {
						return transfers.TransferOutput{}, nil
					},
				},
			},
			want: want{
				statusCode: http.StatusUnauthorized,
				body: json.RawMessage(`
					{
						"code": "app:unauthenticated",
						"message": "unauthenticated"
					}
				`),
			},
		},
		{
			name: "fails when claims account id is different from the request account id",
			args: args{
				accountID:       "acc-id-009",
				claimsAccountID: "acc-id-001",
				authenticated:   true,
				input: DoTransferRequest{
					AccountDestinationID: "acc-id-002",
					Amount:               4_50,
				},
			},
			fields: fields{
				trfUC: &trfUsecaseMock{
					DoTransferFunc: func(_ context.Context, _ transfers.TransferInput) (transfers.TransferOutput, error) {
						return transfers.TransferOutput{}, nil
					},
				},
			},
			want: want{
				statusCode: http.StatusForbidden,
				body: json.RawMessage(`
					{
						"code": "app:forbidden",
						"message": "forbidden"
					}
				`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			hand := &Handler{
				trfUC: tt.fields.trfUC,
			}

			router := chi.NewRouter()
			router.Post("/transfers/{account_id}", rest.Handler(hand.DoTransfer))

			bodyBytes, err := json.Marshal(tt.args.input)
			require.NoError(t, err)

			ctx := context.Background()

			if tt.args.authenticated {
				ctx = types.ContextWithClaims(context.Background(), types.Claims{
					AccountID: tt.args.claimsAccountID,
				})
			}

			req, err := http.NewRequestWithContext(
				ctx,
				http.MethodPost,
				"/transfers/"+tt.args.accountID,
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
