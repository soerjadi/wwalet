package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/soerjadi/wwalet/internal/model"
	"github.com/soerjadi/wwalet/internal/pkg/str"

	"github.com/golang-jwt/jwt/v4"
)

func (u userUsecase) Register(ctx context.Context, req model.UserRegisterRequest) (model.UserRegisterResponse, error) {

	salt := str.GenerateSalt()
	pin := fmt.Sprintf("%s///%s", req.Pin, salt)

	hashed, err := str.HashStr(pin)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	req.Pin = hashed

	res, err := u.repository.Register(ctx, req)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	return model.UserRegisterResponse{
		ID:          res.ID,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		PhoneNumber: res.PhoneNumber,
		Address:     res.Address,
		CreatedAt:   res.CreatedAt.Format("2006-02-01 22:21:20"),
	}, nil
}

func (u userUsecase) Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) {
	user, err := u.repository.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		err = errors.New("User not found")
		return model.LoginResponse{}, err
	}

	pinRaw := fmt.Sprintf("%s///%s", req.Pin, user.Salt)
	passwordMatch := str.CompareHash(user.Pin, pinRaw)
	if !passwordMatch {
		return model.LoginResponse{}, errors.New("Phone Number and PIN doesn't match")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"createTime": time.Now(),
		"exp":        time.Now().Add(time.Minute * time.Duration(u.cfg.Server.JwtTTL)).Unix(),
	})

	accessToken, err := token.SignedString([]byte(u.cfg.Server.SecretKey))

	if err != nil {
		return model.LoginResponse{}, errors.New("failed generate access token")
	}

	refreshToken, err := str.RandStr(256)

	if err != nil {
		return model.LoginResponse{}, errors.New("failed generate refresh token")
	}

	return model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u userUsecase) Update(ctx context.Context, req model.UserUpdateRequest) (model.UserUpdatedResponse, error) {
	res, err := u.repository.UpdateUser(ctx, req)
	if err != nil {
		return model.UserUpdatedResponse{}, errors.New("failed updated user")
	}

	user, err := u.repository.GetUserByPhoneNumber(ctx, req.User.PhoneNumber)
	if err != nil {
		return model.UserUpdatedResponse{}, errors.New("failed to get updated user")
	}

	return model.UserUpdatedResponse{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Address:   res.Address,
		UpdatedAt: user.UpdatedAt.Time.Format("2006-02-01 22:21:20"),
	}, nil
}

func (u userUsecase) GetByID(ctx context.Context, id string) (model.User, error) {
	return u.repository.GetUserByID(ctx, id)
}
