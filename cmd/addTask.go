package cmd

import "github.com/spf13/cobra"

var addTaskCmd = &cobra.Command{
	Use: "Add Task",
	Aliases: []string{"add"},
	Short: "Add tasks to todo list",
	Long: "Add tasks to todo list",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		AddTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
}