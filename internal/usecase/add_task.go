package usecase

import task "github.com/fkw3t/gotask/internal/model"

type AddTaskUseCase struct {
	taskRepo *task.TaskRepo
}

func (u *AddTaskUseCase) Add() error {
	// TODO
	return nil
}

func (u *AddTaskUseCase) AddFromCSV() error {
	// TODO
	return nil
}
