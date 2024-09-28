package model

import "time"

type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Pin         string    `json:"pin"`
	Salt        string    `json:"-"`
	CreatedAt   time.Time `json:"created_dated"`
}
