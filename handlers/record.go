package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/grahamplata/sixers/api"
	"github.com/grahamplata/sixers/utils"
	"github.com/spf13/cobra"
)

// Record ...
func Record(cmd *cobra.Command, args []string) {
	year := cmd.Flag("year")
	var url string
	if year.Value.String() == "" {
		currentYear := time.Now().Format("2006")
		url = api.BuildURL(utils.DecrementString(currentYear), currentYear)
	} else {
		url = api.BuildURL(year.Value.String(), utils.IncrementString(year.Value.String()))
	}
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The request failed with error %s\n", err)
	} else {
		fmt.Println(api.RecordResponse(response))
	}
}
