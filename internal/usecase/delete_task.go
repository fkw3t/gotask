package usecase

import task "github.com/fkw3t/gotask/internal/model"

type DeleteTaskUseCase struct {
	taskRepo *task.TaskRepo
}

func (u *AddTaskUseCase) Delete() error {
	// TODO
	return nil
}
