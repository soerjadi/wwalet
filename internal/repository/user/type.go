package user

import (
	"context"

	"github.com/soerjadi/wwalet/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockUserRepository -destination=../../mocks/user_repo_mock.go -source=type.go
type Repository interface {
	GetUserByID(ctx context.Context, id string) (model.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (user model.User, err error)
	UpdateUser(ctx context.Context, req model.UserUpdateRequest) (user model.User, err error)
	Register(ctx context.Context, req model.UserRegisterRequest) (user model.User, err error)
}

type userRepository struct {
	query prepareQuery
}
