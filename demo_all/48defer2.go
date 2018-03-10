package main

import "fmt"

func main() {
	a := 1
	a = 2  
	defer fmt.Printf("defer a: %v,%p\n", a, &a) //defer a: 2,0xc420012090
	a = 3
	fmt.Printf("outer a: %v,%p\n", a, &a) //outer a: 3,0xc420012090
}