package user

import (
	"github.com/soerjadi/wwalet/internal/config"
	"github.com/soerjadi/wwalet/internal/repository/user"
)

func GetUsecase(repo user.Repository, cfg *config.Config) Usecase {
	return &userUsecase{
		repository: repo,
		cfg:        cfg,
	}
}
