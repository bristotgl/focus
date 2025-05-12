package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use: "status",
	Short: "Check the current session status.",
	Long: "Displays the remaining time or completion status of the current session.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current session status is: test")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}