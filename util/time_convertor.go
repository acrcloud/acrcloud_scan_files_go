package util

import (
	"math"
	"strconv"
)

func plural(count int, singular string) (result string) {
	result = strconv.Itoa(count) + singular
	if count < 10 {
		result = "0" + result
	}
	return
}

func secondsToHuman(input int) (result string) {
	seconds := input % (60 * 60 * 24)
	hours := math.Floor(float64(seconds) / 60 / 60)
	seconds = input % (60 * 60)
	minutes := math.Floor(float64(seconds) / 60)
	seconds = input % 60
	result = plural(int(hours), ":") + plural(int(minutes), ":") + plural(int(seconds), "")
	return
}
