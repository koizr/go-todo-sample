package usecase

import (
	"github.com/koizr/go-todo-sample/task/domain"
)

type CompleteDependencies struct {
	Tasks domain.Tasks
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

	if err := dependencies.Tasks.Update(task.Complete()); err != nil {
		return err
	}

	return nil
}
