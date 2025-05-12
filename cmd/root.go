package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "focus",
	Short: "No noise. Just focus.",
	Long: "A simple CLI tool to start, track, and log focused sessions.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, focus world.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops, an error occurred while executing focus '%s'\n", err)
		os.Exit(1)
	}
}