package config

import "testing"

func TestAPIBaseURL(t *testing.T) {
	if BaseAPIURL != "https://www.balldontlie.io/api/v1/games" {
		t.Error("Base Api URL not set to the expected URL.")
	}
	if BaseAPIURL == "" {
		t.Error("Base Api URL cannot be empty.")
	}
}
