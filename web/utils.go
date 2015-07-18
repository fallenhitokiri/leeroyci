package web

import (
	"strconv"
)

func stringToInt(number string) int {
	conv, err := strconv.Atoi(number)

	if err != nil {
		panic(err)
	}

	return conv
}
