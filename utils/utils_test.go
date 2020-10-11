package utils

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>
*/

import "testing"

var incrementTests = []struct {
	n        string // input
	expected string // expected result
}{
	{"2000", "2001"},
	{"2011", "2012"},
	{"2019", "2020"},
	{"2020", "2021"},
}

func TestIncrementString(t *testing.T) {
	for _, tt := range incrementTests {
		actual := IncrementString(tt.n)
		if actual != tt.expected {
			t.Errorf("incrementString(%s): expected %s, actual %s", tt.n, tt.expected, actual)
		}
	}
}

var decrementTests = []struct {
	n        string // input
	expected string // expected result
}{
	{"2000", "1999"},
	{"2011", "2010"},
	{"2019", "2018"},
	{"2020", "2019"},
}

func TestDecrementString(t *testing.T) {
	for _, tt := range decrementTests {
		actual := DecrementString(tt.n)
		if actual != tt.expected {
			t.Errorf("decrementString(%s): expected %s, actual %s", tt.n, tt.expected, actual)
		}
	}
}
