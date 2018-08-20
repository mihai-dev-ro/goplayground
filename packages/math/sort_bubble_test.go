package math

import (
	"testing"
	"fmt"
)

func TestSortBubble(t *testing.T) {
	type TestPair struct {
		input, expected []int
	} 

	testPair := TestPair{
		[]int{ 78, 67, 8, 10, 5 },
		[]int{ 5, 8, 10, 67, 78 },
	}

	result := SortBubble(testPair.input)
	if !arrayEqual(result, testPair.expected) {
		t.Error(
			"For test pair:", testPair,
			"Expected:", testPair.expected,
			"To be equal to:", result,
		)
	}
}

func arrayEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		fmt.Println("length not equal", "a:", len(a), "b:", len(b))
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			fmt.Println("element not equal at position", i, "a[i]:", a[i], "b[i]", b[i])
			return false
		}
	}

	return true 
}