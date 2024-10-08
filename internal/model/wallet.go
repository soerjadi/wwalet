package model

import (
	"database/sql"
	"time"
)

type Wallet struct {
	ID        string       `db:"id" json:"id"`
	UserID    string       `db:"user_id" json:"user_id"`
	Balance   int64        `db:"balance" json:"balance"`
	CreatedAt time.Time    `db:"created_at" json:"-"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"-"`
}
