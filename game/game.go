package game

import "strconv"

func Say(evaluatingNumber int) string {
	if evaluatingNumber % 3 == 0 {
		return "Fizz"
	}

	return strconv.Itoa(evaluatingNumber)
}
