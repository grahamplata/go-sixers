/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>
*/

package api

// A Request Struct to map request parameters to.
type Request struct {
	Year string
}

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
