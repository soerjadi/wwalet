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
	req.TargetUserID = userLog.ID

	result, err := u.repository.InsertTransaction(ctx, req)

	if err != nil {
		return model.TransactionSingle{}, err
	}

	go func(ctx context.Context) {
		err = u.repository.UpdateBalance(ctx, req.TargetUserID, balanceAfter)
		if err != nil {
			fmt.Printf("error updating balance err : %v\n", err)
		}
	}(ctx)

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

	if wallet.Balance < req.Amount {
		return model.TransactionSingle{}, errors.New("Balance is not enough")
	}

	balanceAfter := wallet.Balance - req.Amount
	req.BalanceAfter = balanceAfter
	req.BalanceBefore = wallet.Balance
	req.TargetUserID = userLog.ID

	result, err := u.repository.InsertTransaction(ctx, req)
	if err != nil {
		return model.TransactionSingle{}, errors.New("failed to insert transaction")
	}

	go func(ctx context.Context) {
		err = u.repository.UpdateBalance(ctx, req.TargetUserID, balanceAfter)
		if err != nil {
			fmt.Printf("error updating balance err : %v\n", err)
		}
	}(ctx)

	transform := result.TransformSingle()

	return transform, nil
}

func (u transactionUsecase) Transfer(ctx context.Context, req model.TransactionRequest) (model.TransactionSingle, error) {
	req.Category = model.TRANSACTION_CATEGORY_TRANSFER
	req.Type = model.TRANSACTION_TYPE_DEBIT
	targetUserID := req.TargetUserID

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

	go func(ctx context.Context) {
		err = u.repository.UpdateBalance(ctx, userLog.ID, req.BalanceAfter)
		if err != nil {
			fmt.Printf("got an error updating balance sender. err : %v\n", err)
			return
		}
	}(ctx)

	// do insert transfer history to target user
	go func(ctx context.Context) {
		wallet, err := u.getUserWallet(ctx, req.TargetUserID)
		if err != nil {
			fmt.Printf("failed get current user balance, %v", err)
			return
		}

		balanceAfter := wallet.Balance + req.Amount
		req.BalanceAfter = balanceAfter
		req.BalanceBefore = wallet.Balance
		req.TargetUserID = targetUserID

		result, err := u.repository.InsertTransaction(ctx, req)
		if err != nil {
			fmt.Printf("Got an error calculate transaction history to target user %v\n", err)
			return
		}

		fmt.Printf("successfully insert transactions history to target user %v\n", result)

		err = u.repository.UpdateBalance(ctx, targetUserID, balanceAfter)
		if err != nil {
			fmt.Printf("got an error updating balance target. err : %v\n", err)
			return
		}
	}(ctx)

	return transform, nil
}

func (u transactionUsecase) transferInternal(ctx context.Context, req model.TransactionRequest, userID string) (model.TransactionRequest, error) {
	wallet, err := u.getUserWallet(ctx, userID)
	if err != nil {
		return req, errors.New("failed get current user balance")
	}

	if wallet.Balance < req.Amount {
		return req, errors.New("Balance is not enough")
	}

	balanceAfter := wallet.Balance - req.Amount
	req.BalanceAfter = balanceAfter
	req.BalanceBefore = wallet.Balance
	req.TargetUserID = userID

	return req, nil
}

func (u transactionUsecase) getUserWallet(ctx context.Context, userID string) (model.Wallet, error) {
	return u.repository.GetUserWallet(ctx, userID)
}
