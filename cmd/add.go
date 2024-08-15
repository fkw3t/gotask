package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var description string
var deadline string

func (r *Root) addTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add item to task list",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Item %s added to task list!\n", args[0])
			if description != "" {
				fmt.Printf("With description: %s", description)
			}
			if deadline != "" {
				fmt.Printf("With deadline date: %s", deadline)
			}

		},
	}

	cmd.Flags().StringVarP(&description, "description", "d", "", "item description")
	cmd.Flags().StringVarP(&deadline, "deadline", "l", "", "item deadline")

	return cmd
}
