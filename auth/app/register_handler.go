package app

import (
	"github.com/koizr/go-todo-sample/auth/domain"
	"github.com/koizr/go-todo-sample/auth/usecase"
	"github.com/koizr/go-todo-sample/common"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type registerDep interface {
	DB() *gorm.DB
}

type registerForm struct {
	LoginID  string `json:"loginId"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type registerSuccessResponse struct {
	Registered bool `json:"registered"`
}

func Register(dependencies registerDep) func(c echo.Context) error {
	return func(c echo.Context) error {
		form := &registerForm{}
		if err := c.Bind(form); err != nil {
			return c.JSON(http.StatusBadRequest, &struct{}{})
		}

		_, err := usecase.Register(
			&usecase.RegisterDependencies{
				Users: NewUsers(dependencies.DB()),
			},
			&domain.ProvisionalUser{
				LoginID:  form.LoginID,
				Password: form.Password,
				Name:     form.Name,
			},
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewServerError("failed to add user."))
		}

		return c.JSON(http.StatusCreated, &registerSuccessResponse{Registered: true})
	}
}
