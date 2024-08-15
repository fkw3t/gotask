package cmd

import (
	"fmt"
	"github.com/fkw3t/gotask/internal/usecase"
	"os"

	"github.com/spf13/cobra"
)

type Root struct {
	addTaskUseCase      *usecase.AddTaskUseCase
	listTaskUseCase     *usecase.ListTaskUseCase
	completeTaskUseCase *usecase.CompleteTaskUseCase
	deleteTaskUsecase   *usecase.DeleteTaskUseCase
}

func NewRoot(
	addTaskUseCase *usecase.AddTaskUseCase,
	listTaskUseCase *usecase.ListTaskUseCase,
	completeTaskUseCase *usecase.CompleteTaskUseCase,
	deleteTaskUsecase *usecase.DeleteTaskUseCase,
) *Root {
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "gotask",
	Short: "Homemade task list manager",
	Long:  `A homemade task list manager built with go to practice language knowledge`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (r *Root) init() {
	rootCmd.AddCommand(r.addTaskCmd(r))
}
