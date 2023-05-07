package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "100club",
		Short: `A CLI tool for running 100 Club draw`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
