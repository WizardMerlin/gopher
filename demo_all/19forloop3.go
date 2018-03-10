package main

import "fmt"

func main() {
	var sum int
	for i :=1 ; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum) //5050
}