package game

import "strconv"

type specialNumber struct{
	num int
	word string
}

var specialNumbers [3]specialNumber = [3]specialNumber{
	{15, "FizzBuzz"},
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
