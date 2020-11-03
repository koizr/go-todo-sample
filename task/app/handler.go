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
	Now() *time.Time
}

type addTaskResponse struct {
	task *domain.Task
}

type addTaskClientErrorResponse struct {
	Error addTaskClientError `json:"error"`
}

type addTaskClientError struct {
	Message string `json:"message"`
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
			return c.JSON(http.StatusBadRequest, &struct{}{}) // TODO: return message
		}

		user, err := auth.Authenticate(c.Get("user").(*auth.Token), dependencies.Now())
		if err != nil {
			// TODO: JWT の認証失敗時のレスポンスの形式を調べて再実装。しかし JWT を使っているという情報がここに漏れ出すのは良くない感じがする
			return c.JSON(http.StatusUnauthorized, nil)
		}

		dueDate, err := time.Parse(time.RFC3339, form.DueDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &addTaskClientErrorResponse{
				Error: addTaskClientError{
					Message: "dueDate parse error. expected dueDate format is RFC3339",
				},
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
			return c.JSON(http.StatusInternalServerError, common.NewServerError("failed to add task"))
		}

		return c.JSON(http.StatusOK, &addTaskResponse{
			task: task,
		})
	}
}
