package usecase

import (
	"fmt"
	"os"

	task "github.com/fkw3t/gotask/internal/model"
)

type CompleteTaskUseCase struct {
	taskRepo task.TaskRepo
}

func NewCompleteTaskUseCase(taskRepo task.TaskRepo) *CompleteTaskUseCase {
	return &CompleteTaskUseCase{taskRepo}
}

func (u *CompleteTaskUseCase) Complete(id uint16) error {
	exists, err := u.taskRepo.Exists(id)
	if err != nil {
		return err
	}

	if !exists {
		fmt.Println("task not found")
		os.Exit(1)
	}

	err = u.taskRepo.Complete(id)
	if err != nil {
		return err
	}

	return nil
}
