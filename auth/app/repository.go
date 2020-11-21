package app

import (
	"github.com/google/uuid"
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

func (u *users) Add(provisionalUser *domain.ProvisionalUser) (*domain.User, error) {
	user, err := newUser(provisionalUser)
	if err != nil {
		return nil, err
	}

	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return &domain.User{ID: user.ID}, nil
}

func newUser(provisionalUser *domain.ProvisionalUser) (*persistent.User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	password, err := hashPassword(provisionalUser.Password)
	if err != nil {
		return nil, err
	}

	return &persistent.User{
		ID:       id.String(),
		LoginID:  provisionalUser.LoginID,
		Password: password,
		Name:     provisionalUser.Name,
	}, nil
}

func hashPassword(rawPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
