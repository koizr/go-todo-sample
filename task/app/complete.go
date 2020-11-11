package app

import (
	auth "github.com/koizr/go-todo-sample/auth/usecase"
	"github.com/koizr/go-todo-sample/common"
	"github.com/koizr/go-todo-sample/task/domain"
	"github.com/koizr/go-todo-sample/task/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type completeTaskDep interface {
	DB() *gorm.DB
	Now() *time.Time
}

type completeTaskResponse struct {
	Completed bool `json:"completed"`
}

func CompleteTask(dependencies completeTaskDep) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")

		user, err := auth.Authenticate(c.Get("user").(*auth.Token))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, common.Error{Message: "invalid or expired jwt"})
		}

		if usecase.Complete(
			&usecase.CompleteDependencies{
				Tasks: &Tasks{db: dependencies.DB()},
				Now:   dependencies.Now(),
				User: &domain.User{
					ID: user.ID,
				},
			},
			&usecase.CompleteInput{TaskID: id},
		) != nil {

			return c.JSON(http.StatusInternalServerError, common.Error{Message: "failed to complete task"})
		}

		return c.JSON(http.StatusOK, &completeTaskResponse{
			Completed: true,
		})
	}
}
