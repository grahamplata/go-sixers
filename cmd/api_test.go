package cmd

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>
*/

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/grahamplata/sixers/schema"
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
		actual := buildURL(tt.n, tt.m)
		if actual != tt.expected {
			t.Errorf("buildURL(%s,%s): expected %s, actual %s", tt.n, tt.m, tt.expected, actual)
		}
	}
}

func TestHandleNextResponse(t *testing.T) {
	dummy := schema.Response{
		Data: []schema.Game{
			schema.Game{
				ID:               1,
				Date:             "2018-10-16T00:00:00.000Z",
				HomeTeamScore:    123,
				VisitorTeamScore: 122,
				Season:           2020,
				Period:           1,
				Status:           "Final",
				Time:             " ",
				PostSeason:       false,
				HomeTeam: schema.Team{
					ID:           1,
					Abbreviation: "BOS",
					City:         "Boston",
					Conference:   "East",
					Division:     "Atlantic",
					FullName:     "Boston Celtics",
					Name:         "Celtics",
				},
				VisitorTeam: schema.Team{
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
	actual := handleNextResponse(q)
	if actual != false {
		t.Errorf("handleNextResponse(stub): expected false, actual %t", actual)
	}
}

func TestHandleRecordResponse(t *testing.T) {
	dummy := schema.Response{
		Data: []schema.Game{
			schema.Game{
				ID:               1,
				Date:             "2018-10-16T00:00:00.000Z",
				HomeTeamScore:    122,
				VisitorTeamScore: 123,
				Season:           2020,
				Period:           1,
				Status:           "Final",
				Time:             " ",
				PostSeason:       false,
				HomeTeam: schema.Team{
					ID:           1,
					Abbreviation: "BOS",
					City:         "Boston",
					Conference:   "East",
					Division:     "Atlantic",
					FullName:     "Boston Celtics",
					Name:         "Celtics",
				},
				VisitorTeam: schema.Team{
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
	actual := handleRecordResponse(q)
	if actual == "1" {
		t.Errorf("handleNextResponse(stub): expected 1, actual %s", actual)
	}
}
