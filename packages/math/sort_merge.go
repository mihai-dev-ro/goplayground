package math

// "fmt"

// SortMerge implements Merge Sort algorithm which is a Divide & Conquer type of algorithm that recursively splits the list into sub-lists until
// it reaches an atomic level of 1-item list
// followed by a recursive merging of the smaller sub-lists back into larger lists
//
// Complexity analysis
//    - time complexity: O(n * log(n))
//    - space complexity: O(n)
func SortMerge(list []int) []int {
	// fmt.Printf("Input list: %v\n", list)

	// empty or 1-item list; just return it as there is nothing to sort
	if len(list) < 2 {
		return list
	}

	// split and merge the sub-lists
	splitIndex := len(list) / 2

	// cSublistLeft := make(chan []int)
	// go func() {
	// 	cSublistLeft <- SortMerge(list[:splitIndex])
	// }()

	// cSublistRight := make(chan []int)
	// go func() {
	// 	cSublistRight <- SortMerge(list[splitIndex:])
	// }()

	// result := merge(<- cSublistLeft, <- cSublistRight)

	// no concurrency
	result := merge(SortMerge(list[:splitIndex]), SortMerge(list[splitIndex:]))

	// CONCLUSION:
	// running several tests with very large lists
	// revealed that concurrency does not improve the speed

	// fmt.Printf("Merge sorted list: %v\n", result)

	return result
}

func merge(a, b []int) []int {
	result := make([]int, len(a)+len(b))

	var i, j = 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result[i+j] = a[i]
			i++
		} else {
			result[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		result[i+j] = a[i]
		i++
	}
	for j < len(b) {
		result[i+j] = b[j]
		j++
	}

	return result
}
