package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var details bool

func (r *RootCmd) listTaskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list task items",
		Example: `gotask list
                  or
                  gotask list -a`,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			switch {
			case details:
				err = r.listTaskUseCase.ListWithDetails()
			default:
				err = r.listTaskUseCase.List()
			}

			if err != nil {
				fmt.Printf("failed to list task items: %v", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().BoolVarP(&details, "aill", "a", false, "list task items with all columns")

	return cmd
}
