package app

import (
	"github.com/koizr/go-todo-sample/auth/domain"
	"github.com/koizr/go-todo-sample/auth/usecase"
	"github.com/koizr/go-todo-sample/common"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type loginDep interface {
	DB() *gorm.DB
	Secret() string
	Now() *time.Time
	AuthenticationExpire() time.Duration
}

type loginResponse struct {
	Token string `json:"token"`
}

func Login(dependencies loginDep) func(c echo.Context) error {
	return func(c echo.Context) error {
		form := &usecase.LoginForm{}
		if err := c.Bind(form); err != nil {
			return c.JSON(http.StatusBadRequest, &struct{}{})
		}

		token, err := usecase.Login(
			&loginDependencies{
				dep: dependencies,
			},
			form,
		)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewServerError(err.Error()))
		}

		return c.JSON(http.StatusOK, &loginResponse{
			Token: token,
		})
	}
}

type loginDependencies struct {
	dep loginDep
}

func (ld *loginDependencies) Secret() string {
	return ld.dep.Secret()
}

func (ld *loginDependencies) Now() *time.Time {
	return ld.dep.Now()
}

func (ld *loginDependencies) Users() domain.Users {
	return NewUsers(ld.dep.DB())
}

func (ld *loginDependencies) AuthenticationExpire() time.Duration {
	return ld.dep.AuthenticationExpire()
}
