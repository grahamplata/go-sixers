package cmd

import "testing"

func TestAPIBaseURL(t *testing.T) {
	if baseAPIURL != "https://www.balldontlie.io/api/v1/games" {
		t.Error("Base Api URL not set to the expected URL.")
	}
	if baseAPIURL == "" {
		t.Error("Base Api URL cannot be empty.")
	}
}
