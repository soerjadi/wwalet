package user

import (
	"context"

	"github.com/soerjadi/wwalet/internal/model"
	"github.com/soerjadi/wwalet/internal/pkg/str"
)

func (r userRepository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (user model.User, err error) {
	err = r.query.getUserByPhoneNumber.GetContext(ctx, &user, phoneNumber)
	if err != nil {
		user = model.User{}
		return
	}

	return
}

func (r userRepository) UpdateUser(ctx context.Context, req model.UserUpdateRequest) (user model.User, err error) {
	if _, err = r.query.updateUser.ExecContext(
		ctx,
		req.FirstName,
		req.LastName,
		req.Address,
		req.User.ID,
	); err != nil {
		return
	}

	err = nil
	user = req.User
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Address = req.Address

	return
}

func (r userRepository) Register(ctx context.Context, req model.UserRegisterRequest) (user model.User, err error) {
	if err = r.query.registerUser.GetContext(
		ctx,
		&user,
		str.GenerateUUID(),
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
		req.Address,
		req.Pin,
		req.Salt,
	); err != nil {
		user = model.User{}
		return
	}

	err = nil

	return
}

func (r userRepository) GetUserByID(ctx context.Context, id string) (model.User, error) {
	var user model.User

	if err := r.query.getUserbyID.GetContext(ctx, &user, id); err != nil {
		return model.User{}, err
	}

	return user, nil

}
