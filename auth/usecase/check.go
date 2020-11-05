package usecase

import (
	"github.com/koizr/go-todo-sample/auth/domain"
	"github.com/koizr/go-todo-sample/auth/jwt"
)

type Token = jwt.Token

func Authenticate(token *Token) (*domain.User, error) {
	userID, err := jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID: userID,
	}, nil
}
