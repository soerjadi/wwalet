package user

import (
	"context"

	"github.com/soerjadi/wwalet/internal/config"
	"github.com/soerjadi/wwalet/internal/model"
	"github.com/soerjadi/wwalet/internal/repository/user"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockUserUsecase -destination=../../mocks/user_usecase_mock.go -source=type.go
type Usecase interface {
	GetByID(ctx context.Context, id string) (model.User, error)
	Register(ctx context.Context, req model.UserRegisterRequest) (model.UserRegisterResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error)
	Update(ctx context.Context, req model.UserUpdateRequest) (model.UserUpdatedResponse, error)
}

type userUsecase struct {
	repository user.Repository
	cfg        *config.Config
}
