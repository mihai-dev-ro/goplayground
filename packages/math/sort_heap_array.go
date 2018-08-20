package math

// SortHeapArray implements the Heap Sort algorithm by using an array to hold the
// heap data structure, given that the heap tree is a complete binary tree and it
// requires that each level in the tree is complete before filling up nodes on a lower level
func SortHeapArray(list []int) []int {
	// fmt.Printf("Input list: %v\n", list)

	// empty or 1-item list; just return it as there is nothing to sort
	if len(list) < 2 {
		return list
	}

	result := make([]int, len(list))

	tempList := list
	// heapify
	heapify(list, len(list)/2)
	index := len(result) - 1
	for len(tempList) > 0 {
		result[index] = tempList[0]
		index--

		//create a new slice by removing the first element
		tempList = tempList[1:]
		// heapify
		heapify(tempList, len(tempList)/2)
	}

	// fmt.Printf("Input list after heapify: %v\n", result)
	return result
}

func heapify(list []int, iMax int) {
	if iMax <= 1 {
		return
	}

	for i := 0; i < iMax; i++ {
		if 2*i+1 < len(list) && list[i] < list[2*i+1] {
			list[i], list[2*i+1] = list[2*i+1], list[i]
		}
		if 2*i+2 < len(list) && list[i] < list[2*i+2] {
			list[i], list[2*i+2] = list[2*i+2], list[i]
		}
	}

	heapify(list, iMax/2)
	return
}
