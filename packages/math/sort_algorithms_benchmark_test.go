package math

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var listSize = 1 << 16
var MaxUInt = ^uint(0)
var MaxInt = int(MaxUInt >> 52)

var sortAlgorithms = map[string]interface{}{
	"SortBubble":           SortBubble,
	"SortMerge":            SortMerge,
	"SortQuicksort":        SortQuicksort,
	"SortSelection":        SortSelection,
	"SortInsertion":        SortInsertion,
	"SortHeap":             SortHeap,
	"SortHeapArray":        SortHeapArray,
	"SortHeapArrayNotMine": SortHeapArrayNotMine,
}

func TestSortAgorithmsEficiency(t *testing.T) {

	// init array to be sorted
	randomArray := make([]int, listSize)
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < len(randomArray); i++ {
		randomArray[i] = rand.Intn(MaxInt)
	}

	// create a list of identical inputs (last one is the sort from the standard go package)
	inputs := make([][]int, len(sortAlgorithms)+1)
	for i := range inputs {
		inputs[i] = make([]int, len(randomArray))
		copy(inputs[i], randomArray)
	}

	// store of the results
	results := make([][]int, len(sortAlgorithms)+1)
	i := 0

	// perform sorting
	for fnName, fnSort := range sortAlgorithms {
		start := time.Now()
		results[i] = fnSort.(func([]int) []int)(inputs[i])
		i++
		end := time.Now()
		fmt.Printf("Duration %v: %v\n", fnName, end.Sub(start))
	}

	// sorting from the sort distribution
	last := len(results) - 1
	results[last] = make([]int, len(inputs[last]))
	copy(results[last], inputs[last])
	start := time.Now()
	sort.Ints(results[last])
	end := time.Now()
	fmt.Printf("Duration implementation of Go standard sort: %v\n", end.Sub(start))

	// make sure all results are equal to the list sorted with the sorted available in the sort package
	for i := 0; i < len(results)-1; i++ {
		if !arrayEqual(results[last], results[i]) {
			t.Error(
				"Results from resultSet", i,
				"not equal to results from sorting with Go Sort algorithm", "\n",
				"Expected", results[i], "\n",
				"Equal to", results[last],
			)
		}
	}
}
