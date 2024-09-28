package model

import "time"

type Transaction struct {
	ID            string    `json:"-"`
	UserID        string    `json:"user_id"`
	Status        string    `json:"status"`
	Type          string    `json:"transaction_type"`
	Category      string    `json:"-"`
	Amount        int64     `json:"amount"`
	Remarks       string    `json:"remarks,omitempty"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedAt     time.Time `json:"created_date"`
}

type TransactionRequest struct {
	TargetUserID  string `json:"target_user,omitempty"`
	Remarks       string `json:"remarks,omitempty"`
	Amount        int64  `json:"amount"`
	Category      string `json:"-"`
	Type          string `json:"-"`
	BalanceAfter  int64  `json:"-"`
	BalanceBefore int64  `json:"-"`
}

type TransactionResult struct {
	ID            string
	Amount        int64
	BalanceAfter  int64
	BalanceBefore int64
	Remarks       string
	CreatedAt     time.Time
}

type TransactionSingle struct {
	// define id key on the response
	TopupID    string `json:"top_up_id,omitempty"`
	TransferID string `json:"transfer_id,omitempty"`
	PaymentID  string `json:"payment_id,omitempty"`

	// define amount on the response
	Amount      int64 `json:"amount,omitempty"`
	AmountTopup int64 `json:"amount_top_up,omitempty"`

	Remarks       string    `json:"remarks,omitempty"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedAt     time.Time `json:"created_date"`
}

type TransactionList struct {
	// define id key on the response
	TopupID    string `json:"top_up_id,omitempty"`
	TransferID string `json:"transfer_id,omitempty"`
	PaymentID  string `json:"payment_id,omitempty"`

	Status          string    `json:"status"`
	TransactionType string    `json:"transaction_type"`
	Amount          int64     `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   int64     `json:"balance_before"`
	BalanceAfter    int64     `json:"balance_after"`
	CreatedAt       time.Time `json:"created_date"`
}

const (
	TRANSACTION_TYPE_DEBIT  = "debit"
	TRANSACTION_TYPE_CREDIT = "credit"

	TRANSACTION_STATUS_SUCCESS = "success"
	TRANSACTION_STATUS_FAILED  = "failed"

	TRANSACTION_CATEGORY_TOPUP    = "topup"
	TRANSACTION_CATEGORY_PAYMENT  = "payment"
	TRANSACTION_CATEGORY_TRANSFER = "transfer"
)

func (t Transaction) TransformSingle() TransactionSingle {
	result := TransactionSingle{
		Remarks:       t.Remarks,
		BalanceAfter:  t.BalanceAfter,
		BalanceBefore: t.BalanceBefore,
		CreatedAt:     t.CreatedAt,
	}

	switch t.Category {
	case TRANSACTION_CATEGORY_TRANSFER:
		result.TransferID = t.ID
		result.Amount = t.Amount
	case TRANSACTION_CATEGORY_TOPUP:
		result.TopupID = t.ID
		result.AmountTopup = t.Amount
	default:
		result.PaymentID = t.ID
		result.Amount = t.Amount
	}

	return result
}

func (Transaction) TransformList(data []Transaction) []TransactionList {
	var result []TransactionList

	for _, d := range data {
		transformData := d.TransformSingle()
		result = append(result, transformData)
	}

	return result
}
