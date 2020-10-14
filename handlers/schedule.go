package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/grahamplata/sixers/api"
	"github.com/grahamplata/sixers/config"
	"github.com/grahamplata/sixers/utils"
	"github.com/spf13/cobra"
)

// Schedule ...
func Schedule(cmd *cobra.Command, args []string) {
	// Create the parameters
	year := time.Now().Format(config.YearFormat)
	url := api.BuildURL(utils.DecrementString(year), year)

	// Do get request and handle the response
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The request failed with an error. %s\n", err)
	} else {
		resp := api.ScheduleResponse(response)
		fmt.Println(resp)
	}
}
