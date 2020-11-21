package usecase

import (
	"github.com/koizr/go-todo-sample/auth/domain"
)

type RegisterDependencies struct {
	Users domain.Users
}

func Register(dep *RegisterDependencies, provisionalUser *domain.ProvisionalUser) (*domain.User, error) {
	return domain.Register(provisionalUser, dep.Users)
}
