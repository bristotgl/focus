package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	Short: "Start a new focus session.",
	Long: "Begin a new focus session that will run for the specified duration.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting focus session.")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}