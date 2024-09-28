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

	if err = r.createWallet(ctx, user.ID); err != nil {
		return user, err
	}

	err = nil

	return
}

func (r userRepository) createWallet(ctx context.Context, userID string) error {
	var strID string
	if err := r.query.createWallet.GetContext(
		ctx,
		&strID,
		str.GenerateUUID(),
		userID,
	); err != nil {
		return err
	}

	return nil
}

func (r userRepository) GetUserByID(ctx context.Context, id string) (model.User, error) {
	var user model.User

	if err := r.query.getUserbyID.GetContext(ctx, &user, id); err != nil {
		return model.User{}, err
	}

	return user, nil

}
