package test

import (
	"al-go-rithm/solutions"
	"testing"
)

func TestNumInList(t *testing.T) {
	r1 := solutions.NumInList([]int{1, 2, 3, 4, 5}, 5) // true
	r2 := solutions.NumInList([]int{3, 3, 3, 3, 3}, 5) // false
	r3 := solutions.NumInList([]int{3, 3, 5, 3, 3}, 5) // true
	r4 := solutions.NumInList(nil, 5)                  // false
	r5 := solutions.NumInList([]int{}, 5)              // false

	if r1 != true || r3 != true {
		t.Errorf("Expected true, but false!")
	}

	if r2 != false || r4 != false || r5 != false {
		t.Errorf("Expected false, but true!")
	}
}
