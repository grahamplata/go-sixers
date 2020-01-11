/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/
package cmd

import (
	"strconv"
)

const baseURL string = "https://www.balldontlie.io/api/v1/games"

func incrementString(str string) string {
	i, _ := strconv.Atoi(str)
	i = i + 1
	str = strconv.FormatInt(int64(i), 10)
	return str
}

func decrementString(str string) string {
	i, _ := strconv.Atoi(str)
	i = i - 1
	str = strconv.FormatInt(int64(i), 10)
	return str
}
