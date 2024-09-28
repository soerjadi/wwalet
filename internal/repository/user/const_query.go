package user

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	getUserbyID          *sqlx.Stmt
	getUserByPhoneNumber *sqlx.Stmt
	updateUser           *sqlx.Stmt
	registerUser         *sqlx.Stmt
}

const (
	getUserbyID = `
	SELECT
		id, 
		first_name,
		last_name,
		phone_number,
		address,
		pin,
		salt,
		created_at,
		updated_at
	FROM
		users
	WHERE id = $1
	LIMIT 1
	`

	getUserByPhoneNumber = `
	SELECT
		id, 
		first_name,
		last_name,
		phone_number,
		address,
		pin,
		salt,
		created_at,
		updated_at
	FROM
		users
	WHERE phone_number = $1
	LIMIT 1
	`

	updateUser = `
	UPDATE 
		users 
	SET
		first_name = $1,
		last_name = $2,
		address = $3,
		updated_at = NOW()
	WHERE
		id = $4
	`

	registerUser = `
	INSERT INTO users (
		id,
		first_name,
		last_name,
		phone_number,
		address,
		pin,
		salt
	) VALUES (
		$1,
		$2,
		$3,
		$4, 
		$5,
		$6,
		$7 
	) RETURNING (
		id,
		first_name,
		last_name,
		phone_number,
		address,
		pin,
		salt,
		created_at 
	)
	`
)
