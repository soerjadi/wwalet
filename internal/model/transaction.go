package model

import "time"

type Transaction struct {
	ID            string    `db:"id" json:"-"`
	UserID        string    `db:"user_id" json:"user_id"`
	Status        string    `db:"status" json:"status"`
	Type          string    `db:"type" json:"transaction_type"`
	Category      string    `db:"category" json:"-"`
	Amount        int64     `db:"amount" json:"amount"`
	Remarks       string    `db:"remarks" json:"remarks,omitempty"`
	BalanceBefore int64     `db:"balance_before" json:"balance_before"`
	BalanceAfter  int64     `db:"balance_after" json:"balance_after"`
	CreatedAt     time.Time `db:"created_at" json:"created_date"`
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

	Remarks       string `json:"remarks,omitempty"`
	BalanceBefore int64  `json:"balance_before"`
	BalanceAfter  int64  `json:"balance_after"`
	CreatedAt     string `json:"created_date"`
}

type TransactionList struct {
	// define id key on the response
	TopupID    string `json:"top_up_id,omitempty"`
	TransferID string `json:"transfer_id,omitempty"`
	PaymentID  string `json:"payment_id,omitempty"`

	Status          string `json:"status"`
	TransactionType string `json:"transaction_type"`
	Amount          int64  `json:"amount"`
	Remarks         string `json:"remarks"`
	BalanceBefore   int64  `json:"balance_before"`
	BalanceAfter    int64  `json:"balance_after"`
	CreatedAt       string `json:"created_date"`
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
		CreatedAt:     t.CreatedAt.Format("2006-09-01 22:20:01"),
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
		transactionList := TransactionList{
			Status:          d.Status,
			TransactionType: d.Type,
			Amount:          d.Amount,
			BalanceBefore:   d.BalanceBefore,
			BalanceAfter:    d.BalanceAfter,
			CreatedAt:       d.CreatedAt.Format("2006-09-01 22:20:21"),
		}

		switch d.Category {
		case TRANSACTION_CATEGORY_TRANSFER:
			transactionList.TransferID = d.ID
		case TRANSACTION_CATEGORY_TOPUP:
			transactionList.TopupID = d.ID
		default:
			transactionList.PaymentID = d.ID
		}

		result = append(result, transactionList)
	}

	return result
}
