package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/grahamplata/sixers/api"
	"github.com/grahamplata/sixers/utils"
	"github.com/spf13/cobra"
)

// Schedule ...
func Schedule(cmd *cobra.Command, args []string) {
	// Create the parameters
	year := cmd.Flag("year")
	var url string
	if year.Value.String() == "" {
		currentYear := time.Now().Format("2006")
		url = api.BuildURL(utils.DecrementString(currentYear), currentYear)
	} else {
		url = api.BuildURL(year.Value.String(), utils.IncrementString(year.Value.String()))
	}
	// Do get request and handle the response
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The request failed with an error. %s\n", err)
	} else {
		resp := api.ScheduleResponse(response)
		fmt.Println(resp)
	}
}
