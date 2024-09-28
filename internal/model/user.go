package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID          string       `db:"id"`
	FirstName   string       `db:"first_name"`
	LastName    string       `db:"last_name"`
	PhoneNumber string       `db:"phone_number"`
	Address     string       `db:"address"`
	Pin         string       `db:"pin"`
	Salt        string       `db:"salt"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
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
