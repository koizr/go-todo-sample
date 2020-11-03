package app

import (
	"github.com/koizr/go-todo-sample/infra/persistent"
	"github.com/koizr/go-todo-sample/task/domain"
	"gorm.io/gorm"
)

type Tasks struct {
	db *gorm.DB
}

func (t *Tasks) Add(task *domain.Task) error {
	return t.db.Create(taskToDataModel(task)).Error
}

func (t *Tasks) UpdateSubject(task *domain.Task) error {
	return t.db.Model(task).Updates(&persistent.Task{
		Subject: task.Subject,
	}).Error
}

func (t *Tasks) Remove(task *domain.Task) error {
	return t.db.Delete(task).Error
}

func (t *Tasks) FindAll(user *domain.User) ([]*domain.Task, error) {
	var taskDataModels []*persistent.Task
	if err := t.db.Where(&persistent.Task{UserID: user.ID}).Find(&taskDataModels).Error; err != nil {
		return nil, err
	}

	var tasks []*domain.Task
	for _, task := range taskDataModels {
		tasks = append(tasks, taskFromDataModel(task))
	}

	return tasks, nil
}

func (t *Tasks) FindById(id domain.TaskID, user *domain.User) (*domain.Task, error) {
	task := &persistent.Task{}
	tx := t.db.
		Where(&persistent.Task{ID: id, UserID: user.ID}).
		First(task)
	if tx.Error != nil {
		return nil, &domain.TaskNotFoundError{
			ID: id,
		}
	}
	return taskFromDataModel(task), nil
}

func taskToDataModel(task *domain.Task) *persistent.Task {
	return &persistent.Task{
		ID:          task.ID,
		UserID:      task.User.ID,
		Subject:     task.Subject,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     *task.DueDate,
	}
}

func taskFromDataModel(task *persistent.Task) *domain.Task {
	return &domain.Task{
		ID:          task.ID,
		User:        &domain.User{ID: task.UserID},
		Subject:     task.Subject,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     &task.DueDate,
	}
}
