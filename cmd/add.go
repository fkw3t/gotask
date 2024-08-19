package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	description string
	deadline    string
)

func (r *RootCmd) addTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add item to task list",
		Args:  cobra.ExactArgs(1),
		Example: `gotask add "task name" --description "task description" --deadline "2024-12-31 16:00:00"
                  or
                  gotask add "task name" -d "task description" -l "2024-12-31 16:00:00"`,
		Run: func(cmd *cobra.Command, args []string) {
			err := r.addTaskUseCase.Add(args[0], description, deadline)
			if err != nil {
				fmt.Println("failed to add item to task list: ", err)
				os.Exit(1)
			}

			fmt.Println("Item added to task list!")
		},
	}

	cmd.Flags().StringVarP(&description, "description", "d", "", "item description")
	cmd.Flags().StringVarP(&deadline, "deadline", "l", "", "item deadline")

	return cmd
}
