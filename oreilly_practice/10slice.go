package main

import (
	"fmt"
)

func changeStr(strs []string) {
	for _, value := range strs {
		fmt.Print(value)
	}
	fmt.Println()
	strs[0] = "ur"
}

func main() {
	words := []string{"the", "world","is","ours."}
	changeStr(words)
	changeStr(words)
}