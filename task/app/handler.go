package app

import (
	auth "github.com/koizr/go-todo-sample/auth/usecase"
	"github.com/koizr/go-todo-sample/common"
	"github.com/koizr/go-todo-sample/infra/persistent"
	"github.com/koizr/go-todo-sample/task/domain"
	"github.com/koizr/go-todo-sample/task/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type dependencies interface {
	DB() *gorm.DB
}

type addTaskResponse struct {
	task *domain.Task
}

type addTaskForm struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
}

func AddTask(dependencies dependencies) func(c echo.Context) error {
	return func(c echo.Context) error {
		form := &addTaskForm{}
		if err := c.Bind(form); err != nil {
			return c.JSON(http.StatusBadRequest, common.Error{Message: err.Error()})
		}

		user, err := auth.Authenticate(c.Get("user").(*auth.Token))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, common.Error{Message: "invalid or expired jwt"})
		}

		dueDate, err := time.Parse(time.RFC3339, form.DueDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.Error{
				Message: "dueDate parse error. expected dueDate format is RFC3339",
			})
		}

		task, err := usecase.Add(
			&usecase.AddDependencies{
				Tasks:          &Tasks{db: dependencies.DB()},
				GenerateTaskID: persistent.GenerateTaskID,
			},
			&usecase.AddInput{
				User:        &domain.User{ID: user.ID},
				Subject:     form.Subject,
				Description: form.Description,
				DueDate:     &dueDate,
			},
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.Error{Message: "failed to add task"})
		}

		return c.JSON(http.StatusOK, &addTaskResponse{
			task: task,
		})
	}
}
