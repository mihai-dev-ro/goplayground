package math

import (
	// "fmt"
)

// Implementation of Selection Sort algorithm
// Selection Sort algorithm iterates through the list, finds the smallest element and places it
// at the beginning of the list and then continues iterativeley for the remaining elements
// until there are no remaining elements to be scanned
//
// Complexity analysis:
//    - time complexity: O(n * n)
//    - space complexity: O(1)
func SortSelection(list []int) []int {
	// fmt.Printf("Input list: %v\n", list)

	// empty or 1-item list; just return it as there is nothing to sort
	if len(list) < 2 {
		return list
	}

	left := 0
	for left < len(list) {

		min := list[left]
		iMin := left
		for i := left + 1; i < len(list); i++ {
			if min > list[i] {
				iMin = i
				min = list[i]
			}
		}
		list[left], list[iMin] = list[iMin], list[left]
		left++
	}

	// fmt.Printf("Selection sorted list: %v\n", list)

	return list; 
}