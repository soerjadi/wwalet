package user

import "github.com/jmoiron/sqlx"

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		q   prepareQuery
		err error
	)

	q.getUserByPhoneNumber, err = db.Preparex(getUserByPhoneNumber)
	if err != nil {
		return q, err
	}

	q.updateUser, err = db.Preparex(updateUser)
	if err != nil {
		return q, err
	}

	q.registerUser, err = db.Preparex(registerUser)
	if err != nil {
		return q, err
	}

	q.getUserbyID, err = db.Preparex(getUserbyID)
	if err != nil {
		return q, err
	}

	q.createWallet, err = db.Preparex(createWallet)
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

	return &userRepository{
		query: query,
	}, nil
}
