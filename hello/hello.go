package main

import (
	"fmt"
	"os/user"
)

func main() {
	usr, err := user.Current()

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("Hello, my name is %v\n", usr.Name)

	a := 1.01
	b := 0.99

	fmt.Printf("%v - %v = %v\n", a, b, a-b) 

	var x uint32 = 321325
	var y uint32 = 424521

	fmt.Printf("%v * %v = %v\n", x, y, x * y) 
}