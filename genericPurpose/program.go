package main

import (
	"fmt"
)

func main() {
	convertCelsius()
	convertFeet()
}

func convertCelsius() {
	var input float32
	fmt.Print("Input the value of Fahrenheit: ")
	fmt.Scanf("%f", &input)

	output := (input - 32.0) * 5 / 9

	fmt.Printf("Value in Celsius is [C = (F - 32) * 5 / 9] = %v\n", output)
}

func convertFeet() {
	var input float32
	fmt.Print("Input the value of Feet: ")
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Printf("Value in Meters is [M = F * 0.3048] = %v\n", output)
}
