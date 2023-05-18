package repository

import (
	"errors"

	"github.com/yong509/go-fiber-playground/types"
)

func FindByCredentials(email, password string) (*types.User, error) {
	if email == "test@mail.com" && password == "test12345" {
		return &types.User{
			ID:    1,
			Email: "test@mail.com",
		}, nil
	}
	return nil, errors.New("user not found")
}
