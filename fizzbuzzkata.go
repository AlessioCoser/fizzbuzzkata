package main

import (
	"fmt"
	"github.com/AlessioCoser/fizzbuzzkata/game"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(game.Say(i))
	}
}
