package app

import (
	"errors"
	"github.com/koizr/go-todo-sample/auth/domain"
	"github.com/koizr/go-todo-sample/infra/persistent"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB) domain.Users {
	return &users{
		db: db,
	}
}

func (u *users) Find(loginID string, password string) (*domain.User, error) {
	user := &persistent.User{}
	hashedPassword, err := persistent.HashPassword(password)
	if err != nil {
		return nil, err
	}
	if u.db.First(user, persistent.User{
		LoginID:  loginID,
		Password: hashedPassword,
	}).Error != nil {
		return nil, errors.New("user not found")
	}

	return &domain.User{
		ID: user.ID,
	}, nil
}
