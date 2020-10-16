package api

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>
*/

import (
	"testing"
)

var urlTests = []struct {
	n        Request // input
	expected string  // expected result
}{
	{Request{Year: "2009"}, "https://www.balldontlie.io/api/v1/games/?seasons[]=2009,2008&postseason=False&team_ids[]=23&per_page=100"},
	{Request{Year: "2020"}, "https://www.balldontlie.io/api/v1/games/?seasons[]=2020,2019&postseason=False&team_ids[]=23&per_page=100"},
}

func TestBuildURL(t *testing.T) {
	for _, tt := range urlTests {
		actual := urlBuilder(tt.n)
		if actual != tt.expected {
			t.Errorf("BuildURL(%s): expected %s, actual %s", tt.n, tt.expected, actual)
		}
	}
}
