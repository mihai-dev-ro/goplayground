package math

//"fmt"

// SortBubble implements the sort algorithm which is resolving the sorting of a list of elements by progressively "bubbling up" the largest values
// to the end of the list. It implies putting adjacent values in the correct order and perform as many iterations as
// necessary until there is no swapping required. The iteration where there "no movement" signals the end of the algorithm
//
// Complexity analysis:
//    - time complexity : O(n*n)
//    - space complexity : O(1)
func SortBubble(list []int) []int {
	// fmt.Printf("Input list: %v\n", list)

	if len(list) < 2 {
		return list
	}

	swapPerformed := true // monitors if any swaping performed
	for swapPerformed {
		swapPerformed = false // reset flag

		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]

				swapPerformed = true // mark the swapping
			}
		}
	}

	// fmt.Printf("Bubble sorted list: %v\n", list)
	return list
}
