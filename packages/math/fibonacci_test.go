package math

import (
	"testing"
	"reflect"
)

func TestCreateFibonacciSequence(t *testing.T) {
	type TestPair struct {
		input, output uint
	}

	testPairs := []TestPair{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	fibonacciSeq := CreateFibonacciSequence()
	fibonacciSeqType := reflect.TypeOf(fibonacciSeq)
	if fibonacciSeqType.String() != "func(uint) uint" {
		t.Error(
			"Expected CreateFibonacciSequence() to return a function with signature",
			"func(uint) uint",

		)	
	}

	var result uint
	for _, pair := range testPairs {
		result = fibonacciSeq(pair.input)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"Expected", pair.output,
				"To be equal to", result,
			)
		}
	}	
}