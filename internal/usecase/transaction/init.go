package transaction

import "github.com/soerjadi/wwalet/internal/repository/transaction"

func GetUsecase(repository transaction.Repository) Usecase {
	return &transactionUsecase{
		repository: repository,
	}
}
