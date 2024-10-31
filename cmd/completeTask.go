package cmd

import (
	"github.com/spf13/cobra"
)

var completeTaskCmd = &cobra.Command{
	Use:     "Mark Task as complete",
	Aliases: []string{"complete"},
	Short:   "Mark task(s) as complete in todo list",
	Long:    "Mark task(s) as complete in todo list",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CompleteTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(completeTaskCmd)
}

