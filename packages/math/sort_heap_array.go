package math

// SortHeapArray implements the Heap Sort algorithm by using an array to hold the
// heap data structure, given that the heap tree is a complete binary tree and it
// requires that each level in the tree is complete before filling up nodes on a lower level
//
// Complexity analysis
//    - time complexity: O(n*log(n))
//    - space complexity: O(1)
func SortHeapArray(list []int) []int {
	// fmt.Printf("Input list: %v\n", list)

	// empty or 1-item list; just return it as there is nothing to sort
	if len(list) < 2 {
		return list
	}

	// heapify
	n := len(list)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(list, i, n)
	}

	last := len(list) - 1
	for last >= 0 {
		list[0], list[last] = list[last], list[0]

		// re-construct the heap data
		heapify(list, 0, last)
		last--
	}

	// fmt.Printf("Input list after heapify: %v\n", result)
	return list
}

// heapify the sub-tree with the root at iNode
func heapify(list []int, iNode int, n int) {
	largest := iNode
	left := 2*iNode + 1
	right := 2*iNode + 2

	// get the largest among the 3 nodes and put it as parent
	if left < n && list[largest] < list[left] {
		largest = left
	}
	if right < n && list[largest] < list[right] {
		largest = right
	}

	if largest != iNode {
		list[largest], list[iNode] = list[iNode], list[largest]

		heapify(list, largest, n)
	}

	return
}
