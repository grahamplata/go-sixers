package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/grahamplata/sixers/api"
	"github.com/grahamplata/sixers/utils"
	"github.com/spf13/cobra"
)

// Next handles the
func Next(cmd *cobra.Command, args []string) {
	var url string
	year := time.Now().Format("2006")
	url = api.BuildURL(utils.DecrementString(year), year)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The request failed with error %s\n", err)
	} else {
		response := api.NextResponse(response)
		if response != true {
			fmt.Println("Sorry, no game tonight.")
		}
	}
}
