package transaction

import (
	"context"

	"github.com/soerjadi/wwalet/internal/model"
	"github.com/soerjadi/wwalet/internal/repository/transaction"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockTransactionUsecase -destination=../../mocks/transaction_usecase_mock.go -source=type.go
type Usecase interface {
	GetTransactionList(ctx context.Context) ([]model.TransactionList, error)
	Topup(ctx context.Context, req model.TransactionRequest) (model.TransactionSingle, error)
	Payment(ctx context.Context, req model.TransactionRequest) (model.TransactionSingle, error)
	Transfer(ctx context.Context, req model.TransactionRequest) (model.TransactionSingle, error)
}

type transactionUsecase struct {
	repository transaction.Repository
}
