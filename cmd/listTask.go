package cmd

import (
	"github.com/spf13/cobra"
)

var listTaskCmd = &cobra.Command{
	Use:     "List Task",
	Aliases: []string{"list"},
	Short:   "List task(s) in todo list",
	Long:    "List task(s) in todo list",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ListTask()
	},
}

func init() {
	rootCmd.AddCommand(listTaskCmd)
}
