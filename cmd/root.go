package cmd

import (
	"fmt"
	"os"

	"github.com/fkw3t/gotask/internal/usecase"

	"github.com/spf13/cobra"
)

type RootCmd struct {
	addTaskUseCase      *usecase.AddTaskUseCase
	listTaskUseCase     *usecase.ListTaskUseCase
	completeTaskUseCase *usecase.CompleteTaskUseCase
	deleteTaskUsecase   *usecase.DeleteTaskUseCase
	rootCmd             *cobra.Command
}

func NewRootCmd(
	addTaskUseCase *usecase.AddTaskUseCase,
	listTaskUseCase *usecase.ListTaskUseCase,
	completeTaskUseCase *usecase.CompleteTaskUseCase,
	deleteTaskUsecase *usecase.DeleteTaskUseCase,
) *RootCmd {
	return &RootCmd{
		addTaskUseCase:      addTaskUseCase,
		listTaskUseCase:     listTaskUseCase,
		completeTaskUseCase: completeTaskUseCase,
		deleteTaskUsecase:   deleteTaskUsecase,
		rootCmd: &cobra.Command{
			Use:   "gotask",
			Short: "Homemade task list manager",
			Long:  `A homemade task list manager built with go to practice language knowledge`,
			Run:   func(cmd *cobra.Command, args []string) {},
		},
	}
}

func (r *RootCmd) Execute() {
	r.init()
	if err := r.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (r *RootCmd) init() {
	r.rootCmd.AddCommand(r.addTaskCmd())
	r.rootCmd.AddCommand(r.listTaskCmd())
	r.rootCmd.AddCommand(r.completeTaskCmd())
	r.rootCmd.AddCommand(r.deleteTaskCmd())
}
