package usecase

import (
	"errors"
	"github.com/koizr/go-todo-sample/auth/domain"
	"github.com/koizr/go-todo-sample/auth/jwt"
	"time"
)

type LoginDependencies interface {
	Users() domain.Users
	Secret() string
	Now() *time.Time
	AuthenticationExpire() time.Duration
}

type LoginForm struct {
	LoginID  string `json:"loginId"`
	Password string `json:"password"`
}

func Login(dep LoginDependencies, form *LoginForm) (string, error) {
	user, err := dep.Users().Find(form.LoginID, form.Password)
	if err != nil {
		return "", errors.New("user not found")
	}

	return jwt.GenerateToken(dep.Secret(), user, dep.Now(), dep.AuthenticationExpire())
}
