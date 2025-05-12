package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use: "stop",
	Short: "Stop and save the current session.",
	Long: "Stops the active focus session and saves it to the session log.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ending focus session.")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}