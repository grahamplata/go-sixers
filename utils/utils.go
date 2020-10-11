package utils

/*
Copyright © 2019 Graham Plata <graham.plata@gmail.com>
*/

import (
	"fmt"
	"strconv"
)

// IncrementString ...
func IncrementString(str string) string {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Unable to Increment string", err)
		return ""
	}
	i = i + 1
	str = strconv.FormatInt(int64(i), 10)
	return str
}

// DecrementString ...
func DecrementString(str string) string {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Unable to Decrement string", err)
		return ""
	}
	i = i - 1
	str = strconv.FormatInt(int64(i), 10)
	return str
}
