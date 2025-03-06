package models

import "time"

type User struct {
	ID                   int       `json:"id"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	Name                 string    `json:"name"`
	Photo                string    `json:"photo"`
	Role                 string    `json:"role"`
	Active               bool      `json:"active"`
	PasswordResetToken   string    `json:"password_reset_toke"`
	PasswordResetExpires time.Time `json:"password_reset_expires"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
	Role  string `json:"role"`
}
