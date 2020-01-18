package cmd

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/

const baseAPIURL string = "https://www.balldontlie.io/api/v1/games"

var sixers string = fmt.Sprintf("%d%ders", aurora.Bold(aurora.Red(7)), aurora.Bold(aurora.Blue(6)))
