package main

import (
	"os"
	"fmt"
)

func main() {
	n, err := fmt.Println("HelloWorld")

	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Println("0 bytes output")
	case n != 11:  /*including \n*/
		fmt.Println("wrong number of characters")
	default:
		fmt.Println("ok;")
	}
	fmt.Printf("\n")
}