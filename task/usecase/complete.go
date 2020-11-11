package usecase

import (
	"github.com/koizr/go-todo-sample/task/domain"
	"time"
)

type CompleteDependencies struct {
	Tasks domain.Tasks
	Now   *time.Time
	User  *domain.User
}

type CompleteInput struct {
	TaskID domain.TaskID
}

func Complete(dependencies *CompleteDependencies, input *CompleteInput) error {
	task, err := dependencies.Tasks.FindById(input.TaskID, dependencies.User)
	if err != nil {
		return err
	}

	if err := dependencies.Tasks.Update(task); err != nil {
		return err
	}

	return nil
}
