package domain

import (
	"fmt"
	"time"
)

type UserID = string

type User struct {
	ID UserID
}

type TaskID = string

type Task struct {
	ID          TaskID
	User        *User
	Subject     string
	Description string
	Status      uint
	DueDate     *time.Time
}

func NewTask(id TaskID, user *User, subject string, description string, dueDate *time.Time) *Task {
	return &Task{
		ID:          id,
		User:        user,
		Subject:     subject,
		Description: description,
		Status:      taskStatusTodo,
		DueDate:     dueDate,
	}
}

func (t *Task) Complete() *Task {
	return &Task{
		ID:          t.ID,
		User:        t.User,
		Subject:     t.Subject,
		Description: t.Description,
		Status:      taskStatusDone,
		DueDate:     t.DueDate,
	}
}

func (t *Task) ChangeSubject(subject string) *Task {
	return &Task{
		ID:          t.ID,
		User:        t.User,
		Subject:     subject,
		Description: t.Description,
		Status:      t.Status,
		DueDate:     t.DueDate,
	}
}

func Add(tasks Tasks, task *Task) error {
	return tasks.Add(task)
}

const (
	taskStatusDone = 0
	taskStatusTodo = 1
)

type Tasks interface {
	Add(task *Task) error
	Update(task *Task) error
	Remove(task *Task) error
	FindAll(user *User) ([]*Task, error)
	FindById(id TaskID, user *User) (*Task, error)
}

type TaskNotFoundError struct {
	ID TaskID
}

func (t *TaskNotFoundError) Error() string {
	return fmt.Sprintf("task is not found. ID: %s", t.ID)
}
