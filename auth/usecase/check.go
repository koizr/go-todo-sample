package usecase

import (
	"github.com/koizr/go-todo-sample/auth/domain"
	"github.com/koizr/go-todo-sample/auth/jwt"
	"time"
)

type Token = jwt.Token

func Authenticate(token *Token, now *time.Time) (*domain.User, error) {
	userID, err := jwt.ParseToken(token, now)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID: userID,
	}, nil
}
