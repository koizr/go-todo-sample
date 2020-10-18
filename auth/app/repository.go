package app

import (
	"github.com/koizr/go-todo-sample/auth/domain"
	"github.com/koizr/go-todo-sample/infra/persistent"
	"golang.org/x/crypto/bcrypt"
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
	if u.db.First(user, persistent.User{LoginID: loginID}).Error != nil {
		return nil, &domain.UserNotFoundError{}
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, &domain.UserNotFoundError{}
	}

	return &domain.User{
		ID: user.ID,
	}, nil
}
