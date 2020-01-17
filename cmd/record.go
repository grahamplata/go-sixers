package cmd

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Gets the 76ers record for the current season.",
	Long:  "Gets the 76ers record for the current season.",
	Run: func(cmd *cobra.Command, args []string) {
		year := cmd.Flag("year")
		var url string
		if year.Value.String() == "" {
			currentYear := time.Now().Format("2006")
			url = buildURL(decrementString(currentYear), currentYear)
		} else {
			url = buildURL(year.Value.String(), incrementString(year.Value.String()))
		}
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("The request failed with error %s\n", err)
		} else {
			fmt.Printf(handleRecordResponse(response))
		}
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)
	// Here you will define your flags and configuration settings.
	recordCmd.Flags().StringP("year", "y", "", "The year of the season you would like to query")
}
