package cmd

import (
	"github.com/spf13/cobra"
)

var showCompletedTask bool
var listTaskCmd = &cobra.Command{
	Use:     "List Task",
	Aliases: []string{"list"},
	Short:   "List uncompleted task(s) in todo list",
	Long:    "List uncompleted task(s) in todo list",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ListTask(showCompletedTask)
	},
}

func init() {
	listTaskCmd.Flags().BoolVarP(&showCompletedTask, "all", "a", false, "list all tasks")
	rootCmd.AddCommand(listTaskCmd)
}

