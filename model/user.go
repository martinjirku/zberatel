package model

import "github.com/segmentio/ksuid"

type UserRegistrationInput struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLogin struct {
	ID       ksuid.KSUID `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
}
