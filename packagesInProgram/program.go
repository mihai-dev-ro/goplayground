package main

import (
	"exercises/packages/math"
)

func main() {
	fibonacciSeq := math.CreateFibonacciSequence()
	fibonacciSeq(10)
}