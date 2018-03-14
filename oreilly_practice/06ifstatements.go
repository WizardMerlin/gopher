package main

import (
	"os"
	"fmt"
)

func main() {

	if numbers, errors := fmt.Println("Hello"); errors != nil {
		os.Exit(1)
	} else {
		fmt.Printf("%d bytes charactors printed.\n", numbers)
	}
}