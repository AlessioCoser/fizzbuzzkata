package game

import "strconv"

type specialNumber struct{
	num int
	word string
}

var specialNumbers [2]specialNumber = [2]specialNumber{
	{5, "Buzz"},
	{3, "Fizz"},
}

func Say(evaluating int) string {

	for _, special := range specialNumbers {
		if evaluating % special.num == 0 {
			return special.word
		}
	}

	return strconv.Itoa(evaluating)
}
