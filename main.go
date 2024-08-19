package main

import (
	"log"
	"os"

	"github.com/fkw3t/gotask/cmd"
	filesystem_repo "github.com/fkw3t/gotask/internal/repository/filesystem"
	"github.com/fkw3t/gotask/internal/usecase"
)

func main() {
	//// database as storage
	// db, err := sql.Open("sqlite3", "./gotask.db")
	// if err != nil {
	// 	log.Panicf("failed to open database: %v", err)
	// }
	// defer db.Close()
	// taskRepo := sqlite_repo.NewTaskRepo(db)

	//// filesystem as storage
	// TODO: set filepath as a environment variable
	currentDir, err := os.Getwd()
	if err != nil {
		log.Panicf("failed to get current directory: %v", err)
	}
	taskRepo, err := filesystem_repo.NewTaskRepo(currentDir, "gotask-storage")
	if err != nil {
		log.Panicf("failed to create task repository: %v", err)
	}

	// setup usecase
	addTaskUseCase := usecase.NewAddTaskUseCase(taskRepo)
	listTaskUseCase := usecase.NewListTaskUseCase(taskRepo)
	completeTaskUseCase := usecase.NewCompleteTaskUseCase(taskRepo)
	deleteTaskUseCase := usecase.NewDeleteTaskUseCase(taskRepo)

	// setup command
	cmd := cmd.NewRootCmd(addTaskUseCase, listTaskUseCase, completeTaskUseCase, deleteTaskUseCase)
	cmd.Execute()
}
