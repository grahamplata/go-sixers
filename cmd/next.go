/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/
package cmd

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"net/http"
	"time"
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
			response := handleNextResponse(response)
			if response == true {
				sixers := fmt.Sprintf("%d%ders", Red(7), Blue(6))
				fmt.Printf("10,9 8 %s!\n", sixers)
			}
			fmt.Println("Sorry, no game tonight.")
		}
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
