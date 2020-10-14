package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "An at a glance view of the Sixers NBA season.",
	Long:  `An at a glance view of the Sixers NBA season.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("schedule called")
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
