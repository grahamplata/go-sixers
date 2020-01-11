/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	"github.com/grahamplata/sixers/schema"
)

func buildURL(val1 string, val2 string) string {
	url := fmt.Sprintf("%s/?seasons[]=%s,%s&postseason=False&team_ids[]=23&per_page=100", baseAPIURL, val1, val2)
	return url
}

func handleNextResponse(response *http.Response) string {
	spin := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	gameFound := false
	t := time.Now().Format("2006-01-02")

	spin.Start()
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObject schema.Response
	json.Unmarshal(responseData, &responseObject)
	for i := 0; i < len(responseObject.Data); i++ {
		cleanTime := fmt.Sprintf("%sT00:00:00.000Z", t)
		if responseObject.Data[i].Date == cleanTime {
			gameFound = true
		}
	}
	spin.Stop()

	if gameFound == true {
		return fmt.Sprintf("10,9 8 76ers! Game tonight! %s", t)
	}
	return fmt.Sprintf("Sorry, no game tonight.")
}

// handleRecordResponse
func handleRecordResponse(response *http.Response) string {
	spin := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	spin.Start()
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObject schema.Response
	json.Unmarshal(responseData, &responseObject)
	var gameCount int = 1
	var winRecord int = 0
	for i := 0; i < len(responseObject.Data); i++ {
		if responseObject.Data[i].VisitorTeamScore != 0 || responseObject.Data[i].HomeTeamScore != 0 {
			var visitorScore int = responseObject.Data[i].VisitorTeamScore
			var homeScore int = responseObject.Data[i].HomeTeamScore
			var homeID int = responseObject.Data[i].HomeTeam.ID
			if homeID == 23 {
				if homeScore > visitorScore {
					winRecord++
				}
			} else {
				if homeScore < visitorScore {
					winRecord++
				}
			}
			gameCount++
		}
	}
	spin.Stop()
	return fmt.Sprintf("Wins: %d Losses: %d\n", winRecord, (gameCount - winRecord))
}
