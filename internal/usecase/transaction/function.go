package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/soerjadi/wwalet/internal/model"
)

func (u transactionUsecase) GetTransactionList(ctx context.Context) ([]model.TransactionList, error) {
	transactions, err := u.repository.GetTransactionList(ctx)
	if err != nil {
		return nil, err
	}

	result := new(model.Transaction).TransformList(transactions)

	return result, nil
}

func (u transactionUsecase) Topup(ctx context.Context, req model.TransactionRequest) (model.TransactionSingle, error) {
	req.Category = model.TRANSACTION_CATEGORY_TOPUP
	req.Type = model.TRANSACTION_TYPE_CREDIT

	var userLog model.User
	userLogStr := ctx.Value("user-key-respondent")
	if userLogStr != nil {
		userLog = userLogStr.(model.User)
	}

	wallet, err := u.getUserWallet(ctx, userLog.ID)
	if err != nil {
		return model.TransactionSingle{}, errors.New("failed get current user balance")
	}

	balanceAfter := wallet.Balance + req.Amount
	req.BalanceBefore = wallet.Balance
	req.BalanceAfter = balanceAfter

	result, err := u.repository.InsertTransaction(ctx, req)

	if err != nil {
		return model.TransactionSingle{}, err
	}

	transform := result.TransformSingle()

	return transform, nil
}

func (u transactionUsecase) Payment(ctx context.Context, req model.TransactionRequest) (model.TransactionSingle, error) {
	req.Category = model.TRANSACTION_CATEGORY_PAYMENT
	req.Type = model.TRANSACTION_TYPE_DEBIT

	var userLog model.User
	userLogStr := ctx.Value("user-key-respondent")
	if userLogStr != nil {
		userLog = userLogStr.(model.User)
	}

	wallet, err := u.getUserWallet(ctx, userLog.ID)
	if err != nil {
		return model.TransactionSingle{}, errors.New("failed get current user balance")
	}

	balanceAfter := wallet.Balance - req.Amount
	req.BalanceAfter = balanceAfter
	req.BalanceBefore = wallet.Balance

	result, err := u.repository.InsertTransaction(ctx, req)
	transform := result.TransformSingle()

	return transform, nil
}

func (u transactionUsecase) Transfer(ctx context.Context, req model.TransactionRequest) (model.TransactionSingle, error) {
	req.Category = model.TRANSACTION_CATEGORY_TRANSFER
	req.Type = model.TRANSACTION_TYPE_DEBIT

	var userLog model.User
	userLogStr := ctx.Value("user-key-respondent")
	if userLogStr != nil {
		userLog = userLogStr.(model.User)
	}

	req, err := u.transferInternal(ctx, req, userLog.ID)
	if err != nil {
		return model.TransactionSingle{}, err
	}

	result, err := u.repository.InsertTransaction(ctx, req)
	transform := result.TransformSingle()

	// do insert transfer history to target user
	go func() {
		req, err = u.transferInternal(ctx, req, req.TargetUserID)
		if err != nil {
			fmt.Printf("Got an error calculate transaction history to target user %v\n", err)
			return
		}

		result, err := u.repository.InsertTransaction(ctx, req)
		if err != nil {
			fmt.Printf("Got an error calculate transaction history to target user %v\n", err)
			return
		}

		fmt.Printf("successfully insert transactions history to target user %v\n", result)
	}()

	return transform, nil
}

func (u transactionUsecase) transferInternal(ctx context.Context, req model.TransactionRequest, userID string) (model.TransactionRequest, error) {
	wallet, err := u.getUserWallet(ctx, userID)
	if err != nil {
		return req, errors.New("failed get current user balance")
	}

	balanceAfter := wallet.Balance - req.Amount
	req.BalanceAfter = balanceAfter
	req.BalanceBefore = wallet.Balance

	return req, nil
}

func (u transactionUsecase) getUserWallet(ctx context.Context, userID string) (model.Wallet, error) {
	return u.repository.GetUserWallet(ctx, userID)
}
