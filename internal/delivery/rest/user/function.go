package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/soerjadi/wwalet/internal/model"
)

func (h Handler) register(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.UserRegisterRequest{}

	if err := dec.Decode(&req); err != nil {
		return nil, errors.New("fail to decode request body")
	}

	user, err := h.usecase.Register(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h Handler) login(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	req := model.LoginRequest{}

	if err := dec.Decode(&req); err != nil {
		return nil, errors.New("fail to decode request body")
	}

	user, err := h.usecase.Login(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h Handler) update(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var userLog model.User
	dec := json.NewDecoder(r.Body)

	req := model.UserUpdateRequest{}

	if err := dec.Decode(&req); err != nil {
		return nil, errors.New("fail to decode request body")
	}

	userLogStr := r.Context().Value("user-key-respondent")
	if userLogStr != nil {
		userLog = userLogStr.(model.User)
	}

	req.User = userLog

	user, err := h.usecase.Update(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
