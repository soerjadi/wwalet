package transaction

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	getTransactionList *sqlx.Stmt
	insertTransaction  *sqlx.Stmt
	updateBalance      *sqlx.Stmt
	getUserWallet      *sqlx.Stmt
}

const (
	getTransactionList = `
	SELECT
		id,
		user_id,
		status,
		type,
		category,
		amount,
		remarks,
		balance_before,
		balance_after,
		created_at
	FROM
		transactions
	`

	insertTransaction = `
	INSERT INTO transactions (
		id,
		user_id,
		status,
		type,
		category,
		amount,
		remarks,
		balance_before,
		balance_after,
		created_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		NOW()
	) RETURNING
		id,
		user_id,
		status,
		type,
		category,
		amount,
		remarks,
		balance_before,
		balance_after,
		created_at
	`

	updateBalance = `
	UPDATE wallets 
	SET
		balance = $1,
		updated_at = NOW()
	WHERE 
		user_id = $2
	`

	getUserWallet = `
	SELECT
		id,
		user_id,
		balance,
		created_at,
		updated_at
	FROM 
		wallets
	WHERE
		user_id = $1
	LIMIT 1
	`
)
