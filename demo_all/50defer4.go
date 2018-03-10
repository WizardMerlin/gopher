package main

import "fmt"

func main() {
	a := 1
	a = 2

	b := 5
	b = 6

	defer func(b int) {
		fmt.Printf("defer b: %d, %p\n", b, &b) 
		fmt.Printf("defer a: %d, %p\n", a, &a) 
	}(b)

	a = 3
	fmt.Printf("outer a: %d, %p\n", a, &a) 
	b = 7
	fmt.Printf("outer b: %d, %p\n", b, &b)
}

/**
outer a: 3, 0xc420012090
outer b: 7, 0xc420012098
defer b: 6, 0xc4200120c0
defer a: 3, 0xc420012090
**/