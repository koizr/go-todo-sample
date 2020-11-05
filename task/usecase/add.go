package usecase

import (
	"github.com/koizr/go-todo-sample/task/domain"
	"time"
)

func Add(dependencies *AddDependencies, input *AddInput) (*domain.Task, error) {
	id, err := dependencies.GenerateTaskID()
	if err != nil {
		return nil, err
	}
	task := domain.NewTask(id, input.User, input.Subject, input.Description, input.DueDate)
	if err := domain.Add(dependencies.Tasks, task); err != nil {
		return nil, err
	}
	return task, nil
}

type AddDependencies struct {
	Tasks          domain.Tasks
	GenerateTaskID func() (domain.TaskID, error)
}

type AddInput struct {
	User        *domain.User
	Subject     string     `json:"subject"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
}
