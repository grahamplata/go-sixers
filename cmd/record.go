/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// A Response Struct to map response to.
type Response struct {
	Data []Game `json:"data"`
}

// A Game Struct to map games to.
type Game struct {
	ID               int    `json:"id"`
	Date             string `json:"date"`
	HomeTeamScore    int    `json:"home_team_score"`
	VisitorTeamScore int    `json:"visitor_team_score"`
	Season           int    `json:"season"`
	Period           int    `json:"period"`
	Status           string `json:"status"`
	Time             string `json:"time"`
	PostSeason       bool   `json:"postseason"`
	HomeTeam         Team   `json:"home_team"`
	VisitorTeam      Team   `json:"visitor_team"`
}

// A Team Struct to map response to.
type Team struct {
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
	City         string `json:"ƒcity"`
	Conference   string `json:"conference"`
	Division     string `json:"division"`
	FullName     string `json:"full_name"`
	Name         string `json:"name"`
}

// A Meta Struct to map response to.
type Meta struct {
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	PerPage     int `json:"per_page"`
	TotalCount  int `json:"total_count"`
}

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Get the 76ers record for the current season.",
	Long:  "Get the 76ers record for the current season.",
	Run: func(cmd *cobra.Command, args []string) {
		year := cmd.Flag("year")
		fmt.Printf("Fetching 76ers record for %s \n", year.Value.String())
		response, err := http.Get(`https://www.balldontlie.io/api/v1/games/?seasons[]=2019,2020&postseason=False&team_ids[]=23&per_page=100`)
		if err != nil {
			fmt.Printf("The request failed with error %s\n", err)
		} else {
			responseData, _ := ioutil.ReadAll(response.Body)
			var responseObject Response
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
			fmt.Printf("Wins: %d Losses: %d\n", winRecord, (gameCount - winRecord))
		}
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)

	// Here you will define your flags and configuration settings.
	recordCmd.Flags().StringP("year", "y", "", "The year to query")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
