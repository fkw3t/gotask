package usecase

import task "github.com/fkw3t/gotask/internal/model"

type DeleteTaskUseCase struct {
	taskRepo task.TaskRepo
}

func NewDeleteTaskUseCase(taskRepo task.TaskRepo) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{taskRepo}
}

func (u *DeleteTaskUseCase) Delete(taskId uint16) error {
	err := u.taskRepo.Delete(taskId)
	if err != nil {
		return err
	}

	return nil
}
