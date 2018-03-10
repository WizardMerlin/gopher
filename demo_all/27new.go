package main

import "fmt"

func main() {
	p := new([3]int)
	fmt.Println(*p)

	p[2] = 4
	fmt.Println(*p)
}