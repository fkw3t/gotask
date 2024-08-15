package usecase

import task "github.com/fkw3t/gotask/internal/model"

type CompleteTaskUseCase struct {
	taskRepo *task.TaskRepo
}

func (u *CompleteTaskUseCase) Complete() error {
	// TODO
	return nil
}
