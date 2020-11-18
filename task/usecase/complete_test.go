package usecase

import (
	"errors"
	"fmt"
	"github.com/koizr/go-todo-sample/task/domain"
	"reflect"
	"testing"
	"time"
)

func TestComplete__call_complete_then_target_task_is_completed(t *testing.T) {
	tasks := &tasksMock{tasks: []domain.Task{
		{
			ID:          "task1",
			User:        &domain.User{ID: "user_id"},
			Subject:     "task1_subject",
			Description: "task1_description",
			Status:      1,
			DueDate:     dateTime(2020, 1, 1, 1, 1, 1, 1),
		},
		{
			ID:          "task2",
			User:        &domain.User{ID: "user_id"},
			Subject:     "task2_subject",
			Description: "task2_description",
			Status:      1,
			DueDate:     dateTime(2020, 2, 2, 2, 2, 2, 2),
		},
		{
			ID:          "task3",
			User:        &domain.User{ID: "user_id_other"},
			Subject:     "task3_subject",
			Description: "task3_description",
			Status:      1,
			DueDate:     dateTime(3030, 3, 3, 3, 3, 3, 3),
		},
	}}

	err := Complete(
		&CompleteDependencies{
			Tasks: tasks,
			User:  &domain.User{ID: "user_id"},
		},
		&CompleteInput{
			TaskID: "task2",
		},
	)

	if err != nil {
		t.Errorf("error returned. %s", err.Error())
	}

	if !reflect.DeepEqual(
		tasks.tasks,
		[]domain.Task{
			{
				ID:          "task1",
				User:        &domain.User{ID: "user_id"},
				Subject:     "task1_subject",
				Description: "task1_description",
				Status:      1,
				DueDate:     dateTime(2020, 1, 1, 1, 1, 1, 1),
			},
			{
				ID:          "task2",
				User:        &domain.User{ID: "user_id"},
				Subject:     "task2_subject",
				Description: "task2_description",
				Status:      0, // この値だけ変わってる
				DueDate:     dateTime(2020, 2, 2, 2, 2, 2, 2),
			},
			{
				ID:          "task3",
				User:        &domain.User{ID: "user_id_other"},
				Subject:     "task3_subject",
				Description: "task3_description",
				Status:      1,
				DueDate:     dateTime(3030, 3, 3, 3, 3, 3, 3),
			},
		},
	) {
		t.Errorf("tasks is not expected")
	}
}

func TestComplete__apply_ID_that_does_not_exist_then_return_error(t *testing.T) {
	tasks := &tasksMock{tasks: []domain.Task{
		{
			ID:          "task1",
			User:        &domain.User{ID: "user_id"},
			Subject:     "task1_subject",
			Description: "task1_description",
			Status:      1,
			DueDate:     dateTime(2020, 1, 1, 1, 1, 1, 1),
		},
	}}

	err := Complete(
		&CompleteDependencies{
			Tasks: tasks,
			User:  &domain.User{ID: "user_id"},
		},
		&CompleteInput{
			TaskID: "task2",
		},
	)

	if err == nil {
		t.Errorf("error is not nil")
	}
}

func dateTime(year int, month time.Month, day int, hour int, min int, sec int, nanoSec int) *time.Time {
	t := time.Date(year, month, day, hour, min, sec, nanoSec, time.UTC)
	return &t
}

type tasksMock struct {
	tasks []domain.Task
}

func (t *tasksMock) Add(task *domain.Task) error {
	panic("implement me")
}

func (t *tasksMock) Update(task *domain.Task) error {
	var newTasks []domain.Task
	for _, oldTask := range t.tasks {
		if oldTask.ID == task.ID {
			newTasks = append(newTasks, *task)
		} else {
			newTasks = append(newTasks, oldTask)
		}
	}
	t.tasks = newTasks
	return nil
}

func (t *tasksMock) Remove(task *domain.Task) error {
	panic("implement me")
}

func (t *tasksMock) FindAll(user *domain.User) ([]*domain.Task, error) {
	panic("implement me")
}

func (t *tasksMock) FindById(id string, user *domain.User) (*domain.Task, error) {
	for _, task := range t.tasks {
		if task.ID == id && task.User.ID == user.ID {
			return &task, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("task %s is not found", id))
}
