package transaction

import (
	"context"

	"github.com/google/uuid"
	"github.com/soerjadi/wwalet/internal/model"
)

func (r transactionRepository) GetTransactionList(ctx context.Context) ([]model.Transaction, error) {
	var dbRes []model.Transaction

	err := r.query.getTransactionList.SelectContext(
		ctx,
		&dbRes,
	)

	if err != nil {
		return nil, err
	}

	return dbRes, nil
}

func (r transactionRepository) InsertTransaction(ctx context.Context, req model.TransactionRequest) (model.Transaction, error) {
	var (
		res model.Transaction
		err error
	)

	if err = r.query.insertTransaction.GetContext(
		ctx,
		&res,
		uuid.NewString(),
		req.TargetUserID,
		model.TRANSACTION_STATUS_SUCCESS,
		req.Type,
		req.Category,
		req.Amount,
		req.Remarks,
		req.BalanceBefore,
		req.BalanceAfter,
	); err != nil {
		return model.Transaction{}, err
	}

	return res, nil
}

func (r transactionRepository) UpdateBalance(ctx context.Context, userID string, balance int64) error {
	var err error

	if _, err = r.query.updateBalance.ExecContext(
		ctx,
		balance,
		userID,
	); err != nil {
		return err
	}

	return nil
}

func (r transactionRepository) GetUserWallet(ctx context.Context, userID string) (model.Wallet, error) {
	var wallet model.Wallet

	if err := r.query.getUserWallet.GetContext(ctx, &wallet, userID); err != nil {
		return model.Wallet{}, err
	}

	return wallet, nil
}
