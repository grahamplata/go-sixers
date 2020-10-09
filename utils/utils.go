package utils

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>
*/

import (
	"strconv"
)

func IncrementString(str string) string {
	i, _ := strconv.Atoi(str)
	i = i + 1
	str = strconv.FormatInt(int64(i), 10)
	return str
}

func DecrementString(str string) string {
	i, _ := strconv.Atoi(str)
	i = i - 1
	str = strconv.FormatInt(int64(i), 10)
	return str
}
