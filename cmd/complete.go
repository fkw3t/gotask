package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func (r *RootCmd) completeTaskCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "complete",
		Short:   "complete a task item",
		Example: `gotask complete 1`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskId := formatTaskId(args[0])
			err := r.completeTaskUseCase.Complete(taskId)
			if err != nil {
				fmt.Printf("failed to complete task item: %v\n", err)
				os.Exit(1)
			}
		},
	}
}

func formatTaskId(arg string) uint16 {
	taskId, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("task id must be a number")
		os.Exit(1)
	}

	if taskId == 0 {
		fmt.Println("task id must be greater than 0")
		os.Exit(1)
	}

	return uint16(taskId)
}
