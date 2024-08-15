package task

import (
	"github.com/fkw3t/gotask/internal/enum"
	"time"
)

type Task struct {
	Id          *uint16
	Name        string
	Description *string
	Status      enum.Status
	Deadline    *time.Time
	DueDate     *time.Time
	CreatedAt   *time.Time
}

type TaskRepo interface {
	Add(task *Task) error
	List() ([]*Task, error)
	Complete(taskId uint16) error
	Delete(taskId uint16) error
}

func NewTask(
	id *uint16,
	name string,
	description *string,
	status enum.Status,
	deadline *time.Time,
	dueDate *time.Time,
	createdAt *time.Time,
) *Task {
	return &Task{
		Id:          id,
		Name:        name,
		Description: description,
		Status:      status,
		Deadline:    deadline,
		DueDate:     dueDate,
		CreatedAt:   createdAt,
	}
}
