package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (r *RootCmd) deleteTaskCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "delete",
		Short:   "dele a task item",
		Example: `gotask delete 1`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskId := formatTaskId(args[0])
			err := r.deleteTaskUsecase.Delete(taskId)
			if err != nil {
				fmt.Printf("failed to delete task item: %v\n", err)
				os.Exit(1)
			}
		},
	}
}
