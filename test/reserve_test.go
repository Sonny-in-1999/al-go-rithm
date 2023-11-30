package test

import (
	"al-go-rithm/solutions"
	"testing"
)

func TestReserve(t *testing.T) {
	r1 := solutions.Reserve("cat")
	r2 := solutions.Reserve("alphabet")

	if r1 != "tac" {
		t.Errorf("Expected tac, but otherwise")
	}

	if r2 != "tebahpla" {
		t.Errorf("Expected tebahpla, but otherwise")
	}
}
