package app

import (
	"github.com/koizr/go-todo-sample/common"
	"github.com/koizr/go-todo-sample/infra/persistent"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type registerDep interface {
	DB() *gorm.DB
}

func Register(dependencies registerDep) func(c echo.Context) error {
	return func(c echo.Context) error {
		provisionalUser := &persistent.ProvisionalUser{}
		if err := c.Bind(provisionalUser); err != nil {
			return c.JSON(http.StatusBadRequest, &struct{}{})
		}

		user, err := persistent.NewUser(provisionalUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewServerError("failed to add user."))
		}

		if dependencies.DB().Create(user).Error != nil {
			return c.JSON(http.StatusInternalServerError, common.NewServerError("failed to add user."))
		}

		return c.JSON(http.StatusCreated, &struct{}{})
	}
}
