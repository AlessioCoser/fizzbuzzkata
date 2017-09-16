package tests

import (
	"testing"
	"github.com/AlessioCoser/fizzbuzzkata/game"
)

func TestReturnsNumbersAsString(t *testing.T) {
	equal("1", game.Say(1), t)
	equal("2", game.Say(2), t)
}

func equal(expected interface{}, actual interface{}, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected Number `%q` is not equals to `%q`", expected, actual)
	}
}
