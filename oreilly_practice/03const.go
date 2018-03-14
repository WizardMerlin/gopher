package main

import "fmt"

const (
	answer1 = iota*2
	answer2
)

func main() {
	fmt.Println(answer1, answer2)
}