package cmd

import (
	"github.com/grahamplata/sixers/handlers"
	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Fetches the record for the current season.",
	Long:  "Fetches the record for the current season.",
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Record(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)
	// Here you will define your flags and configuration settings.
	recordCmd.Flags().StringP("year", "y", "", "The year of the season you would like to query")
}
