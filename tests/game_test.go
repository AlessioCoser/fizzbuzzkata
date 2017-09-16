package tests

import (
	"testing"
	"github.com/AlessioCoser/fizzbuzzkata/game"
)

func TestReturnNumberAsString(t *testing.T) {
	equal("1", game.Say(1), t)
}

func equal(expected interface{}, actual interface{}, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected Number `%d` is not equals to `%d`", expected, actual)
	}
}
