package api

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/grahamplata/sixers/config"
)

var urlTests = []struct {
	n        string // input
	m        string // input
	expected string // expected result
}{
	{"a", "b", "https://www.balldontlie.io/api/v1/games/?seasons[]=a,b&postseason=False&team_ids[]=23&per_page=100"},
	{"1", "2", "https://www.balldontlie.io/api/v1/games/?seasons[]=1,2&postseason=False&team_ids[]=23&per_page=100"},
	{"21", "2009", "https://www.balldontlie.io/api/v1/games/?seasons[]=21,2009&postseason=False&team_ids[]=23&per_page=100"},
	{"2019", "2020", "https://www.balldontlie.io/api/v1/games/?seasons[]=2019,2020&postseason=False&team_ids[]=23&per_page=100"},
}

func TestBuildURL(t *testing.T) {
	for _, tt := range urlTests {
		actual := BuildURL(tt.n, tt.m)
		if actual != tt.expected {
			t.Errorf("BuildURL(%s,%s): expected %s, actual %s", tt.n, tt.m, tt.expected, actual)
		}
	}
}

var dummy = Response{
	Data: []Game{
		{
			ID:               1,
			Date:             "2018-10-16T00:00:00.000Z",
			HomeTeamScore:    123,
			VisitorTeamScore: 122,
			Season:           2020,
			Period:           1,
			Status:           "Final",
			Time:             " ",
			PostSeason:       false,
			HomeTeam: Team{
				ID:           1,
				Abbreviation: "BOS",
				City:         "Boston",
				Conference:   "East",
				Division:     "Atlantic",
				FullName:     "Boston Celtics",
				Name:         "Celtics",
			},
			VisitorTeam: Team{
				ID:           23,
				Abbreviation: "PHI",
				City:         "Philadelphia",
				Conference:   "East",
				Division:     "Atlantic",
				FullName:     "Philadelphia 76ers",
				Name:         "76ers",
			},
		},
	},
}

func TestNextResponse(t *testing.T) {
	body, _ := json.Marshal(dummy)
	q := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}
	actual := NextResponse(q)
	if actual != "Sorry, there are not any available games." {
		t.Errorf("NextResponse(stub): expected false, actual %s", actual)
	}
}

func TestNextResponseFinal(t *testing.T) {
	dummy.Data[0].Status = "8PM"
	body, _ := json.Marshal(dummy)
	q := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}
	actual := NextResponse(q)
	gameTime := strings.TrimRight(dummy.Data[0].Time, " ")
	if actual != fmt.Sprintf("10,9 8 %s! There is a game currently @ %s %+s\n", config.SixersLogo, dummy.Data[0].Status, gameTime) {
		t.Errorf("NextResponse(stub): expected false, actual %s", actual)
	}
}

func TestRecordResponse(t *testing.T) {
	body, _ := json.Marshal(dummy)
	q := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}
	actual := RecordResponse(q)
	if actual == "1" {
		t.Errorf("NextResponse(stub): expected 1, actual %s", actual)
	}
}
