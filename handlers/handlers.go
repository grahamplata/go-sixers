package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bndr/gotabulate"
	"github.com/grahamplata/sixers/api"
	"github.com/grahamplata/sixers/config"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

// Next handles when the next sixers game is available
func Next() {
	requestParams := api.Request{Year: time.Now().Format(config.YearFormat)}
	results := api.Fetch(requestParams)
	for _, v := range results.Data {
		if v.Status != "Final" {
			gameTime := strings.TrimRight(v.Time, " ")
			nextResponse := fmt.Sprintf("10,9 8 %s! There is a game currently @ %s %+s\n", config.SixersLogo, v.Status, gameTime)
			fmt.Println(nextResponse)
		}
	}
}

// Record handles the record statistics
func Record(cmd *cobra.Command) {
	var gameCount, winRecord int
	var requestParams api.Request

	yearFlag := cmd.Flag("year")

	if yearFlag.Value.String() == "" {
		requestParams.Year = time.Now().Format(config.YearFormat)
	} else {
		// TODO add format check
		requestParams.Year = yearFlag.Value.String()
	}

	results := api.Fetch(requestParams)

	for _, v := range results.Data {
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
	final := fmt.Sprintf("%s %s %s", wins, losses, pct)

	fmt.Println(final)
}

// Schedule handles when the next sixers game is available
func Schedule(cmd *cobra.Command) {
	var headers = []string{"Game", "Date", "Home", "Away", "Winner"}
	var table [][]string
	var requestParams api.Request

	yearFlag := cmd.Flag("year")

	if yearFlag.Value.String() == "" {
		requestParams.Year = time.Now().Format(config.YearFormat)
	} else {
		// TODO add format check
		requestParams.Year = yearFlag.Value.String()
	}

	results := api.Fetch(requestParams)

	for i, v := range results.Data {
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
	fmt.Println(tabulate.Render("simple"))
}
