package cmd

import (
	"github.com/grahamplata/sixers/handlers"
	"github.com/spf13/cobra"
)

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "An at a glance view of the Sixers NBA season.",
	Long:  `An at a glance view of the Sixers NBA season.`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Schedule(cmd)
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
	// Here you will define your flags and configuration settings.
	scheduleCmd.Flags().StringP("year", "y", "", "The year of the season you would like to query")
}
