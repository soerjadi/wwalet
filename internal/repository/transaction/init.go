package transaction

import "github.com/jmoiron/sqlx"

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		q   prepareQuery
		err error
	)

	q.getTransactionList, err = db.Preparex(getTransactionList)
	if err != nil {
		return q, err
	}

	q.insertTransaction, err = db.Preparex(insertTransaction)
	if err != nil {
		return q, err
	}

	q.updateBalance, err = db.Preparex(updateBalance)
	if err != nil {
		return q, err
	}

	return q, nil
}

func GetRepository(db *sqlx.DB) (Repository, error) {
	query, err := prepareQueries(db)
	if err != nil {
		return nil, err
	}

	return &transactionRepository{
		query: query,
	}, nil
}
