package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "utsukushii",
	Short: "solution for application testing report",
	Long:  `utsukushii is a utility for show your unit-test reports in beautiful format.`,
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}