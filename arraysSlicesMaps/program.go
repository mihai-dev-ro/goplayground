package main

import (
	"fmt"
)

func main() {
	findSmallestNumber()	 
}

func slicesExercises() {

	/* CONCLUSIONS: 
	   1) slices x & y points to the same elements in the associated arrays behind x
       2) append function will create a new memory space (associated array) for the created slice 
	   3) copy function will replace the elements starting with the 0 index in the destination slice with the elements in the source slice
    */

	x := make([]float64, 5, 6)
	for i := 0; i < 5; i++ { 
		x[i] = float64(i+1) 
	}
	printSlice("Init slice X", x, x)  // [1 2 3 4 5] [1 2 3 4 5]

	y := x[1:3]
	printSlice("Derivate slice y := x[1:3]", y, x)  // [2 3] [1 2 3 4 5]

	copy(y, make([]float64, 2))
	printSlice("Slice y has overitten elements copy(y, make([]float64, 2))", y, x)  // [0 0] [1 0 0 4 5]	

	z := append(x, 6, 7, 8, 9, 10) // creates a new slice and returns it
	printSlice("Appended slice z := append(x, 6, 7, 8, 9, 10)", z, x)  // [1 0 0 4 5 6 7 8 9 10]

	// copies starting with the index 0
	copy(z, make([]float64, 2)) // returns the number of elements copied; 1st argument is the destination, 2nd elem is the source
	printSlice("Copied slice z <- copy(z, make([]float64, 2)", z, x)	 // [0 0 0 4 5 6 7 8 9 10] [1 0 0 4 5]

	// copies only a small segment of a slice
	t := make([]float64, 5)
	copy(t, z)
	printSlice("Copied slice t <- copy(make([]float64, 5), z)", t, z) // [0 0 0 4 5] [0 0 0 4 5 6 7 8 9 10]
}

func printSlice(text string, slice []float64, initialSlice []float64) {
	fmt.Print("/// ", text, "\n")
	fmt.Println("=== START === ")
	fmt.Print(slice, "   ", initialSlice, "\n")
	fmt.Println("=== END === ")
}

func findSmallestNumber() {
	x := []int{
	 48,96,86,68,
	 57,82,63,70,
	 37,34,83,27,
	 19,97, 9,17,
	 48,96,86,68,
	 57,82,63,70,
	 37,34,83,27,
	 19,97, 9,17,
	 48,96,86,68,
	 57,82,63,70,
	 37,34,83,27,
	 19,97, 9,17,
	 48,96,86,68,
	 57,82,63,70,
	 37,34,83,27,
	 19,97, 9,17,
	}

	/* possible solutions:
		1. iterate through the array and retain the smallest number -> has a time complexity of O(n)
		2. divide & conquer: split the array until you get to the leaves, calculate the minimun and then merge (merge sort)
	*/	

	fmt.Printf("Array used as input: %v\n length: %v\n", x, len(x))

	orderedArray := mergeSort(x)
	fmt.Printf("Smallest number is %v\n", orderedArray[0])
}

var indexFnCall int
var indexSortCall int

func min(a []int) (int) {
	indexFnCall++	
	fmt.Printf("Function is called: %v\n", indexFnCall)
	fmt.Printf("input: %v\n", a)

	if len(a) == 1 {
		indexSortCall++
		return a[0]
	} else if len(a) == 2 {
		indexSortCall++
		if a[0] < a[1] {
			return a[0]
		} else {
			return a[1]
		}
	} else {
		// get the split position (= median)
		splitPosition := int(len(a) / 2)
		partition1 := a[:splitPosition]
		partition2 := a[splitPosition:]

		minPartition1 := min(partition1)
		minPartition2 := min(partition2)

		if minPartition1 < minPartition2 {
			return minPartition1
		} else {
			return minPartition2
		}
	}
}

var indexMergeSortFnCall int
func mergeSort(a []int) []int {
	indexMergeSortFnCall++	
	fmt.Printf("Function <<Merge Sort>> is called: %v\n", indexMergeSortFnCall)
	fmt.Printf("input: %v\n", a)

	if len(a) == 1 {
		return a
	} else {
		// get the split position (= median)
		splitPosition := int(len(a) / 2)
		partition1 := a[:splitPosition]
		partition2 := a[splitPosition:]

		return mergePartitions(mergeSort(partition1), mergeSort(partition2))
	}
}

var indexMergePartitionsFnCall int
func mergePartitions(a []int, b []int) []int {
	panic("not implemented")

	indexMergePartitionsFnCall++	
	fmt.Printf("Function <<Merge Partitions>> is called: %v\n", indexMergePartitionsFnCall)
	fmt.Printf("input-> array 1: %v, array 2: %v\n", a, b)

	result := make([]int, len(a) + len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[i] {
			result[j+j] = a[i]
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

	fmt.Printf("sorted array: %v\n", result)
	return result
}

