package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID          string       `json:"id"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	PhoneNumber string       `json:"phone_number"`
	Address     string       `json:"address"`
	Pin         string       `json:"pin"`
	Salt        string       `json:"-"`
	CreatedAt   time.Time    `json:"created_dated"`
	UpdatedAt   sql.NullTime `json:"-"`
}

type UserRegisterResponse struct {
	ID          string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	CreatedAt   string `json:"created_date"`
}

type UserUpdatedResponse struct {
	ID        string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	UpdatedAt string `json:"updated_date"`
}

type UserUpdateRequest struct {
	User      User
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

type UserRegisterRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Pin         string `json:"pin"`
	Salt        string `json:"-"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"pin"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
