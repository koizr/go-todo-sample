package usecase

import (
	"errors"
	"github.com/koizr/go-todo-sample/auth/domain"
)

type user struct {
	ID       string
	LoginID  string
	Password string
	Name     string
}

type usersMock struct {
	users []user
}

func (u *usersMock) Find(loginID string, password string) (*domain.User, error) {
	for _, user := range u.users {
		if user.LoginID == loginID && user.Password == password {
			return &domain.User{
				ID: user.ID,
			}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (u *usersMock) Add(provisionalUser *domain.ProvisionalUser) (*domain.User, error) {
	u.users = append(u.users, user{
		ID:       "newId",
		LoginID:  provisionalUser.LoginID,
		Password: provisionalUser.Password,
		Name:     provisionalUser.Name,
	})
	return &domain.User{ID: "new_id"}, nil
}
