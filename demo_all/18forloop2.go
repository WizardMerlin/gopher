package main

import "fmt"

func main() {
	//带有单个条件的 for 循环

	a := 1
	for a > 0 {
		fmt.Println("循环中...");
		a--
	}
	fmt.Println("循环结束了...")
}