package test

import (
	"al-go-rithm/solutions"
	"testing"
)

func TestSum(t *testing.T) {
	r1 := solutions.Sum([]int{1, 2, 3, 4, 5}) // 15
	r2 := solutions.Sum([]int{3, 3, 3, 3, 3}) // 15
	r3 := solutions.Sum([]int{3, 3, 2, 5, 2}) // 15

	if r1 != 15 || r2 != 15 || r3 != 15 {
		t.Errorf("Expected 15, but different")
	}
}
