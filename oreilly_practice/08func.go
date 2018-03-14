package main

import (
	"fmt"
)

func main() {
	printer("mine")
}

func printer(msg string) error {
	defer fmt.Println("over")
	defer fmt.Print("?\n");
	_, err := fmt.Println(msg)
	return err
}