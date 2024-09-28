package transaction

import (
	"github.com/soerjadi/wwalet/internal/config"
	"github.com/soerjadi/wwalet/internal/usecase/transaction"
	"github.com/soerjadi/wwalet/internal/usecase/user"
)

type Handler struct {
	usecase     transaction.Usecase
	userUsecase user.Usecase
	cfg         *config.Config
}
