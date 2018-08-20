package main

import (
	"fmt"
)

func main() {
	//printFor()
	//printIf()
	//printSwitch()

	printFizzBuzz()
}

func printFor() {
	// printeaza toate valorile posibile de stocat intr-un byte
	fmt.Printf("Print all unicode bytes from %v to %v\n\n", 0, 1 << 8)

	for i := 0; i < (1 << 8); i++ {
		fmt.Printf("value: %v is %q\n", i, i)
	}
	fmt.Println()
}

func printIf() {
	// printeaza numai numerele prime intr-un byte
	fmt.Printf("Print all primes numbers in the range from %v to %v\n\n", 0, 1 << 8)

	for i := 1; i < (1 << 8); i++ {
		if isPrimeNumber(i) {
			fmt.Printf("value: %v is prime\n", i)
		}
	}
	fmt.Println()
}

func isPrimeNumber(num int) (bool){
	if num < 4 {
		return true
	}
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func printSwitch() {
	/* cats of the wilds
		cheetah(= ghepard), 
		lion, 
		caracal(= serval), 
		lynx, 
		leopard, 
		cougar (puma), 
		tiger, 
		jaguar
	*/

	for i := 0; i < 8; i++ {
		switch i % 8 {
			case 0:
				fmt.Printf("Value %v is %s\n", i, "Cheetah");
			case 1:
				fmt.Printf("Value %v is %s\n", i, "Lion");
			case 2:
				fmt.Printf("Value %v is %s\n", i, "Caracal");
			case 3:
				fmt.Printf("Value %v is %s\n", i, "Lynx");
			case 4:
				fmt.Printf("Value %v is %s\n", i, "Leopard");
			case 5:
				fmt.Printf("Value %v is %s\n", i, "Cougar");
			case 6:
				fmt.Printf("Value %v is %s\n", i, "Tiger");
			case 7:
				fmt.Printf("Value %v is %s\n", i, "Jaguar");
			default: 
				fmt.Printf("Not found\n");

		}
	}
}

func printFizzBuzz() {
	for i :=1; i < 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	} 
}
