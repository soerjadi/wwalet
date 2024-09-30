package transaction

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

type prepareQueryMock struct {
	getTransactionList *sqlmock.ExpectedPrepare
	insertTransaction  *sqlmock.ExpectedPrepare
	updateBalance      *sqlmock.ExpectedPrepare
	getUserWallet      *sqlmock.ExpectedPrepare
}

func expectPrepareMock(mock sqlmock.Sqlmock) prepareQueryMock {
	prepareQuery := prepareQueryMock{}

	prepareQuery.getTransactionList = mock.ExpectPrepare(`
	SELECT
		id,
		user_id,
		status,
		type,
		category,
		amount,
		remarks,
		balance_before,
		balance_after
		created_at
	FROM
		transactions
	`)

	prepareQuery.insertTransaction = mock.ExpectPrepare(`
	INSERT INTO transactions \(
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
	\) VALUES \(
		(.*),
		(.*),
		(.*),
		(.*),
		(.*),
		(.*),
		(.*),
		(.*),
		NOW\(\)
	\) RETURNING
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
	`)

	prepareQuery.updateBalance = mock.ExpectPrepare(`
	UPDATE INTO 
		wallets 
	SET
		balance = (.*),
		updated_at = NOW\(\)
	WHERE 
		user_id = (.*)
	`)

	prepareQuery.getUserWallet = mock.ExpectPrepare(`
	SELECT
		id,
		user_id,
		balance,
		created_at,
		updated_at
	FROM 
		wallets
	WHERE
		user_id = (.*)
	LIMIT 1
	`)

	return prepareQuery
}

func TestGetRepository(t *testing.T) {

	tests := []struct {
		name     string
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     func(db *sqlx.DB) Repository
		wantErr  bool
	}{
		{
			name: "success",
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				expectPrepareMock(mock)
				expectPrepareMock(mock)
				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: func(db *sqlx.DB) Repository {
				q, _ := prepareQueries(db)

				return &transactionRepository{
					query: q,
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()

			got, err := GetRepository(db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := tt.want(db)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetRepository() = %v, want %v", got, want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}
