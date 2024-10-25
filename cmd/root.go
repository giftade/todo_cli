package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo_cli",
	Short: "Todo_Cli is a cli application for managing tasks in the terminal",
	Long: "Todo_Cli is a cli application for managing tasks in the terminal",
  Run: func(cmd *cobra.Command, args []string) {

	},

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing %s/n", err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


