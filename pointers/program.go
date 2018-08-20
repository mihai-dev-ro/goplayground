package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 2
	swap(&x, &y)
}

func swap(x *int, y *int) {
	fmt.Printf("input -> x: %v, y: %v\n", *x, *y)

	tmp := y
	y = x
	x = tmp

	fmt.Printf("output -> x: %v, y: %v\n", *x, *y)
}