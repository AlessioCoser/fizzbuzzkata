package tests

import (
	"testing"
	"github.com/AlessioCoser/fizzbuzzkata/game"
)

func TestReturnNumberAsString(t *testing.T) {

	if game.Say(1) != "1" {
		t.Errorf("not implemented yet")
	}
}
