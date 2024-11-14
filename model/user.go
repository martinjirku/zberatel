package model

import (
	"encoding/gob"
)

func init() {
	gob.Register(UserLogin{})
}

type UserRegistrationInput struct {
	Username string `json:"username" validate:"required,min=3,max=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserLoginInput(username, password string) UserLoginInput {
	return UserLoginInput{
		Username: username,
		Password: password,
	}
}

type UserLogin struct {
	ID       KSUID  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
