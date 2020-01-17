package cmd

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/

import (
	"encoding/json"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/grahamplata/sixers/schema"
	"github.com/logrusorgru/aurora"
	"io/ioutil"
	"net/http"
	"time"
)

func buildURL(val1 string, val2 string) string {
	url := fmt.Sprintf("%s/?seasons[]=%s,%s&postseason=False&team_ids[]=23&per_page=100", baseAPIURL, val1, val2)
	return url
}

func handleNextResponse(response *http.Response) bool {
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
			sixers := fmt.Sprintf("%d%ders", aurora.Bold(aurora.Red(7)), aurora.Bold(aurora.Blue(6)))
			fmt.Printf("10,9 8 %s!\nYou're in luck! There is a game today @ %s %s!\n", sixers, responseObject.Data[i].Status, responseObject.Data[i].Time)
			gameFound = true
		}
	}
	spin.Stop()

	if gameFound == true {
		return true
	}
	return false
}

// handleRecordResponse
func handleRecordResponse(response *http.Response) string {
	spin := spinner.New(spinner.CharSets[21], 100*time.Millisecond)
	spin.Start()
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObject schema.Response
	json.Unmarshal(responseData, &responseObject)
	var gameCount int
	var winRecord int
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
	pct := (float64(winRecord) / float64(gameCount))
	return fmt.Sprintf("Wins: %d Losses: %d Winrate: %.3f\n", winRecord, (gameCount - winRecord), pct)
}
