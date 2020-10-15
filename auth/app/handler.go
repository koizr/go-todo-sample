package app

import (
	"github.com/koizr/go-todo-sample/common"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type LoginForm struct {
	LoginID  string `json:"loginId"`
	Password string `json:"password"`
}

type dependencies interface {
	DB() *gorm.DB
	Secret() string
	Now() *time.Time
}

type loginResponse struct {
	Token string `json:"token"`
}

func Login(dependencies dependencies) func(c echo.Context) error {
	return func(c echo.Context) error {
		form := &LoginForm{}
		if err := c.Bind(form); err != nil {
			return c.JSON(http.StatusBadRequest, &struct{}{})
		}

		users := NewUsers(dependencies.DB())
		user, err := users.Find(form.LoginID, form.Password)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewServerError("user not found"))
		}

		token, err := GenerateToken(dependencies.Secret(), user, dependencies.Now())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewServerError("failed to generate token"))
		}

		return c.JSON(http.StatusOK, &loginResponse{
			Token: token,
		})
	}
}
