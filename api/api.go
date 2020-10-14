package api

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/briandowns/spinner"
	"github.com/grahamplata/sixers/config"
	"github.com/logrusorgru/aurora"
)

// BuildURL takes the query parameters and returns a string to be used to retrieve data from the api
func BuildURL(seasonYearStart string, seasonYearEnd string) string {
	// TODO currently paging is hard coded address this later
	url := fmt.Sprintf("%s/?seasons[]=%s,%s&postseason=False&team_ids[]=%v&per_page=100", config.APIURL, seasonYearStart, seasonYearEnd, config.TeamID)
	return url
}

// NextResponse ...
func NextResponse(response *http.Response) {
	var responseObject Response

	responseData, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Data); i++ {
		// dirty way to handle the way the api flattens the time
		if responseObject.Data[i].Status != "Final" {
			status := responseObject.Data[i].Status
			gameTime := strings.TrimRight(responseObject.Data[i].Time, " ")
			fmt.Printf("10,9 8 %s! There is a game currently @ %s %+s\n", config.SixersLogo, status, gameTime)
			break
		}
	}
	fmt.Println("Sorry, there are not any available games.")
	return

}

// RecordResponse ...
func RecordResponse(response *http.Response) string {
	spin := spinner.New(config.SpinnerType, config.SpinnerDuration)
	spin.Start()
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObject Response
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
	wins := fmt.Sprintf("%s %d", aurora.Green("Wins:"), winRecord)
	losses := fmt.Sprintf("%s %d", aurora.Red("Losses:"), (gameCount - winRecord))
	pct := fmt.Sprintf("%s %.3f", aurora.Yellow("Win pct:"), (float64(winRecord) / float64(gameCount)))
	return fmt.Sprintf("%s\n%s\n%s", wins, losses, pct)
}
