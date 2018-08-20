package main

import (
	"fmt"
)

func main() {
	fmt.Printf("SUM:\n")
	var _ = sum([]int {3, 67, 100, 7});

	fmt.Printf("HALF:\n")
	var _,_ = half(3)
	var _,_ = half(4)

	fmt.Printf("GREATEST:\n")
	var _ = greatest(56, 7, 100, -10)

	fmt.Printf("ODD NUMEBERS GENERATOR:\n")
	oddGenerator := makeOddGenerator()
	var _ = oddGenerator()
	var _ = oddGenerator()
	var _ = oddGenerator()	

	fmt.Printf("FIBONACCI MEMOIZATION:\n")
	fibonacci := makeFibonacci()
	var _ = fibonacci(100)

	fmt.Printf("FIBONACCI RECURSIVE:\n")
	var _ = fibonacciRecursive(10)
}

func sum(numbers []int) int {
	fmt.Printf("array: %v\n", numbers)

	total := 0
	for _,v := range numbers {
		total += v
	}

	fmt.Printf("sum: %v\n", total)
	return total
}

func half(number int) (result int, ok bool) {
	fmt.Printf("input: %v\n", number)

	result = int(number / 2)
	ok = number - (result * 2) == 0

	fmt.Printf("result: (%v, %v)\n", result, ok)
	return
}

func greatest(args ...int) int {
	fmt.Printf("input: %v\n", args)

	MaxUint := ^uint(0)
	MaxInt := int(MaxUint >> 1)
	MinInt := -MaxInt - 1
	fmt.Printf("(maxUint: %v, maxInt: %v, minInt: %v)\n", MaxUint, MaxInt, MinInt)

	max := MinInt
	for _,v := range args {
		if v > max {
			max = v
		}
	}

	fmt.Printf("result: %v\n", max)
	return max
}

func makeOddGenerator() func() uint {
	sequence := uint(1)

	return func() (result uint) {
		result = sequence
		sequence += 2

		fmt.Printf("result: %v\n", result)
		return
	}
}

func makeFibonacci() func(uint) uint {
	memoizationValues := make(map[uint]uint)

	var fibonacci func(uint) uint

	fibonacci = func (num uint) (result uint) {
		fmt.Printf("input: %v\n", num)

		if num < 2 {
			memoizationValues[num] = num
			result,_ = memoizationValues[num]
		} else {
			if value,found := memoizationValues[num-1]; !found {
				value = fibonacci(num-1)
				memoizationValues[num-1] = value
			}
			if value,found := memoizationValues[num-2]; !found {
				value = fibonacci(num-2)
				memoizationValues[num-2] = value
			}

			memoizationValues[num] = memoizationValues[num-1] + memoizationValues[num-2]
			result = memoizationValues[num]
		}

		fmt.Printf("result: %v\n", result)
		return
	}

	return fibonacci 
}

func fibonacciRecursive(num uint) (result uint) {
	fmt.Printf("input: %v\n", num)

	if num < 2 {
		result = num
	} else {
		result = fibonacciRecursive(num - 1) + fibonacciRecursive(num - 2)
	}

	fmt.Printf("result: %v\n", result)
	return
}