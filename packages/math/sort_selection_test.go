package math

import (
	"testing"
)

func TestSortSelection(t *testing.T) {
	type TestPair struct {
		input, expected []int
	} 

	testPair := TestPair{
		[]int{ 78, 67, 8, 10, 5 },
		[]int{ 5, 8, 10, 67, 78 },
	}

	result := SortSelection(testPair.input)
	if !arrayEqual(result, testPair.expected) {
		t.Error(
			"For test pair:", testPair,
			"Expected:", testPair.expected,
			"To be equal to:", result,
		)
	}
}