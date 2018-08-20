package math

// "fmt"

// SortInsertion implements the Insertion Sort algorithm which is creating a new list and progressively moving elements from the original list
// to the correct position in the new sorted list
//
// Complexity analysis:
//    - time complexity: O(n * n)
//    - space complexity: O(1)
func SortInsertion(list []int) []int {
	// fmt.Printf("Input list: %v\n", list)

	// empty or 1-item list; just return it as there is nothing to sort
	if len(list) < 2 {
		return list
	}

	result := make([]int, len(list))

	upper := 0
	for i := 0; i < len(list); i++ {
		j := 0
		for j < upper && list[i] > result[j] {
			j++
		}
		if j == upper {
			result[upper] = list[i]
		} else {
			for k := upper; k > j; k-- {
				result[k-1], result[k] = result[k], result[k-1]
			}
			result[j] = list[i]
		}
		upper++
	}

	// fmt.Printf("Insertion sorted list: %v\n", result)
	return result
}
