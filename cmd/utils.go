package cmd

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>

*/

import (
	"strconv"
)

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
