package math

import (
	"fmt"
)

// SortHeapArrayNotMine is implementing the solution found here:
// https://www.geeksforgeeks.org/heap-sort/
func SortHeapArrayNotMine(list []int) []int {
	// empty or 1-item list; just return it as there is nothing to sort
	if len(list) < 2 {
		return list
	}

	n := len(list)
	for i := n/2 - 1; i > 0; i-- {
		heapifyNotMine(list, n, i)
	}

	fmt.Printf("Heapified list: %v\n", list)

	// one by one extract an elemengt from the heap
	for i := n - 1; i >= 0; i-- {
		// move the first which is the greatest integer to the last position
		list[0], list[i] = list[i], list[0]

		// call max heapify on the reduced heap
		heapifyNotMine(list, i, 0)
	}

	return list
}

func heapifyNotMine(list []int, n int, i int) {

	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && list[l] > list[largest] {
		largest = l
	}

	if r < n && list[r] > list[largest] {
		largest = r
	}

	if largest != i {
		list[largest], list[i] = list[i], list[largest] // swap values

		// recursively heapify the remaining sub-tree
		heapifyNotMine(list, n, largest)
	}
}
