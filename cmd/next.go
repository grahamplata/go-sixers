/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Gets the next available sixers game date and time.",
	Long:  "Gets the next available sixers game date and time.",
	Run: func(cmd *cobra.Command, args []string) {
		var url string
		year := time.Now().Format("2006")
		url = buildURL(decrementString(year), year)
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("The request failed with error %s\n", err)
		} else {
			fmt.Println(handleNextResponse(response))
		}
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
