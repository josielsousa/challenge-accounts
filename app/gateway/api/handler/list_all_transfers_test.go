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

	"github.com/josielsousa/challenge-accounts/app/domain/transfers"
	"github.com/josielsousa/challenge-accounts/app/gateway/api/rest"
	"github.com/josielsousa/challenge-accounts/app/types"
)

func TestListAllTransfers(t *testing.T) {
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
	}

	tests := []struct {
		name   string
		args   args
		fields fields
		want   want
	}{
		{
			name: "list all transfers success",
			args: args{
				accountID:       "acc-id-001",
				claimsAccountID: "acc-id-001",
				authenticated:   true,
			},
			fields: fields{
				trfUC: &trfUsecaseMock{
					DoTransferFunc: func(_ context.Context, _ transfers.TransferInput) (transfers.TransferOutput, error) {
						return transfers.TransferOutput{
							ID:                   "trf-id-001",
							AccountOriginID:      "acc-id-001",
							AccountDestinationID: "acc-id-002",
							Amount:               4_50,
							CreatedAt:            time.Date(2025, time.January, 22, 18, 5, 0, 0, time.UTC),
						}, nil
					},
					ListTransfersAccountFunc: func(_ context.Context, _ string) ([]transfers.TransferOutput, error) {
						return []transfers.TransferOutput{
							{
								ID:                   "trf-id-001",
								AccountOriginID:      "acc-id-001",
								AccountDestinationID: "acc-id-002",
								Amount:               4_50,
								CreatedAt:            time.Date(2025, time.January, 22, 18, 5, 0, 0, time.UTC),
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
								"id": "trf-id-001",
								"account_origin_id": "acc-id-001",
								"account_destination_id": "acc-id-002",
								"amount": 450,
								"created_at": "2025-01-22T18:05:00Z"
							}
						],
						"success": true
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
			router.Get("/transfers/{account_id}", rest.Handler(hand.ListTransfers))

			ctx := context.Background()

			if tt.args.authenticated {
				ctx = types.ContextWithClaims(context.Background(), types.Claims{
					AccountID: tt.args.claimsAccountID,
				})
			}

			req, err := http.NewRequestWithContext(
				ctx,
				http.MethodGet,
				"/transfers/"+tt.args.accountID,
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
