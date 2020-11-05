package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	dueDate := time.Date(2020, 10, 20, 12, 13, 14, 15, time.UTC)
	actual := NewTask("taskID", &User{ID: "userID"}, "some task", "do something", &dueDate)

	expected := &Task{
		ID:          "taskID",
		User:        &User{ID: "userID"},
		Subject:     "some task",
		Description: "do something",
		Status:      1,
		DueDate:     &dueDate,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("New task doesn't equals expected task. expected=%#v, actual=%#v", expected, actual)
	}
}

func TestTask_ChangeSubject(t *testing.T) {
	dueDate := time.Date(2020, 10, 20, 12, 13, 14, 15, time.UTC)
	actual := NewTask("taskID", &User{ID: "userID"}, "some task", "do something", &dueDate).ChangeSubject("changed task subject")

	expected := &Task{
		ID:          "taskID",
		User:        &User{ID: "userID"},
		Subject:     "changed task subject",
		Description: "do something",
		Status:      1,
		DueDate:     &dueDate,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Failed to change task subject. expected=%#v, actual=%#v", expected, actual)
	}
}

func TestTask_Complete(t *testing.T) {
	dueDate := time.Date(2020, 10, 20, 12, 13, 14, 15, time.UTC)
	actual := NewTask("taskID", &User{ID: "userID"}, "some task", "do something", &dueDate).Complete()

	expected := &Task{
		ID:          "taskID",
		User:        &User{ID: "userID"},
		Subject:     "some task",
		Description: "do something",
		Status:      0,
		DueDate:     &dueDate,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Failed to complete task. expected=%#v, actual=%#v", expected, actual)
	}
}
