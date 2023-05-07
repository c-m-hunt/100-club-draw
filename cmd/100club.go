package cmd

import (
	"github.com/c-m-hunt/100-club-draw/pkg/hundredClub"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showEntriesCmd)
	rootCmd.AddCommand(drawCmd)
}

var defaultEntriesFile = "entries.csv"

var showEntriesCmd = &cobra.Command{
	Use:   "entries",
	Short: "Shows the entries from file",
	Long:  `Shows all the entries and summarises the number of entries per person`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			defaultEntriesFile = args[0]
		}
		hc := hundredClub.New(defaultEntriesFile)
		hc.DisplayEntries()
		hc.DisplayEntriesSummary()
	},
}

var drawCmd = &cobra.Command{
	Use:   "draw",
	Short: "Draws the prizes",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			defaultEntriesFile = args[0]
		}
		hc := hundredClub.New(defaultEntriesFile)
		hc.DrawAndDisplay()
	},
}
