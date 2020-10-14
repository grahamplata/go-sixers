package api

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bndr/gotabulate"
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
func NextResponse(response *http.Response) string {
	var responseObject Response

	responseData, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseData, &responseObject)

	for _, v := range responseObject.Data {
		if v.Status != "Final" {
			gameTime := strings.TrimRight(v.Time, " ")
			resp := fmt.Sprintf("10,9 8 %s! There is a game currently @ %s %+s\n", config.SixersLogo, v.Status, gameTime)
			return resp
		}
	}
	return "Sorry, there are not any available games."
}

// RecordResponse ...
func RecordResponse(response *http.Response) string {
	var gameCount, winRecord int
	var responseObject Response

	responseData, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseData, &responseObject)

	for _, v := range responseObject.Data {
		if v.VisitorTeamScore != 0 || v.HomeTeamScore != 0 {
			if v.HomeTeam.ID == config.TeamID {
				if v.HomeTeamScore > v.VisitorTeamScore {
					winRecord++
				}
			} else {
				if v.HomeTeamScore < v.VisitorTeamScore {
					winRecord++
				}
			}
			gameCount++
		}
	}
	wins := fmt.Sprintf("%s %d", aurora.Green("Wins:"), winRecord)
	losses := fmt.Sprintf("%s %d", aurora.Red("Losses:"), (gameCount - winRecord))
	pct := fmt.Sprintf("%s %.3f", aurora.Yellow("Win pct:"), (float64(winRecord) / float64(gameCount)))
	return fmt.Sprintf("%s %s %s", wins, losses, pct)
}

// ScheduleResponse ...
func ScheduleResponse(response *http.Response) string {
	var headers = []string{"Game", "Date", "Home", "Away", "Winner"}
	var responseObject Response
	var table [][]string

	responseData, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseData, &responseObject)

	for i, v := range responseObject.Data {
		var winner string
		parsedTime, _ := time.Parse(time.RFC3339, v.Date)
		formattedTime := parsedTime.Format(config.LayoutUS)
		gameNum := strconv.Itoa(i + 1)
		home := fmt.Sprintf("%s: %v", v.HomeTeam.Abbreviation, v.HomeTeamScore)
		away := fmt.Sprintf("%s: %v", v.VisitorTeam.Abbreviation, v.VisitorTeamScore)

		if v.VisitorTeamScore != 0 || v.HomeTeamScore != 0 {
			if v.HomeTeam.ID == config.TeamID {
				if v.HomeTeamScore > v.VisitorTeamScore {
					winner = v.HomeTeam.FullName
				}
			} else {
				if v.HomeTeamScore < v.VisitorTeamScore {
					winner = v.HomeTeam.FullName
				}
			}
			winner = v.VisitorTeam.FullName
		}

		game := []string{gameNum, formattedTime, home, away, winner}
		table = append(table, game)
	}
	tabulate := gotabulate.Create(table)
	tabulate.SetHeaders(headers)
	return tabulate.Render("simple")
}
