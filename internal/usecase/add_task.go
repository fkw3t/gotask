package usecase

import (
	"fmt"
	"time"

	utils "github.com/fkw3t/gotask/internal"
	"github.com/fkw3t/gotask/internal/enum"
	task "github.com/fkw3t/gotask/internal/model"
)

type AddTaskUseCase struct {
	taskRepo task.TaskRepo
}

func NewAddTaskUseCase(taskRepo task.TaskRepo) *AddTaskUseCase {
	return &AddTaskUseCase{taskRepo}
}

func (u *AddTaskUseCase) Add(name, description, deadline string) error {
	var parsedDeadline *time.Time
	var dueDate *time.Time
	createdAt := time.Now()
	fmt.Println("createdAt: ", createdAt)
	var err error

	id, err := u.taskRepo.GetNextId()
	if err != nil {
		return err
	}

	if deadline != "" {
		parsedDeadline, err = utils.ParseDate(deadline)
		if err != nil {
			return err
		}
	}

	task, err := task.NewTask(
		id,
		name,
		description,
		enum.StatusPending,
		parsedDeadline,
		dueDate,
		&createdAt,
	)
	if err != nil {
		return err
	}

	err = u.taskRepo.Add(task)
	if err != nil {
		return err
	}

	return nil
}

func (u *AddTaskUseCase) AddFromCSV() error {
	// TODO
	return nil
}
