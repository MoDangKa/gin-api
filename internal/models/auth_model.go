package models

type Auth struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
