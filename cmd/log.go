package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use: "log",
	Short: "Show session history.",
	Long:  "Displays a log of all saved focus sessions.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sessions log: lorem ipsum...")
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}