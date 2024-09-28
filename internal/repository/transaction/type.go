package transaction

import (
	"context"

	"github.com/soerjadi/wwalet/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockTransactionRepository -destination=../../mocks/transaction_repo_mock.go -source=type.go
type Repository interface {
	GetTransactionList(ctx context.Context) ([]model.Transaction, error)
	InsertTransaction(ctx context.Context, req model.TransactionRequest) (model.Transaction, error)
	UpdateBalance(ctx context.Context, userID string, balance int64) error
	GetUserWallet(ctx context.Context, userID string) (model.Wallet, error)
}

type transactionRepository struct {
	query prepareQuery
}
