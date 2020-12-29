/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>
*/

package config

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

// API Configuration

// APIURL is the base endpoint string
const APIURL string = "https://www.balldontlie.io/api/v1/games"

// Team Configuration

// TeamID is the sixers team id from the api
const TeamID int = 23

// SixersLogo is a colored string representing the sixers logo
var SixersLogo string = fmt.Sprintf("%d%ders", aurora.Bold(aurora.Red(7)), aurora.Bold(aurora.Blue(6)))

// Date Time Configuration

// YearFormat year format
const YearFormat string = "2006"

// LayoutISO year format
const LayoutISO string = "2006-01-02"

// LayoutUS year format
const LayoutUS string = "January 2, 2006"
