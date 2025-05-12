package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var abortCmd = &cobra.Command{
	Use: "abort",
	Short: "Abort the current session.",
	Long: "Cancels the session without saving it to the log.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Session aborted.")
	},
}

func init() {
	rootCmd.AddCommand(abortCmd)
}