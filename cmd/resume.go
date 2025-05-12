package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use: "resume",
	Short: "Resumes the current focus session.",
	Long: "Resumes the current focus session.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Session resumed.")
	},
}

func init() {
	rootCmd.AddCommand(resumeCmd)
}