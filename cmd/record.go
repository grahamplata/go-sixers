package cmd

import (
	"github.com/grahamplata/sixers/handlers"
	"github.com/spf13/cobra"
)

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Fetches the record for the current season.",
	Long:  "Fetches the record for the current season.",
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Record(cmd)
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)
	recordCmd.Flags().StringP("year", "y", "", "The year of the season you would like to query")
}
