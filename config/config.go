package config

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/logrusorgru/aurora"
)

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>
*/

// API Configuration
////////////////////////////////////////////////////////////

// APIURL is the base endpoint string
const APIURL string = "https://www.balldontlie.io/api/v1/games"

// Team Configuration
////////////////////////////////////////////////////////////

// TeamID is the sixers team id from the api
const TeamID int = 23

// SixersLogo is a colored string representing the sixers logo
var SixersLogo string = fmt.Sprintf("%d%ders", aurora.Bold(aurora.Red(7)), aurora.Bold(aurora.Blue(6)))

// Date Time Configuration
////////////////////////////////////////////////////////////

// YearFormat year format
const YearFormat = "2006"

// DateFormat year format
const DateFormat = "2006-01-02"

// Spinner Configuration
////////////////////////////////////////////////////////////
// using #21
// ⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈ spinner pattern
// from https://github.com/briandowns/spinner

// SpinnerType ...
var SpinnerType []string = spinner.CharSets[21]

// SpinnerDuration ...
var SpinnerDuration time.Duration = 100 * time.Millisecond
