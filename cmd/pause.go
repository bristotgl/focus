package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use: "pause",
	Short: "Pauses the current focus session.",
	Long: "Pauses the current focus session.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Session paused.")
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}