package sqlite_repo

import (
	"database/sql"
	"github.com/fkw3t/gotask/internal/model"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db}
}

func (r *TaskRepo) Add(task *task.Task) error {
	// TODO
	return nil
}

func (r *TaskRepo) List() ([]*task.Task, error) {
	// TODO
	return nil, nil
}

func (r *TaskRepo) Complete(taskId uint16) error {
	// TODO
	return nil
}

func (r *TaskRepo) Delete(taskId uint16) error {
	// TODO
	return nil
}
