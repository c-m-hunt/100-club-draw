package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// Set at build time
var Version = "development"

// Set at build time
var CommitHash = "not set"

// Set at build time
var BuildDate = "not set"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hundred Club CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Devbox Version: %s\nCommit: %s\nBuild Date: %s", Version, CommitHash, BuildDate)
	},
}