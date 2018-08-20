// Package math package that exposes the mathematics used
package math

// "fmt"

// CreateFibonacciSequence is a generator of a fibonacci sequence that uses the
// memoization technique to retain the calculated values in the sequence
func CreateFibonacciSequence() func(uint) uint {
	fibonacciValues := make(map[uint]uint)

	var fibonacciSequence func(uint) uint

	fibonacciSequence = func(n uint) uint {
		if _, found := fibonacciValues[n]; !found {
			var result uint
			if n < 2 {
				result = n
			} else {
				result = fibonacciSequence(n-2) + fibonacciSequence(n-1)
			}
			fibonacciValues[n] = result
			// fmt.Printf("Fibonacci(%v): %v\n", n, result)
		}

		var val, _ = fibonacciValues[n]
		return val
	}

	return fibonacciSequence
}
