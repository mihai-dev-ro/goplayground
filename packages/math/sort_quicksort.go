package math

//"fmt"

// SortQuicksort implements Quicksort algorithm which is a Divide & Conquer algorithm that uses the pivoting technique to split the list in
// sub-lists and then proceed with sorting the sub-lists using the same method for identifying the pivot
// until they are sorted.
// After choosing (first, middle, last, etc - it doesn't matter) what pivot is used,
// move all the elements smaller than the pivot to the left and all the elements greater than the pivot to the right of it
//
// Complexity analysis
//    - time complexity: O(n * log(n))
//    - space complexity: O(n)
func SortQuicksort(list []int) []int {
	// fmt.Printf("Input list: %v\n", list)

	// empty or 1-item list; just return it as there is nothing to sort
	if len(list) < 2 {
		return list
	}

	// use the median element as the pivot (that reduces risks if list is already sorted in the reverse order)
	pivot := len(list) / 2 //
	left, right := 0, len(list)-1

	// put the pivot at the end of the list
	list[pivot], list[right] = list[right], list[pivot]

	// iterate in the list and put all the smaller elements at the left
	// and the larger element at the right
	for i := range list {
		if list[i] < list[right] {
			list[i], list[left] = list[left], list[i] // always put the smaller element to the right of the left position
			left++                                    // left marks the location of the first element greater than the pivot value
		}
	}

	// swap the pivot value into position
	list[left], list[right] = list[right], list[left]

	if left > 0 {
		SortQuicksort(list[:left]) // sort the sublist to the left of the pivot
	}
	if left < len(list)-1 {
		SortQuicksort(list[left+1:]) // sort the sublist to the right  of the pivot
	}

	// fmt.Printf("Quicksort sorted list: %v\n", list)
	return list
}
