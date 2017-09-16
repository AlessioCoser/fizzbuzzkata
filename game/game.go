package game

import "strconv"

func Say(evaluatingNumber int) string {
	if evaluatingNumber % 3 == 0 {
		return "Fizz"
	}

	if evaluatingNumber % 5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(evaluatingNumber)
}
