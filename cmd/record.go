package cmd

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/

import (
	"github.com/grahamplata/sixers/handlers"
	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Gets the 76ers record for the current season.",
	Long:  "Gets the 76ers record for the current season.",
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Record(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)
	// Here you will define your flags and configuration settings.
	recordCmd.Flags().StringP("year", "y", "", "The year of the season you would like to query")
}
