package user

import (
	"github.com/soerjadi/wwalet/internal/config"
	"github.com/soerjadi/wwalet/internal/usecase/user"
)

type Handler struct {
	usecase user.Usecase
	cfg     *config.Config
}
